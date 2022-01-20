package Router

import (
	"fmt"

	"web-service-gin/Controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func InitRouter() {
	fmt.Println("router staring")
	if true {
		RouteV1()
	}
	router.Run("localhost:8080")
}

func RouteV1() {
	v1 := router.Group("v1")
	v1.GET("/login", Controllers.GetAlbums)
}
