package API

import (
	"Dapp/DataBase"
	"Dapp/DataForm"
	"bytes"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	mspID        = "Org1MSP"
	cryptoPath   = "../../fabric/project/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com"
	certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/User1@org1.example.com-cert.pem"
	keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
	tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	peerEndpoint = "localhost:7051"
	gatewayPeer  = "peer0.org1.example.com"
)

func HandleSubmitOrder(c *gin.Context) {
	var order DataForm.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		log.Fatal(err)
	}

	productDB, err := DataBase.Dbconnect(username, password, productDBname)
	defer productDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	err, number := DataBase.DeleteInventoryToDB(productDB, "productinfo", order.ProductName, order.SellerCompanyName, order.Number)
	if err != nil {
		log.Fatal(err)
	}
	//contract, connection := connectToFabric()
	//defer connection.Close()

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection()

	id := newIdentity()
	sign := newSign()

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "order"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname

	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	err = createOrder(contract, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":          http.StatusBadRequest,
			"error message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"number": number,
		})
	}
	getAllOrders(contract, "world", "", "")

}

// order/historyOrder
func HandleSearchHistoryOrder(c *gin.Context) {
	var userreq DataForm.UserReq
	err := c.ShouldBindJSON(&userreq)
	if err != nil {
		log.Fatal(err)
	}

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection()

	id := newIdentity()
	sign := newSign()

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "order"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname

	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)
	orderstate := strconv.Itoa(userreq.OrderState)
	//orderstate = 0为未完成,目前暂定用0测试
	orders, err := getAllOrders(contract, userreq.Role, userreq.Token, orderstate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":          http.StatusBadRequest,
			"error message": err,
		})
	} else {
		c.JSON(http.StatusOK, orders)
	}
}

// seller/updateOrder?orderId = & state =
func HandleUpdateOrderState(c *gin.Context) {
	orderId := c.Query("orderId")
	state := c.Query("state")

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection()

	id := newIdentity()
	sign := newSign()

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "order"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname

	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)
	err = updateOrder(contract, orderId, state)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":          http.StatusBadRequest,
			"error message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	}
}

func updateOrder(contract *client.Contract, orderId string, OrderState string) error {
	fmt.Printf("\n--> Submit Transaction: UpdateOrder \n")

	_, err := contract.SubmitTransaction("UpdateOrder", orderId, OrderState)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
		return err
	}

	fmt.Printf("*** Transaction committed successfully\n")
	return err
}

// config = "world"并且另一个参数为""返回全部, config = "buyer"并且得包含自己token返回buyer自己的订单,conifg = "seller" token需要包含companyname
func getAllOrders(contract *client.Contract, config string, BuyerToken string, OrderState string) ([]DataForm.Order, error) {
	fmt.Println("\n--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	var orders []DataForm.Order
	evaluateResult, err := contract.EvaluateTransaction("GetAllOrders", config, BuyerToken, OrderState)
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
		return nil, err
	}
	err = json.Unmarshal(evaluateResult, &orders)
	if err != nil {
		log.Fatalf("Error occurred during unmarshaling. Error: %s", err.Error())
		return nil, err
	}

	//用于测试，print
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)

	return orders, nil
}

// Submit a transaction synchronously, blocking until it has been committed to the ledger.
func createOrder(contract *client.Contract, order DataForm.Order) error {
	fmt.Printf("\n--> Submit Transaction: CreateOrder \n")

	//虽然定义合约中的createasset函数的时候，size是int类型，但是调用这个函数的时候必须要用string类型进行传递
	//合约接口会自动尝试将string类型转换成相应类型，如果出错会返回error
	_, err := contract.SubmitTransaction("CreateOrder", order.Id, order.BuyerToken, order.SellerCompanyName, order.ProductName, fmt.Sprintf("%.2f", order.UnitPrice), order.Address, order.OrderTime, fmt.Sprintf("%d", order.Number), fmt.Sprintf("%.2f", order.TotalPrice), fmt.Sprintf("%d", order.OrderState))
	if err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("failed to submit transaction: %w", err))
		return err
	}

	fmt.Printf("*** Transaction committed successfully\n")
	return err
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		panic(err)
	}

	//使用 Org1 用户的 X.509 证书作为客户端身份，并使用签名实现 基于该用户的私钥
	certPool := x509.NewCertPool()
	//添加证书
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	files, err := os.ReadDir(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(keyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

// Format JSON data
func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}
