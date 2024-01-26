package API

import (
	"Dapp/DataBase"
	"Dapp/DataForm"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleShelf(c *gin.Context) {
	var product DataForm.Product
	var temp DataForm.TempProduct

	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Fatal(err)
	}

	//如果string，解析不能自动转换成[]byte，所以得建一个临时结构体
	imageData, err := base64.StdEncoding.DecodeString(temp.ImageData)
	if err != nil {
		// 处理 Base64 解码错误
	}

	product = DataForm.Product{
		Id:           temp.Id,
		MerchantName: temp.MerchantName,
		ProductName:  temp.ProductName,
		UnitPrice:    temp.UnitPrice,
		ProductType:  temp.ProductType,
		Inventory:    temp.Inventory,
		ImageName:    temp.ImageName,
		ImageData:    imageData, // 已解码的数据
		ProductState: temp.ProductState,
	}

	err = shelfProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":          http.StatusBadRequest,
			"error message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	}
}

func HandleShopState(c *gin.Context) {
	Seller := c.Query("seller")
	States := c.Query("state")
	State, err := strconv.Atoi(States)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "wrong input State type",
		})
		return
	}

	err, products := getProductByState(Seller, State)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "wrong type",
		})
	} else {
		// 转换二进制图片数据为 Base64
		for i, product := range products {
			base64Image := base64.StdEncoding.EncodeToString(product.ImageData)
			products[i].ImageData = []byte(base64Image) // 更新为 Base64 编码字符串
		}
		//
		//fmt.Println(string(products[0].ImageData))
		//file, err := os.OpenFile("example.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer file.Close()
		//
		//// 写入数据
		//_, err = file.WriteString(string(products[0].ImageData))
		//if err != nil {
		//	log.Fatal(err)
		//}

		c.JSON(http.StatusOK, products)
	}
}

func getProductByState(MerchantName string, ProductState int) (error, []DataForm.Product) {
	productDB, err := DataBase.Dbconnect(username, password, productDBname)
	defer productDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	err, products := DataBase.SearchProductFromDBbyState(productDB, "productinfo", MerchantName, ProductState)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return err, products
}

func shelfProduct(product DataForm.Product) error {
	productDB, err := DataBase.Dbconnect(username, password, productDBname)
	defer productDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = DataBase.ShelfProductProductToDB(productDB, "productinfo", product)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
