package routes

import (
	"github.com/pavle995/images/config"
	"github.com/pavle995/images/dal"

	"github.com/gin-gonic/gin"
)

type Router struct {
	dal dal.Dal
}

func NewRouter() *Router {
	config := config.GetConfig()
	fs := dal.NewFileService(config)
	return &Router{
		dal: fs,
	}
}

func InitRouter() *gin.Engine {
	engine := gin.Default()
	router := NewRouter()
	engine.POST("/image", router.uploadImage)
	engine.GET("/image", router.getAll)
	return engine
}
