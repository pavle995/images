package dal

import (
	"os"

	"github.com/pavle995/images/config"
)

func (fs *FileService) StoreFile(buffer []byte, fileName string) error {
	cfg := config.GetConfig()
	dst := cfg.App.ImageDirPath + fileName + ".jpg"

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}
