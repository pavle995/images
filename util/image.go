package util

import (
	"crypto/sha256"
	"fmt"
	"mime/multipart"
)

func GetImageName(buf []byte) string {
	h := sha256.New()
	h.Write(buf)

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func ReadImage(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	buf := make([]byte, file.Size)
	src.Read(buf)

	return buf, nil
}
