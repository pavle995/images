package routes

import (
	"github.com/pavle995/images/dal"

	"github.com/gin-gonic/gin"
)

type Router struct {
	dal dal.Dal
}

func NewRouter() *Router {
	fs := dal.FileService{}
	return &Router{
		dal: &fs,
	}
}

func InitRouter() *gin.Engine {
	engine := gin.Default()
	router := NewRouter()
	engine.POST("/image", router.uploadImage)
	return engine
}
