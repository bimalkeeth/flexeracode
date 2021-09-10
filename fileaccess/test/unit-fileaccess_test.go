package test

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"flexeracode/fileaccess"
	"flexeracode/fileaccess/models"
)

func Test_When_GetUsersCopiesByAppId_Error_WhenReader_Error(t *testing.T) {

	fileAccess := fileaccess.New("")
	retChannel := fileAccess.GetUsersCopiesByAppId("374")

	select {
	case item := <-retChannel:
		assert.Error(t, item.ErrorMessage)
		assert.Contains(t, item.ErrorMessage.Error(), "open : no such file or directory")
	}
}

func Test_When_ReadFile_Success_Should_Return_License_Entities(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")

	configPath, _ := filepath.Abs(fmt.Sprintf("%v/%v", dir, "sample-small-test.csv"))
	var responseList []*models.UserApplication

	fileAccess := fileaccess.New(configPath)
	retChannel := fileAccess.GetUsersCopiesByAppId("374")

	for item := range retChannel {
		assert.NoError(t, item.ErrorMessage)
		responseList = append(responseList, item.UserCopy)
	}

	assert.Greater(t, len(responseList), 0)

}
