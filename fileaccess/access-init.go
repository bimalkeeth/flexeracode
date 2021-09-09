package fileaccess

import (
	"encoding/csv"
	"os"

	"flexeracode/fileaccess/models"
)

// IFileAccess interface abstraction for file access
type IFileAccess interface {
	GetUsersCopiesByAppId(appId string) (response <-chan *models.Response)
}

//init struct for file access
type fileAccess struct {
	fileName   string
	reader     *csv.Reader
	openedFile *os.File
}

// New constructor for open file
func New(fileName string) IFileAccess {
	return &fileAccess{fileName: fileName}
}
