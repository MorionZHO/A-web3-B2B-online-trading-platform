package DataBase

import (
	"Dapp/DataForm"
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"testing"
)

func TestExtract(t *testing.T) {
	file, err := os.Open("../WebRoot/products.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// 创建一个 Person 结构体实例
	var products []DataForm.Product

	// 创建一个 JSON 解码器
	decoder := json.NewDecoder(file)

	// 使用解码器解码 JSON 数据到结构体
	err = decoder.Decode(&products)
	fmt.Println(products[0].UnitPrice)
	fmt.Println(products[1])
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for i := 0; i < len(products); i++ {
		imagePath := "../WebRoot/product_image/" + products[i].ImageName
		products[i].ImageData, err = ioutil.ReadFile(imagePath)
		if err != nil {
			fmt.Println("Error reading image file:", err)
		}
	}

	fmt.Println(products[28].ProductName, products[28].ImageName, products[28].ImageData)
	if err != nil {
		fmt.Println("Error walking through directory:", err)
		return
	}

	outputPath := "../WebRoot/output.jpg"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// 将图片二进制数据解码为图片对象
	img, _, err := image.Decode(bytes.NewReader(products[28].ImageData))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// 将图片对象编码为 JPEG 格式并写入文件
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}
}
