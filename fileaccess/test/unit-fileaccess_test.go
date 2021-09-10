package test

import (
	"fmt"
	"testing"

	"flexeracode/fileaccess"
)

func Test_When_GetUsersCopiesByAppId_Error_WhenReader_Error(t *testing.T) {

	fileAccess := fileaccess.New("")
	retChannel := fileAccess.GetUsersCopiesByAppId("374")

	for item := range retChannel {
		fmt.Print(item.ErrorMessage)
	}

}
