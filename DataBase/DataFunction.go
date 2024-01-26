package DataBase

import (
	"Dapp/DataForm"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
)

func InsertUserToDB(db *sql.DB, tablename string, user DataForm.User) error {
	//有点小bug，插入失败也会递增id
	query := fmt.Sprintf(
		`INSERT INTO %s (PhoneNumber,PassWord,UserRole,Token,Company)
                VALUES ('%s','%s','%s','%s','%s')`, tablename, user.PhoneNumber, user.PassWord, user.Role, user.Token, user.Company)

	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		if err.Error()[20] == 'D' && err.Error()[21] == 'u' {
			return &Error{Message: "Duplicated user form"}
		} else {
			return &Error{Message: "wrong user form"}
		}
	}

	return err
}

func ShelfProductProductToDB(db *sql.DB, tablename string, product DataForm.Product) error {
	//不能用这个，这样会因为字符串转义的时候有些特殊字符，比如“'”而出错
	//query := fmt.Sprintf(
	//	`INSERT INTO %s (MerchantName,ProductName,UnitPrice,ProductType,Inventory,ImageName,ImageData,ProductState)
	//            VALUES ('%s','%s','%f','%s','%d','%s','%v','%d')`, tablename, product.MerchantName, product.ProductName, product.UnitPrice, product.ProductType, product.Inventory, product.ImageName, product.ImageData, product.ProductState)
	searchquery := fmt.Sprintf(
		`SELECT id,MerchantName,ProductName FROM %s 
		WHERE MerchantName = ? AND ProductName = ?`, tablename)
	var merchantname, productname string
	var id int
	err := db.QueryRow(searchquery, product.MerchantName, product.ProductName).Scan(&id, &merchantname, &productname)
	if err != nil {
		if err == sql.ErrNoRows {
			// 查询未找到任何匹配的行,执行插入
			query := fmt.Sprintf(
				`INSERT INTO %s (MerchantName,ProductName,UnitPrice,ProductType,Inventory,ImageName,ImageData,ProductState)
				VALUES (?,?,?,?,?,?,?,?)`, tablename)

			_, err := db.Exec(query, product.MerchantName, product.ProductName, product.UnitPrice, product.ProductType, product.Inventory, product.ImageName, product.ImageData, product.ProductState)

			//_, err := db.Exec(query)
			if err != nil {
				fmt.Println(err)
				if err.Error()[20] == 'D' && err.Error()[21] == 'u' {
					return &Error{Message: "Duplicated product form"}
				} else {
					return &Error{Message: "wrong product form"}
				}
			} else {
				return nil
			}
		} else {
			fmt.Println(err)
			return err
			// 处理其他错误
		}
	}
	//执行修改
	product.Id = id
	err = editProductToDB(db, tablename, product)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

// 根据名字和卖家找商品
func SearchProductFromDbByName(db *sql.DB, tablename string, MerchantName string, ProductName string) (error, DataForm.Product) {
	query := fmt.Sprintf(
		`SELECT id,MerchantName,ProductName,UnitPrice,ProductType,Inventory,ImageData,ProductState FROM %s 
		WHERE MerchantName = ? AND ProductName = ?`, tablename)

	//执行查询并且返回匹配的行(只返回一行)
	row := db.QueryRow(query, MerchantName, ProductName)
	var output DataForm.Product

	err := row.Scan(&output.Id, &output.MerchantName, &output.ProductName, &output.UnitPrice, &output.ProductType, &output.Inventory, &output.ImageData, &output.ProductState)
	if err == sql.ErrNoRows {
		return &Error{Message: "incorrect MerchantName or ProductName"}, output
	} else if err != nil {
		log.Fatal(err)
		return nil, output
	} else {
		return nil, output
	}
	return nil, output
}

// 根据名字模糊搜索
func SearchProductFromDbByNameF(db *sql.DB, tablename string, ProductName string) (error, []DataForm.Product) {
	query := fmt.Sprintf(
		`SELECT id,MerchantName,ProductName,UnitPrice,ProductType,Inventory,ImageData,ProductState FROM %s 
		WHERE ProductName LIKE ?`, tablename)

	//执行查询并且返回匹配的行(只返回一行)
	rows, err := db.Query(query, "%"+ProductName+"%")
	if err != nil {
		log.Fatal(err)
		return err, nil
	}
	defer rows.Close() // 不要忘记关闭结果集

	var outputs []DataForm.Product
	for rows.Next() {
		var output DataForm.Product
		// 使用 Scan 从当前行中提取数据
		err = rows.Scan(&output.Id, &output.MerchantName, &output.ProductName, &output.UnitPrice, &output.ProductType, &output.Inventory, &output.ImageData, &output.ProductState)
		if err != nil {
			log.Fatal(err)
			return err, nil
		}
		outputs = append(outputs, output)
	}
	// 检查遍历过程中是否有任何错误发生
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return err, nil
	}
	return nil, outputs
}

// 根据商品id找
func SearchProductFromDbByid(db *sql.DB, tablename string, id int) (error, DataForm.Product) {
	query := fmt.Sprintf(
		`SELECT MerchantName,ProductName,UnitPrice,ProductType,Inventory,ImageData,ProductState FROM %s 
		WHERE id = ?`, tablename)

	//执行查询并且返回匹配的行(只返回一行)
	row := db.QueryRow(query, id)
	var output DataForm.Product

	err := row.Scan(&output.MerchantName, &output.ProductName, &output.UnitPrice, &output.ProductType, &output.Inventory, &output.ImageData, &output.ProductState)
	if err == sql.ErrNoRows {
		return &Error{Message: "incorrect product id"}, output
	} else if err != nil {
		log.Fatal(err)
		return nil, output
	} else {
		return nil, output
	}
	return nil, output
}

func SearchProductFromDBbyState(db *sql.DB, tablename string, MerchantName string, State int) (error, []DataForm.Product) {
	query := fmt.Sprintf(
		`SELECT id,MerchantName,ProductName,UnitPrice,ProductType,Inventory,ImageData,ProductState FROM %s 
		WHERE MerchantName = ? AND ProductState = ?`, tablename)

	rows, err := db.Query(query, MerchantName, State)
	if err != nil {
		log.Fatal(err)
		return err, nil
	}
	defer rows.Close() // 不要忘记关闭结果集

	var outputs []DataForm.Product
	for rows.Next() {
		var output DataForm.Product
		// 使用 Scan 从当前行中提取数据
		err = rows.Scan(&output.Id, &output.MerchantName, &output.ProductName, &output.UnitPrice, &output.ProductType, &output.Inventory, &output.ImageData, &output.ProductState)
		if err != nil {
			log.Fatal(err)
			return err, nil
		}
		outputs = append(outputs, output)
	}
	// 检查遍历过程中是否有任何错误发生
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return err, nil
	}
	return nil, outputs
}

// 根据种类找商品
func SearchProductFromDBbyType(db *sql.DB, tablename string, ProductType string) (error, []DataForm.Product) {
	query := fmt.Sprintf(
		`SELECT id,MerchantName,ProductName,UnitPrice,ProductType,Inventory,ImageData,ProductState FROM %s 
		WHERE ProductType = ?`, tablename)

	rows, err := db.Query(query, ProductType)
	if err != nil {
		log.Fatal(err)
		return err, nil
	}
	defer rows.Close() // 不要忘记关闭结果集

	var outputs []DataForm.Product
	for rows.Next() {
		var output DataForm.Product
		// 使用 Scan 从当前行中提取数据
		err = rows.Scan(&output.Id, &output.MerchantName, &output.ProductName, &output.UnitPrice, &output.ProductType, &output.Inventory, &output.ImageData, &output.ProductState)
		if err != nil {
			log.Fatal(err)
			return err, nil
		}
		outputs = append(outputs, output)
	}
	// 检查遍历过程中是否有任何错误发生
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return err, nil
	}
	return nil, outputs
}

// 查找所有商品
func SearchAllProducts(db *sql.DB, tablename string) (error, []DataForm.Product) {
	query := fmt.Sprintf(
		`SELECT id,MerchantName,ProductName,UnitPrice,ProductType,Inventory,ImageData,ProductState FROM %s`, tablename)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return err, nil
	}
	defer rows.Close() // 不要忘记关闭结果集

	var outputs []DataForm.Product
	for rows.Next() {
		var output DataForm.Product
		// 使用 Scan 从当前行中提取数据
		err = rows.Scan(&output.Id, &output.MerchantName, &output.ProductName, &output.UnitPrice, &output.ProductType, &output.Inventory, &output.ImageData, &output.ProductState)
		if err != nil {
			log.Fatal(err)
			return err, nil
		}
		outputs = append(outputs, output)
	}
	// 检查遍历过程中是否有任何错误发生
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return err, nil
	}
	return nil, outputs
}

func DeleteInventoryToDB(db *sql.DB, tablename string, productname string, merchantname string, number int) (error, int) {
	err, product := SearchProductFromDbByName(db, tablename, merchantname, productname)
	if err != nil {
		return err, -1
	}
	FinalNumber := product.Inventory - number
	query := fmt.Sprintf(
		`UPDATE %s 
            SET Inventory = ?
            WHERE MerchantName = ? AND ProductName = ?`, tablename)

	_, err = db.Exec(query, FinalNumber, merchantname, productname)
	if err != nil {
		return err, -1
	}
	return err, FinalNumber
}

// 修改商品信息
func editProductToDB(db *sql.DB, tablename string, product DataForm.Product) error {
	query := fmt.Sprintf(
		`UPDATE %s 
            SET id = ?,ProductName = ?, UnitPrice = ?, ProductType = ?, Inventory = ?, ImageName = ?, ImageData = ?, ProductState = ?
            WHERE MerchantName = ? AND ProductName = ?`, tablename)

	_, err := db.Exec(query, product.Id, product.ProductName, product.UnitPrice, product.ProductType, product.Inventory, product.ImageName, product.ImageData, product.ProductState, product.MerchantName, product.ProductName)
	if err != nil {
		fmt.Println(err)
		if err.Error()[20] == 'D' && err.Error()[21] == 'u' {
			return &Error{Message: "Duplicated product form"}
		} else {
			return &Error{Message: "wrong product form"}
		}
	}
	return err
}

func InsertProductToChain() error {
	return nil
}

func CheckUserFromDB(db *sql.DB, tablename string, user DataForm.User) (error, string, string, string, string) {
	query := fmt.Sprintf(
		`SELECT PhoneNumber,UserRole,Token,Company FROM %s WHERE PhoneNumber = ? AND PassWord = ?`, tablename)
	//执行查询并且返回匹配的行
	row := db.QueryRow(query, user.PhoneNumber, user.PassWord)

	var PhoneNumber, UserRole, Token, Company string
	err := row.Scan(&PhoneNumber, &UserRole, &Token, &Company)
	if err == sql.ErrNoRows {
		return &Error{Message: "incorrect username or password"}, "", "", "", ""
	} else if err != nil {
		log.Fatal(err)
	} else {
		return nil, PhoneNumber, UserRole, Token, Company
	}
	return nil, PhoneNumber, UserRole, Token, Company
}

func CreateProductTable(db *sql.DB, tablename string) error {

	//image_data 列的类型是 LONGBLOB，用于存储大量的二进制数据，如图片
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s(
        id INT AUTO_INCREMENT PRIMARY KEY,
        MerchantName VARCHAR(255) NOT NULL,
        ProductName VARCHAR(255) NOT NULL,
        UnitPrice INT NOT NULL,
        ProductType VARCHAR(64) NOT NULL,
        Inventory INT NOT NULL,
        ImageName VARCHAR(255) NOT NULL,
        ImageData LONGBLOB,
    	ProductState INT NOT NULL,
        UNIQUE KEY Unique_Merchant_Product(MerchantName,Productname)
        )`, tablename)

	_, err := db.Exec(query)
	return err
}

func CreateUserTable(db *sql.DB, tablename string) error {
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
        id INT AUTO_INCREMENT PRIMARY KEY,
        PhoneNumber VARCHAR(20) NOT NULL UNIQUE,
        Password VARCHAR(20) NOT NULL,
        UserRole ENUM('buyer','seller') NOT NULL
        )`, tablename)

	_, err := db.Exec(query)
	return err
}

// Dbconnect 数据存放位置为E:\MySQL\mysql-8.0.27-winx64\data
func Dbconnect(username string, password string, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", username, password, dbname)
	db, err := sql.Open("mysql", dsn) //意思是mysql驱动程序
	return db, err
}

func LocalIMGDataExtractToDB(db *sql.DB, tablename string) error {
	file, err := os.Open("../WebRoot/products.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return err
	}
	defer file.Close()

	// 创建一个 Person 结构体实例
	var products []DataForm.Product

	// 创建一个 JSON 解码器
	decoder := json.NewDecoder(file)

	// 使用解码器解码 JSON 数据到结构体
	err = decoder.Decode(&products)
	//fmt.Println(products[0].UnitPrice)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err
	}

	for i := 0; i < len(products); i++ {
		imagePath := "../WebRoot/product_image/" + products[i].ImageName
		products[i].ImageData, err = ioutil.ReadFile(imagePath)
		if err != nil {
			fmt.Println("Error reading image file:", err)
			return err
		}
		err = ShelfProductProductToDB(db, tablename, products[i])
		if err != nil {
			fmt.Println("Error insert to mysql:", err)
			return err
		}
	}

	//fmt.Println(products[0].ImageName, products[0].ImageData)
	if err != nil {
		fmt.Println("Error walking through directory:", err)
		return err
	}

	return nil
}

// 从数据库中获取数据，需要输入表名和商品信息,商品名和卖家
func GetProductDataFromDBbyName(db *sql.DB, tablename string, MerchantName string, ProductName string) error {

	outputPath := "../WebRoot/output.jpg"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return err
	}
	defer outputFile.Close()

	fmt.Println(ProductName)
	// 将图片二进制数据解码为图片对象
	err, p := SearchProductFromDbByName(db, tablename, MerchantName, ProductName)

	if err != nil {
		fmt.Println("Error Search:", err)
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(p.ImageData))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return err
	}

	// 将图片对象编码为 JPEG 格式并写入文件
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return err
	}

	return nil
}

// 从数据库中获取数据，需要输入表名和商品id
func GetProductDataFromDBbyid(db *sql.DB, tablename string, Productid int) error {

	outputPath := "../WebRoot/outputID.jpg"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return err
	}
	defer outputFile.Close()

	// 将图片二进制数据解码为图片对象
	err, p := SearchProductFromDbByid(db, tablename, Productid)

	if err != nil {
		fmt.Println("Error Search:", err)
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(p.ImageData))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return err
	}

	// 将图片对象编码为 JPEG 格式并写入文件
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return err
	}

	return nil
}
