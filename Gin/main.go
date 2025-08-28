package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

//func main() {
//	router := gin.Default() //means it will create a router with the default middleware (logger and recovery)
//	//all the HTTP methods are available here
//	router.GET("/getData", func(c *gin.Context) { // this gin Context has the information regarding the response as well as the request
//		c.JSON(200, gin.H{ //context.JSON - whatever we will have in c.JSON, will be responsed.....200 is the response code and hello world will be returned as an output in the form of json
//			"hello": "world",
//		})
//	})
//
//	//router.Use(middleware.Authenticate) // it will be applied to all the endpoints
//	router.GET("/getQueryString", middleware.Authenticate, getQueryString) // now only to this route middleware is applied
//	router.GET("/getUrlData/:name/:age", getUrlData)
//	router.POST("/getDataPost", getDataPost)
//
//	//authentication
//
//	auth := gin.BasicAuth(gin.Accounts{ //gin is using basic authetication , of type account, that means user has to paas username and password to authenticate
//		"user":  "pass",
//		"user2": "pass2",
//		"user3": "pass3",
//	})
//
//	//below is Route Grouping
//	//http://localhost:8080/admin/login
//	admin := router.Group("/admin", auth)
//	{
//		admin.POST("/login", login)
//	}
//
//	//http://localhost:8080/client/getData
//	client := router.Group("/client", auth)
//	{
//		client.GET("/getData", getData)
//	}
//
//	//router.Run(":8080") // default is 8080, if we want to run on any other port, then we can specify
//
//	//http.ListenAndServe(":9090", router) //it is same as router.run, in this i have done explicit work, while in the previos one, gin framework is doing the same stuff
//
//	server := &http.Server{ // by this method, you can pass custom HTTP configurations
//		Addr:         ":8080",
//		Handler:      router,
//		ReadTimeout:  10 * time.Second,
//		WriteTimeout: 10 * time.Second,
//	}
//
//	server.ListenAndServe()
//
//}
//
//// if you want to make your code a bit modular, you can simply remove the func(c *gin.Context) from there, give that function a name, and put the function name over there
//func getData(c *gin.Context) {
//	c.JSON(299, gin.H{
//		"hello": "world",
//	})
//}
//
//func login(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"login": "successful",
//	})
//}
//
//func getDataPost(c *gin.Context) {
//	body := c.Request.Body
//	value, _ := io.ReadAll(body)
//	c.JSON(200, gin.H{
//		"hi":   " I AM POSTING DATA ",
//		"body": string(value),
//	})
//}
//
//// http://localhost:8080/getQueryString?name=Mark&age=30
//// anything after the question mark is the query string
//
//func getQueryString(c *gin.Context) {
//	name := c.Query("name")
//	age := c.Query("age")
//	c.JSON(200, gin.H{
//		"data": "Hi, this is QueryString method",
//		"name": name,
//		"age":  age,
//	})
//}
//
//// http://localhost:8080/getUrlData/name/Mark/age/30
//func getUrlData(c *gin.Context) {
//	name := c.Param("name")
//	age := c.Param("age")
//	c.JSON(200, gin.H{
//		"data": "Hi, this is UrlData method",
//		"name": name,
//		"age":  age,
//	})
//}

// now i want to write the logging data in a file
func main() {
	router := gin.Default()
	f, _ := os.Create("ginLogging.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // this is writing the output in the file as well as in the standard input output.
	router.GET("/getdata", getData)
	router.Run()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hello world",
	})
}

//gin.forceConsoleColor() -> gives color encding to the output which is already in the mac
