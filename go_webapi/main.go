package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sunwenlei/golang/go_webapi/httpd"
)

func main() {
	fmt.Println("This is main")

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", httpd.Ping)
		api.GET("/user", httpd.Getusers)
		api.GET("/user/:personcd", httpd.Getuser)
		api.POST("/user", httpd.Createuser)
		api.PUT("/user/:personcd", httpd.Updateuser)
		api.DELETE("/user/:personcd", httpd.Deleteuser)

	}
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
