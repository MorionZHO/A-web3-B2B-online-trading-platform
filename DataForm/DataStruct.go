package DataForm

import "github.com/dgrijalva/jwt-go"

type User struct {
	PhoneNumber string `form:"phoneNumber"`
	PassWord    string `form:"password"`
	Role        string `form:"role"`
	Company     string `form:"companyName"`
	Token       string
}

type Claim struct {
	PhoneNumber string
	PassWord    string
	Role        string
	Company     string
	jwt.StandardClaims
}

type TempProduct struct {
	Id           int     `json:"id"`
	MerchantName string  `json:"merchantName"`
	ProductName  string  `json:"productName"`
	UnitPrice    float64 `json:"unitPrice"`
	ProductType  string  `json:"productType"`
	Inventory    int     `json:"inventory"`
	ImageName    string  `json:"imageName"`
	ImageData    string  `json:"imageData"` // 字符串类型
	ProductState int     `json:"productState"`
}

type Product struct {
	Id           int     `json:"id"`
	MerchantName string  `json:"merchantName"`
	ProductName  string  `json:"productName"`
	UnitPrice    float64 `json:"unitPrice"`
	ProductType  string  `json:"productType"`
	Inventory    int     `json:"inventory"`
	ImageName    string  `json:"imageName"`
	ImageData    []byte  `json:"imageData"`
	//1 表示上架，0表示售罄，不可购买等
	ProductState int `json:"productState"`
}

type SendProduct struct {
	Id           int     `json:"id"`
	MerchantName string  `json:"merchantName"`
	ProductName  string  `json:"productName"`
	UnitPrice    float64 `json:"unitPrice"`
	ProductType  string  `json:"productType"`
	Inventory    int     `json:"inventory"`
	ImageData    string  `json:"imageData"`
	//1 表示上架，0表示售罄，不可购买等
	ProductState int `json:"productState"`
}

func ConvertSend(product Product) *SendProduct {
	return &SendProduct{
		Id:           product.Id,
		MerchantName: product.MerchantName,
		ProductName:  product.ProductName,
		UnitPrice:    product.UnitPrice,
		ProductType:  product.ProductType,
		Inventory:    product.Inventory,
		ImageData:    string(product.ImageData),
		ProductState: product.ProductState,
	}
}

type Order struct {
	BuyerToken        string `json:"buyerToken"`
	SellerCompanyName string `json:"sellerCompanyName"`
	ProductName       string `json:"productName"`
	//买家收货地址
	Address   string `json:"address"`
	OrderTime string `json:"orderTime"`
	Id        string `json:"id"`

	UnitPrice  float64 `json:"unitPrice"`
	TotalPrice float64 `json:"totalPrice"`
	Number     int     `json:"number"`
	//0表示未发货，1表示已发货，2表示订单完成
	OrderState int `json:"orderState"`
}

type UserReq struct {
	Role       string `json:"role"`
	Token      string `json:"Token"`
	OrderState int    `json:"orderState"`
}
