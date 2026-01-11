package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeMiddleware(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	since := time.Since(startTime)
	// 获取当前请求所对应的函数
	f := c.HandlerName()
	fmt.Printf("函数 %s 耗时 %d\n", f, since)
}

func main() {
	router := gin.Default()
	router.Use(TimeMiddleware)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	// 结构体转json
	router.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		type Msg struct {
			Name    string `json:"user"`
			Message string `json:"message"`
			Number  int    `json:"number"`
		}
		msg := Msg{"fengfeng", "hey", 21}
		// 注意 msg.Name 变成了 "user" 字段
		// 以下方式都会输出 :   {"user": "hanru", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
