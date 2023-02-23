package routes

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/pavle995/images/util"

	"github.com/gin-gonic/gin"
)

func (r *Router) uploadImage(c *gin.Context) {
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if image == nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	buffer, err := util.ReadImage(image)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	imgName := util.GetImageName(buffer)

	err = r.dal.StoreFile(buffer, imgName)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	fmt.Println("------------IMAGE------------")
	//fmt.Printf("%v\n", image)
	//fmt.Println(reflect.TypeOf(image.Header))
	fmt.Println("------------FILE------------")
	fmt.Printf("%v\n", file)
	fmt.Println(reflect.TypeOf(file))
	c.IndentedJSON(http.StatusCreated, "image name: "+imgName)

}
