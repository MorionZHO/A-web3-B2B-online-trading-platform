package API

import (
	"Dapp/DataBase"
	"Dapp/DataForm"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

const PrivateKey = "6344aa8cb57b62a974352044693097e4cc36a975f92e20468e7539db37d6d654"

// 处理post请求，请求参数为json
func HandleRegister(c *gin.Context) {
	var user DataForm.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.PhoneNumber)
	user.Token, err = generateToken(user.PhoneNumber, user.PassWord, user.Role, user.Company)

	if err != nil {
		log.Fatal(err)
	}
	err = insertUserData(user)
	if err != nil {
		//两种，"Duplicated user form" 或者 "wrong user form"
		c.JSON(http.StatusBadRequest, gin.H{
			"code":          http.StatusBadRequest,
			"error message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":          http.StatusOK,
			"error message": err,
		})
	}

}

// 处理post请求，请求参数为json
func HandleLogin(c *gin.Context) {
	//获取json参数，并且绑定结构体
	var user DataForm.User
	//通过判断content-type参数，如果参数是json，那么绑定，不是，那么报错
	err := c.ShouldBind(&user)
	if err != nil {
		log.Fatal(err)
	}
	err, phonenumber, role, token, company := checkUserData(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":          http.StatusNotFound,
			"error message": err.Error(),
		})
	} else {
		if company != "" {
			c.JSON(http.StatusOK, gin.H{
				"userInfo": gin.H{
					"phoneNumber": phonenumber,
					"role":        role,
					"companyName": company},
				"token": token,
				"code":  http.StatusOK,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"userInfo": gin.H{
					"phoneNumber": phonenumber,
					"role":        role},
				"token": token,
				"code":  http.StatusOK,
			})
		}
	}

	//fmt.Println(user)
}

func checkUserData(user DataForm.User) (error, string, string, string, string) {
	userDB, err := DataBase.Dbconnect(username, password, userDBname)
	defer userDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	err, phonenumber, role, token, company := DataBase.CheckUserFromDB(userDB, "useraccount", user)
	if err != nil {
		log.Println(err)
		return err, "", "", "", ""
	}
	return nil, phonenumber, role, token, company
}

func insertUserData(user DataForm.User) error {
	userDB, err := DataBase.Dbconnect(username, password, userDBname)
	defer userDB.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = DataBase.InsertUserToDB(userDB, "useraccount", user)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

// 生成 JWT Token
func generateToken(phonenumber string, password string, role string, company string) (string, error) {
	// 设置 token 有效期等信息
	claims := DataForm.Claim{
		phonenumber,
		password,
		role,
		company,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "Web3Trade",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(PrivateKey)) // 使用密钥签名 token
}
