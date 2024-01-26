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

// shop?type = ...
func HandleShopByType(c *gin.Context) {
	Type := c.Query("type")
	err, products := searchByType(Type)
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

// shop?id=...
func HandleShopByid(c *gin.Context) {
	ids := c.Query("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "wrong input id type",
		})
		return
	}

	err, product := searchByid(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "incorrect product id",
		})
	} else {
		// 转换二进制图片数据为 Base64
		base64Image := base64.StdEncoding.EncodeToString(product.ImageData)
		product.ImageData = []byte(base64Image) // 更新为 Base64 编码字符串
		c.JSON(http.StatusOK, product)
	}
}

func HandleSearch(c *gin.Context) {
	Name := c.Query("KeyWord")
	err, products := searchByName(Name)
	sendproducts := make([]DataForm.SendProduct, len(products))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err,
		})
		return
	} else {
		// 转换二进制图片数据为 Base64
		for i, product := range products {
			base64Image := base64.StdEncoding.EncodeToString(product.ImageData)
			products[i].ImageData = []byte(base64Image) // 更新为 Base64 编码字符串
			sendproducts[i] = *DataForm.ConvertSend(products[i])
		}
		c.JSON(http.StatusOK, sendproducts)
	}

}

// 优先级 type > id
func HandleShop(c *gin.Context) {
	Type := c.Query("type")
	ids := c.Query("id")
	if Type != "" {
		err, products := searchByType(Type)
		sendproducts := make([]DataForm.SendProduct, len(products))
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
				sendproducts[i] = *DataForm.ConvertSend(products[i])
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

			c.JSON(http.StatusOK, sendproducts)
		}
	} else if ids != "" {
		id, err := strconv.Atoi(ids)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "wrong input id type",
			})
			return
		}
		err, product := searchByid(id)

		//base64Image := base64.StdEncoding.EncodeToString(product.ImageData)
		//product.ImageData = []byte(base64Image) // 更新为 Base64 编码字符串
		//fmt.Println(product.ProductName)
		//fmt.Println(product.MerchantName)
		//
		//fmt.Println(string(product.ImageData))
		//file, err := os.OpenFile("example.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer file.Close()
		//
		//// 写入数据
		//_, err = file.WriteString(string(product.ImageData))
		//if err != nil {
		//	log.Fatal(err)
		//}

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "incorrect product id",
			})
		} else {
			// 转换二进制图片数据为 Base64
			base64Image := base64.StdEncoding.EncodeToString(product.ImageData)
			product.ImageData = []byte(base64Image) // 更新为 Base64 编码字符串
			sendproduct := *DataForm.ConvertSend(product)
			c.JSON(http.StatusOK, sendproduct)
		}
	} else if ids == "" && Type == "" {
		err, products := returnallproducts()
		sendproducts := make([]DataForm.SendProduct, len(products))
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
				sendproducts[i] = *DataForm.ConvertSend(products[i])
			}
			c.JSON(http.StatusOK, sendproducts)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "empty database",
		})
	}
}

func searchByType(productType string) (error, []DataForm.Product) {
	productDB, err := DataBase.Dbconnect(username, password, productDBname)
	defer productDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	err, products := DataBase.SearchProductFromDBbyType(productDB, "productinfo", productType)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, products
}

func searchByid(id int) (error, DataForm.Product) {
	productDB, err := DataBase.Dbconnect(username, password, productDBname)
	defer productDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	err, product := DataBase.SearchProductFromDbByid(productDB, "productinfo", id)
	product.Id = id
	if err != nil {
		log.Println(err)
		return err, DataForm.Product{}
	}
	return nil, product
}

func searchByName(name string) (error, []DataForm.Product) {
	productDB, err := DataBase.Dbconnect(username, password, productDBname)
	defer productDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	err, products := DataBase.SearchProductFromDbByNameF(productDB, "productinfo", name)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, products
}

func returnallproducts() (error, []DataForm.Product) {
	productDB, err := DataBase.Dbconnect(username, password, productDBname)
	defer productDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	err, products := DataBase.SearchAllProducts(productDB, "productinfo")
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, products
}
