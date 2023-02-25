package routes

import (
	"github.com/pavle995/images/config"
	"github.com/pavle995/images/dal"

	"github.com/gin-gonic/gin"
)

type Router struct {
	dal    dal.Dal
	config *config.Config
}

func NewRouter(dal dal.Dal, cfg *config.Config) *Router {

	return &Router{
		dal:    dal,
		config: cfg,
	}
}

func InitRouter() *gin.Engine {
	engine := gin.Default()
	config := config.GetConfig()
	fs := dal.NewFileService(config)
	router := NewRouter(fs, config)
	engine.POST("/image", router.uploadImage)
	engine.GET("/image", router.getAll)
	engine.DELETE("/image/:fileName", router.delete)
	engine.GET("/image/:fileName", router.download)
	return engine
}
