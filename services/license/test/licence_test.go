package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"flexeracode/services/license"
)

func Test_When_CalculateLicense_Should_Return_Error_When_ApplicationId_Empty(t *testing.T) {
	licenceService := license.New(mockFleAccess)
	_, err := licenceService.CalculateLicense("")
	assert.Error(t, err)
}

func Test_When_CalculateLicense_Should_Licence_Counts(t *testing.T) {
	licenceService := license.New(mockFleAccess)
	licenceCount, err := licenceService.CalculateLicense("374")
	assert.NoError(t, err)
	assert.Equal(t, licenceCount, 5)

}
