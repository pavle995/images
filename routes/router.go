package routes

import (
	"github.com/pavle995/images/config"
	"github.com/pavle995/images/dal"

	"github.com/gin-gonic/gin"
)

type Router struct {
	dal    dal.Dal
	config config.Config
}

func NewRouter() *Router {
	config := config.GetConfig()
	fs := dal.NewFileService(config)
	return &Router{
		dal:    fs,
		config: *config,
	}
}

func InitRouter() *gin.Engine {
	engine := gin.Default()
	router := NewRouter()
	engine.POST("/image", router.uploadImage)
	engine.GET("/image", router.getAll)
	engine.DELETE("/image/:fileName", router.delete)
	engine.GET("/image/:fileName", router.download)
	return engine
}
