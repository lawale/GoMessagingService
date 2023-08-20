package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.New()
}

func main() {

	_ = router.Run()
}
