package fileaccess

import (
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

func Test_analyze(t *testing.T) {

	fileAccess := fileAccess{fileName: "fileName"}

	t.Run("Ope File with exiting file", func(t *testing.T) {
		if _, err := fileAccess.fileOpen(); (err != nil) != false {
			t.Errorf("analyze() error = %v", err)
		}
	})
}
