package main

import (
	"Dapp/API"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(corsMiddleware)
	user := r.Group("/api/user")
	shop := r.Group("/api/shop")
	seller := r.Group("api/seller")
	search := r.Group("api/search")
	buyer := r.Group("api/buyer")
	order := r.Group("api/order")

	user.POST("/login", API.HandleLogin)
	user.POST("/register", API.HandleRegister)

	//shop.GET("", API.HandleShopByType)
	//shop.GET("", API.HandleShopByid)
	//shop?type= &id=
	shop.GET("", API.HandleShop)

	seller.POST("/shelfProduct", API.HandleShelf)
	//seller?seller= &state=
	seller.GET("", API.HandleShopState)
	seller.POST("updateOrder", API.HandleUpdateOrderState)

	//search?KeyWord=
	search.GET("", API.HandleSearch)

	buyer.POST("submitOrder", API.HandleSubmitOrder)

	order.POST("historyOrder", API.HandleSearchHistoryOrder)
	r.Run()

}

func corsMiddleware(c *gin.Context) {
	//Access-Control-Allow-Origin: 设置允许访问的来源，* 表示允许任何来源，你也可以设置为具体的域名，例如 "http://localhost:3000"。
	//Access-Control-Allow-Credentials: 允许浏览器发送包含凭证（如 cookies）的请求。如果需要在请求中发送 cookies，这个头部必须设置为 true。
	//Access-Control-Allow-Headers: 允许浏览器在实际请求中携带的首部字段，这里包括了一些常见的请求头。
	//Access-Control-Allow-Methods: 允许的请求方法，这里包括了 PUT、POST、GET、DELETE、OPTIONS。
	//OPTIONS 请求处理： 如果是预检请求（OPTIONS 请求），中间件直接返回 200 状态码，表示允许实际请求。

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	//请求成功后，服务器得发送200状态码才算连接成功
	//表示如果是预检请求（OPTIONS 请求），中间件直接返回 200 状态码，表示允许实际请求
	if c.Request.Method == "OPTIONS" {
		// 处理预检请求（preflight request）
		c.AbortWithStatus(200)
		return
	}

	c.Next()
}

//// LoadHTMLGlob("templates/*.html") 表示加载 templates 目录下所有扩展名为 .html 的文件
//func formatAsDate(t time.Time) string {
//	year, month, day := t.Date()
//	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
//}

//func main() {
//	router := gin.Default()
//	router.Delims("{[{", "}]}") //修改模板的分隔符，分隔符用于标记动态生成的内容
//	router.SetFuncMap(template.FuncMap{
//		"formatAsDate": formatAsDate,
//	})
//	router.LoadHTMLFiles("./templates/raw.tmpl")
//
//	router.GET("/raw", func(c *gin.Context) {
//		// | 表示管道操作符，用于将 .now 变量的值传递给名为 formatAsDate 的模板函数。
//		// .now 可以想象成一个map，取map["now"]中对应的数据,
//		//因为JavaScript中的Map是对象，访问键为key的值时候会使用Map.get(key)，这是为了和JS同步
//		//如果使用gin.H，也是一样
//		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
//			"_now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC), //关于接口的表
//		})
//	})
//
//	router.Run(":8080")
//}

//func main() {
//	r := gin.Default()
//	//定义路由的GET方法及响应处理函数
//	//当网页给后端端口发送get请求时，会使用下面这个r.GET函数
//	r.GET("/student", func(c *gin.Context) {
//		//将发送的信息封装成JSON发送给浏览器
//		c.JSON(http.StatusOK, gin.H{
//			//这是我们定义的数据
//			"message": "search student info successfully",
//		})
//	})
//
//	r.POST("/create_student", func(context *gin.Context) {
//		context.JSON(http.StatusOK, gin.H{
//			"message": "create student info successfully",
//		})
//	})
//	r.Run() //默认在本地8080端口启动服务
//}
