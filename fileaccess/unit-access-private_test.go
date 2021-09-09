package fileaccess

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FileOPen_Should_Return_Error_If_FileNotFound(t *testing.T) {
	fileAccess := fileAccess{
		fileName: "fileName",
	}
	_, err := fileAccess.fileOpen()
	assert.Error(t, err)
}

func Test_FileOPen_Should_Return_Error_If_FileSuccess(t *testing.T) {

	fileAccess := fileAccess{fileName: "../sample-small-test.csv"}

	t.Run("Ope File with exiting file", func(t *testing.T) {
		_, err := fileAccess.fileOpen()
		assert.NoError(t, err)
	})
}

func Test_When_Prepare_Should_return_Error_When_File_With_Error(t *testing.T) {

	fileAccess := fileAccess{fileName: "../sample-small-test.csv"}
	stringFile := []string{"839", "8809", "374", "laptop,Exported from System A", "839", "8809"}

	_, err := fileAccess.prepare(stringFile, fmt.Errorf("error in file"))
	assert.Error(t, err)

}
