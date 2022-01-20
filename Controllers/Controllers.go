package Controllers

import (
	"fmt"
	"net/http"

	"web-service-gin/Services"

	"github.com/gin-gonic/gin"
)

func InitControllers() {
	fmt.Println("controllers starting")
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Services.GetAlbum())
}
