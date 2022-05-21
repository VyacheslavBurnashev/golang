package application

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func ApplicationStart() {
	mapUrls()
	router.Run(":8081")
}
