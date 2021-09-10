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

func Test_When_Prepare_Should_return_Success_With_Object(t *testing.T) {

	fileAccess := fileAccess{fileName: "../sample-small-test.csv"}
	stringFile := []string{"839", "8809", "374", "laptop", "Exported from System A", "839", "8809"}

	entity, err := fileAccess.prepare(stringFile, nil)

	t.Run("should not contain error", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("UserId should be equal to 8809", func(t *testing.T) {
		assert.Equal(t, entity.UserID, "8809")
	})

	t.Run("ComputerId should be equal to 839", func(t *testing.T) {
		assert.Equal(t, entity.ComputerID, "839")
	})

	t.Run("ComputerType should be equal to 'laptop'", func(t *testing.T) {
		assert.Equal(t, entity.ComputerType, "laptop")
	})
}
