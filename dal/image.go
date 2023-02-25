package dal

import (
	"io/ioutil"
	"os"
)

func (fs *FileService) StoreFile(buffer []byte, fileName string) error {
	dst := fs.config.App.ImageDirPath + fileName + ".jpg"

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

func (fs *FileService) GetFile(fileName string) ([]byte, error) {
	filePath := fs.config.App.ImageDirPath + fileName + ".jpg"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (fs *FileService) GetAllFilesNames() ([]string, error) {
	files, err := ioutil.ReadDir(fs.config.App.ImageDirPath)
	if err != nil {
		return nil, err
	}

	retVal := []string{}
	for _, file := range files {
		retVal = append(retVal, file.Name())
	}

	return retVal, nil
}

func (fs *FileService) DeleteFile(fileName string) error {
	filePath := fs.config.App.ImageDirPath + fileName
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
