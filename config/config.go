package config

import (
	"flexeracode/fileaccess"
	"flexeracode/services/license"
)

type IApplicationConfig interface {
	LicenseService() license.ILicenseApplication
	FileAccess() fileaccess.IFileAccess
}

type appConfig struct {
	fileName string
}

// LicenseService getting instance of the license server
func (a appConfig) LicenseService() license.ILicenseApplication {
	return license.New(a.FileAccess())
}

// FileAccess getting instance of file access service
func (a appConfig) FileAccess() fileaccess.IFileAccess {
	return fileaccess.New(a.fileName)
}

func NewConfig(fileName string) IApplicationConfig {
	return &appConfig{fileName: fileName}
}
