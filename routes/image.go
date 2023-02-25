package routes

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"sort"

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

	// check if image exists
	_, err = r.dal.GetFile(imgName)
	if !errors.Is(err, os.ErrNotExist) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

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

func (r *Router) getAll(c *gin.Context) {
	imgNames, err := r.dal.GetAllFilesNames()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	sort.Strings(imgNames)

	c.IndentedJSON(http.StatusOK, imgNames)
}

func (r *Router) delete(c *gin.Context) {
	fileName := c.Param("fileName")
	err := r.dal.DeleteFile(fileName)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, fileName+" deleted")
}

func (r *Router) download(c *gin.Context) {
	fileName := c.Param("fileName")
	filePath := r.config.App.ImageDirPath + fileName
	c.File(filePath)
}
