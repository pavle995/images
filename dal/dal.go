package dal

import (
	"github.com/pavle995/images/config"
)

type Dal interface {
	StoreFile(buffer []byte, fileName, extension string) error
	GetFile(fileName, extension string) ([]byte, error)
	GetAllFilesNames() ([]string, error)
	DeleteFile(fileName string) error
}

type FileService struct {
	config *config.Config
}

func NewFileService(config *config.Config) *FileService {
	return &FileService{config: config}
}
