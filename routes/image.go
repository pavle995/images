package routes

import (
	"errors"
	"net/http"
	"os"
	"sort"

	"github.com/pavle995/images/util"

	"github.com/gin-gonic/gin"
)

func (r *Router) uploadImage(c *gin.Context) {
	_, image, err := c.Request.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	buffer, err := util.ReadImage(image)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "server failed to read data from file")
		return
	}
	imgName := util.GetImageName(buffer)
	ext := util.GetFileExtension(image.Filename)

	// check if image exists
	_, err = r.dal.GetFile(imgName, ext)
	if !errors.Is(err, os.ErrNotExist) {
		c.AbortWithStatusJSON(http.StatusBadRequest, "image already exists")
		return
	}

	err = r.dal.StoreFile(buffer, imgName, ext)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "server failed to store image")
		return
	}

	c.IndentedJSON(http.StatusCreated, "image name: "+imgName)

}

func (r *Router) getAll(c *gin.Context) {
	imgNames, err := r.dal.GetAllFilesNames()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "server failed to fetch images")
		return
	}
	sort.Strings(imgNames)

	c.IndentedJSON(http.StatusOK, imgNames)
}

func (r *Router) delete(c *gin.Context) {
	fileName := c.Param("fileName")
	err := r.dal.DeleteFile(fileName)
	if os.IsNotExist(err) {
		c.AbortWithStatusJSON(http.StatusNotFound, fileName+" not exists")
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "server failed to delete image")
		return
	}

	c.IndentedJSON(http.StatusOK, fileName+" deleted")
}

func (r *Router) download(c *gin.Context) {
	fileName := c.Param("fileName")
	filePath := r.config.App.ImageDirPath + fileName
	c.File(filePath)
}
