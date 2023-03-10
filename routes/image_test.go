package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pavle995/images/config"
	"github.com/pavle995/images/dal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpload(t *testing.T) {
	mockDal := &dal.MockDal{}
	cfg := config.GetConfig(true)
	r := NewRouter(mockDal, cfg)

	mockDal.On("GetFile", mock.Anything, mock.Anything).Return([]byte{1, 2, 3}, os.ErrNotExist)
	mockDal.On("StoreFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	filePath := "../test/gopher.jpg"
	body := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(body)
	fileHeader := make(textproto.MIMEHeader)
	fileHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "image", "gopher.jpg"))
	fileHeader.Set("Content-Type", "text/plain")
	writer, _ := multipartWriter.CreatePart(fileHeader)
	file, _ := os.Open(filePath)
	io.Copy(writer, file)
	multipartWriter.Close()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = httptest.NewRequest(http.MethodPost, "/image", body)
	ctx.Request.Header.Add("Content-Type", multipartWriter.FormDataContentType())

	r.uploadImage(ctx)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAll(t *testing.T) {
	mockDal := &dal.MockDal{}
	cfg := config.GetConfig(true)
	r := NewRouter(mockDal, cfg)

	fileNames := []string{"test.jpg", "test2.jpg"}

	mockDal.On("GetAllFilesNames", mock.Anything, mock.Anything).Return(fileNames, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	r.getAll(ctx)

	returnedNames := make([]string, 2)
	json.Unmarshal(w.Body.Bytes(), &returnedNames)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fileNames, returnedNames)
}

func TestDelete(t *testing.T) {
	mockDal := &dal.MockDal{}
	cfg := config.GetConfig(true)
	r := NewRouter(mockDal, cfg)

	mockDal.On("DeleteFile", mock.Anything).Return(nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	r.delete(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDownload(t *testing.T) {
	mockDal := &dal.MockDal{}
	cfg := config.GetConfig(true)
	r := NewRouter(mockDal, cfg)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/image/gopher.jph", nil)

	ctx.AddParam("fileName", "gopher.jpg")

	r.download(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
}
