package FabricFunc

import (
	"Dapp/DataForm"
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"github.com/hyperledger/fabric-protos-go-apiv2/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"os"
	"path"
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

//var now = time.Now()
//var orderId = fmt.Sprintf("order%d", now.Unix()*1e3+int64(now.Nanosecond())/1e6)

func StartConnectToFabric() {

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

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
	config := "world"
	orderState := "0"
	//orderid := "order1"
	order := &DataForm.Order{
		Id:                "order1",
		SellerCompanyName: "ZHH",
		ProductName:       "alaniz cup",
		UnitPrice:         19,
		OrderState:        0,
		Number:            10,
		TotalPrice:        190,
		Address:           "ZHH HOME",
		OrderTime:         "2024/1/25",
		BuyerToken:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQaG9uZU51bWJlciI6Iis2NTEyMyIsIlBhc3NXb3JkIjoiMTIzIiwiUm9sZSI6ImJ1eWVyIiwiQ29tcGFueSI6IiIsImV4cCI6MTcwNjIzNzE2NiwiaWF0IjoxNzA2MTUwNzY2LCJpc3MiOiJXZWIzVHJhZGUifQ.AKDFf2-jOrkSzr1DnCy3h5NAiCQjrulQgvz_m90Acaw",
	}
	createOrder(contract, *order)
	updateOrder(contract, "order1", "1")
	getAllOrders(contract, config, "", orderState)
	exampleErrorHandling(contract)
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

// config = "world"并且另一个参数为""返回全部, config = "buyer"并且得包含自己token返回buyer自己的订单,conifg = "seller" token需要包含companyname
func getAllOrders(contract *client.Contract, config string, BuyerToken string, OrderState string) {
	fmt.Println("\n--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("GetAllOrders", config, BuyerToken, OrderState)
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
}

func updateOrder(contract *client.Contract, orderId string, OrderState string) {
	fmt.Printf("\n--> Submit Transaction: UpdateOrder \n")

	_, err := contract.SubmitTransaction("UpdateOrder", orderId, OrderState)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

// Submit a transaction synchronously, blocking until it has been committed to the ledger.
func createOrder(contract *client.Contract, order DataForm.Order) {
	fmt.Printf("\n--> Submit Transaction: CreateOrder \n")

	//虽然定义合约中的createasset函数的时候，size是int类型，但是调用这个函数的时候必须要用string类型进行传递
	//合约接口会自动尝试将string类型转换成相应类型，如果出错会返回error
	_, err := contract.SubmitTransaction("CreateOrder", order.Id, order.BuyerToken, order.SellerCompanyName, order.ProductName, fmt.Sprintf("%.2f", order.UnitPrice), order.Address, order.OrderTime, fmt.Sprintf("%d", order.Number), fmt.Sprintf("%.2f", order.TotalPrice), fmt.Sprintf("%d", order.OrderState))
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

// Submit transaction, passing in the wrong number of arguments ,expected to throw an error containing details of any error responses from the smart contract.
func exampleErrorHandling(contract *client.Contract) {
	fmt.Println("\n--> Submit Transaction: UpdateAsset order, order does not exist and should return an error")

	_, err := contract.SubmitTransaction("UpdateOrder", "order1", "2")
	if err == nil {
		panic("******** FAILED to return an error")
	}

	fmt.Println("*** Successfully caught the error:")

	switch err := err.(type) {
	case *client.EndorseError:
		fmt.Printf("Endorse error for transaction %s with gRPC status %v: %s\n", err.TransactionID, status.Code(err), err)
	case *client.SubmitError:
		fmt.Printf("Submit error for transaction %s with gRPC status %v: %s\n", err.TransactionID, status.Code(err), err)
	case *client.CommitStatusError:
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Printf("Timeout waiting for transaction %s commit status: %s", err.TransactionID, err)
		} else {
			fmt.Printf("Error obtaining commit status for transaction %s with gRPC status %v: %s\n", err.TransactionID, status.Code(err), err)
		}
	case *client.CommitError:
		fmt.Printf("Transaction %s failed to commit with status %d: %s\n", err.TransactionID, int32(err.Code), err)
	default:
		panic(fmt.Errorf("unexpected error type %T: %w", err, err))
	}

	// Any error that originates from a peer or orderer node external to the gateway will have its details
	// embedded within the gRPC status error. The following code shows how to extract that.
	statusErr := status.Convert(err)

	details := statusErr.Details()
	if len(details) > 0 {
		fmt.Println("Error Details:")

		for _, detail := range details {
			switch detail := detail.(type) {
			case *gateway.ErrorDetail:
				fmt.Printf("- address: %s, mspId: %s, message: %s\n", detail.Address, detail.MspId, detail.Message)
			}
		}
	}
}

// Format JSON data
func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}
