package main

import (
	"DirectChat/internal/routing"
	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	routing.RegisterRoute(&r.RouterGroup)
	routing.RegisterSwagger(&r.RouterGroup)	
	r.Run() // listen and serve on 0.0.0.0:8080
}
