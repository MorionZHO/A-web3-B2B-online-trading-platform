package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

type Error struct {
	Message string
}

func (E *Error) Error() string {
	return E.Message
}

type Order struct {
	BuyerToken        string `json:"buyerToken"`
	SellerCompanyName string `json:"sellerCompanyName"`
	ProductName       string `json:"productName"`
	Address           string `json:"address"`
	OrderTime         string `json:"orderTime"`
	Id                string `json:"id"`

	UnitPrice  float64 `json:"unitPrice"`
	TotalPrice float64 `json:"totalPrice"`
	Number     int     `json:"number"`
	OrderState int     `json:"orderState"`
}

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) CreateOrder(ctx contractapi.TransactionContextInterface, Id string, BuyerToken string, CompanyName string, ProductName string, UintPrice float64, Address string, OrderTime string, Number int, TotalPrice float64, OrderState int) error {
	exists, err := s.OrderExist(ctx, Id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the order %s already exists", Id)
	}

	Order := Order{
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

func (s *SmartContract) ReadOrder(ctx contractapi.TransactionContextInterface, id string) (*Order, error) {
	orderJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if orderJSON == nil {
		return nil, fmt.Errorf("the order %s does not exist", id)
	}

	var order Order
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

	// overwriting original asset with new order
	order.OrderState = OrderState
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(Id, orderJSON)
}

// config = "world" and state = ""return all, config = "buyer" and have token return buyer
func (s *SmartContract) GetAllOrders(ctx contractapi.TransactionContextInterface, config string, Token string, OrderState int) ([]*Order, error) {
	if config == "world" {
		resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
		if err != nil {
			return nil, err
		}
		defer resultsIterator.Close()

		var orders []*Order
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return nil, err
			}
			log.Println(queryResponse.Key)
			var order Order
			err = json.Unmarshal(queryResponse.Value, &order)
			if err != nil {
				return nil, err
			}
			orders = append(orders, &order)
		}
		return orders, nil
	} else if config == "buyer" {
		queryString := fmt.Sprintf(`{"selector":{"buyerToken":"%s","orderState":%d}}`, Token, OrderState)
		resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
		if err != nil {
			return nil, err
		}
		defer resultsIterator.Close()
		var orders []*Order
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return nil, err
			}
			log.Println(queryResponse.Key)
			var order Order
			err = json.Unmarshal(queryResponse.Value, &order)
			if err != nil {
				return nil, err
			}
			orders = append(orders, &order)
		}
		return orders, nil
	} else if config == "seller" {
		queryString := fmt.Sprintf(`{"selector":{"sellerCompanyName":"%s"}}`, Token)
		resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
		if err != nil {
			return nil, err
		}
		defer resultsIterator.Close()
		var orders []*Order
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return nil, err
			}
			log.Println(queryResponse.Key)
			var order Order
			err = json.Unmarshal(queryResponse.Value, &order)
			if err != nil {
				return nil, err
			}
			orders = append(orders, &order)
		}
		return orders, nil
	} else {
		return nil, &Error{Message: "error parameter"}
	}
}

// return true if exist
func (s *SmartContract) OrderExist(ctx contractapi.TransactionContextInterface, Id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(Id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

func main() {
	orderChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating order-transfer-basic chaincode: %v", err)
	}
	if err := orderChaincode.Start(); err != nil {
		log.Panicf("Error starting order-transfer-basic chaincode: %v", err)
	}
}
