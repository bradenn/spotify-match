package server

import (
	"github.com/gin-gonic/gin"
	"spotify-match/controllers"
)

func Init() {
	r := NewRouter()
	_ = r.Run(":8080")
}

func NewRouter() *gin.Engine {
	r := gin.New()

	api := r.Group("api")
	{
		compare := new(controllers.Compare)
		api.POST("/compare", compare.CompareController)
	}
	return r
}
