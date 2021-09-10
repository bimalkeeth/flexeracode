package fileaccess

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"flexeracode/fileaccess/models"
)

//open file data store to read
func (a *fileAccess) fileOpen() (*os.File, error) {
	applicationDataFile, err := os.OpenFile(a.fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Printf("unable to open file: %v", err)
		return nil, err
	}
	return applicationDataFile, nil
}

//prepare current row of date into application data entity
func (a *fileAccess) prepare(line []string, errResponse error) (application *models.UserApplication, err error) {

	if errResponse != nil {
		return nil, errResponse
	}
	return &models.UserApplication{
		ComputerID:    line[ComputerID],
		UserID:        line[UserID],
		ApplicationID: line[ApplicationID],
		ComputerType:  line[ComputerType],
		Comment:       line[Comment]}, nil
}

// read file function open file and initiate file sv file reader
func (a *fileAccess) readFile() (err error) {
	a.openedFile, err = a.fileOpen()
	if err != nil {
		return err
	}
	a.reader = csv.NewReader(a.openedFile)
	_, err = a.reader.Read()
	return err
}

// readLine read current line from the file
func (a *fileAccess) readLine() (line *models.UserApplication, err error) {
	return a.prepare(a.reader.Read())
}

func (a *fileAccess) fileCloser() {
	_ = a.openedFile.Close()
}

// GetUsersCopiesByAppId get data from file data store by application id
func (a *fileAccess) GetUsersCopiesByAppId(appId string) <-chan *models.Response {

	brodCastChan := make(chan *models.Response)
	err := a.readFile()
	if err != nil {
		a.fileCloser()
		brodCastChan <- &models.Response{UserCopy: nil, ErrorMessage: err}
		close(brodCastChan)
	}
	//sending read output to the caller through channel
	go func(broadChan chan *models.Response) {

		defer a.fileCloser()

		line, err := a.readLine()
		for err == nil {
			if appId == line.ApplicationID {
				broadChan <- &models.Response{UserCopy: line, ErrorMessage: nil}
			}
			line, err = a.readLine()
		}
		if !strings.Contains(err.Error(), "EOF") {
			broadChan <- &models.Response{UserCopy: nil, ErrorMessage: err}
		}
		close(broadChan)
	}(brodCastChan)

	return brodCastChan
}
