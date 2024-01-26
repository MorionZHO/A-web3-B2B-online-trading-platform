package DataBase

import (
	"Dapp/DataForm"
	"fmt"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, err := Dbconnect("root", "zhh0812", "user_data")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = CreateUserTable(db, "UserAccount")

	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateProduct(t *testing.T) {
	db, err := Dbconnect("root", "zhh0812", "product_data")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = CreateProductTable(db, "ProductInfo")

	if err != nil {
		log.Fatal(err)
	}
}

func TestInsetIMGToDB(t *testing.T) {
	db, err := Dbconnect("root", "zhh0812", "product_data")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = LocalIMGDataExtractToDB(db, "ProductInfo")

	if err != nil {
		log.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	db, err := Dbconnect("root", "zhh0812", "user_data")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	user := DataForm.User{
		PhoneNumber: "1010101001",
		PassWord:    "zhh",
		Role:        "buyer",
	}
	err = InsertUserToDB(db, "useraccount", user)
	if err != nil {
		log.Fatal(err)
	}
}

func TestCheck(t *testing.T) {
	db, err := Dbconnect("root", "zhh0812", "user_data")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	user := DataForm.User{
		PhoneNumber: "1010101001",
		PassWord:    "zhh",
	}
	err, phonenumber, role, token, company := CheckUserFromDB(db, "useraccount", user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hello! ", phonenumber, " +  ", role, " +  ", token, " +  ", company)
}

func TestGetProductDataFromDB(t *testing.T) {
	db, err := Dbconnect("root", "zhh0812", "product_data")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	var product DataForm.Product
	product.ProductName = "women's platform thick sole snow boots"
	product.MerchantName = "ZZH"
	err = GetProductDataFromDBbyName(db, "ProductInfo", product.MerchantName, product.ProductName)

	err = GetProductDataFromDBbyid(db, "ProductInfo", 1)

	if err != nil {
		log.Fatal(err)
	}
}
