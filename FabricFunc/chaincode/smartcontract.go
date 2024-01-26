package chaincode

import (
	"Dapp/DataBase"
	"Dapp/DataForm"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

type SmartContract struct {
	contractapi.Contract
}

// 参数都得是string,所以前面结构体得把其中的int属性改成string
func (s *SmartContract) CreateOrder(ctx contractapi.TransactionContextInterface, Id string, BuyerToken string, CompanyName string, ProductName string, UintPrice float64, Address string, OrderTime string, Number int, TotalPrice float64, OrderState int) error {
	exists, err := s.OrderExist(ctx, Id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", Id)
	}

	Order := DataForm.Order{
		Id:                Id,
		BuyerToken:        BuyerToken,
		SellerCompanyName: CompanyName,
		ProductName:       ProductName,
		UnitPrice:         UintPrice,
		Address:           Address,
		OrderState:        OrderState,
		OrderTime:         OrderTime,
		Number:            Number,
		TotalPrice:        TotalPrice,
	}

	OrderJSON, err := json.Marshal(Order)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(Id, OrderJSON)
}

// 通过id（key）寻找订单
func (s *SmartContract) ReadOrder(ctx contractapi.TransactionContextInterface, id string) (*DataForm.Order, error) {
	orderJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if orderJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var order DataForm.Order
	err = json.Unmarshal(orderJSON, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *SmartContract) UpdateOrder(ctx contractapi.TransactionContextInterface, Id string, OrderState int) error {
	order, err := s.ReadOrder(ctx, Id)
	if err != nil {
		return err
	}

	// overwriting original asset with new asset
	order.OrderState = OrderState
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	//查找，修改，覆盖
	return ctx.GetStub().PutState(Id, orderJSON)
}

// config = "world"并且另一个参数为""返回全部, config = "buyer"并且得包含自己token返回buyer自己的订单
func (s *SmartContract) GetAllOrders(ctx contractapi.TransactionContextInterface, config string, buyerToken string, OrderState int) ([]*DataForm.Order, error) {
	if config == "world" {
		resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
		if err != nil {
			return nil, err
		}
		defer resultsIterator.Close()

		var orders []*DataForm.Order
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return nil, err
			}
			log.Println(queryResponse.Key)
			var order DataForm.Order
			err = json.Unmarshal(queryResponse.Value, &order)
			if err != nil {
				return nil, err
			}
			orders = append(orders, &order)
		}
		return orders, nil
	} else if config == "buyer" {
		//使用富查询
		queryString := fmt.Sprintf(`{"selector":{"buyerToken":"%s","orderState":"%d"}}`, buyerToken, OrderState)
		resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
		if err != nil {
			return nil, err
		}
		defer resultsIterator.Close()
		var orders []*DataForm.Order
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return nil, err
			}
			log.Println(queryResponse.Key)
			var order DataForm.Order
			err = json.Unmarshal(queryResponse.Value, &order)
			if err != nil {
				return nil, err
			}
			orders = append(orders, &order)
		}
		return orders, nil
	} else {
		return nil, &DataBase.Error{Message: "error parameter"}
	}
	return nil, &DataBase.Error{Message: "error parameter"}
}

// 存在则return true
func (s *SmartContract) OrderExist(ctx contractapi.TransactionContextInterface, Id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(Id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}
