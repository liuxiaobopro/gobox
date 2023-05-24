package main

import (
	"github.com/liuxiaobopro/gobox/gin/ctx"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", ctx.Use())

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
