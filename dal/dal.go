package dal

type Dal interface {
	StoreFile(buffer []byte, fileName string) error
}

type FileService struct{}
