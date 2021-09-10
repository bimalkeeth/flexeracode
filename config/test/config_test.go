package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"flexeracode/config"
)

func Test_When_LicenseService_Should_Return_ILicenseApplication(t *testing.T) {
	configService := config.NewConfig("sample-small-test.csv")
	licenseService := configService.LicenseService()
	assert.IsType(t, fmt.Sprintf("%T\n", licenseService), "license.ILicenseApplication", "this is not correct")

}

func Test_When_FileAccess_Should_Return_IFileAccess(t *testing.T) {
	configService := config.NewConfig("sample-small-test.csv")
	licenseService := configService.LicenseService()
	assert.IsType(t, fmt.Sprintf("%T\n", licenseService), "fileaccess.IFileAccess", "this is not correct")

}
