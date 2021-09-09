package config

import (
	"flexeracode/fileaccess"
	"flexeracode/services/licence"
)

type IApplicationConfig interface {
	LicenseService() licence.ILicenseApplication
	FileAccess() fileaccess.IFileAccess
}

type appConfig struct {
	fileName string
}

// LicenseService getting instance of the license server
func (a appConfig) LicenseService() licence.ILicenseApplication {
	return licence.New(a.FileAccess())
}

// FileAccess getting instance of file access service
func (a appConfig) FileAccess() fileaccess.IFileAccess {
	return fileaccess.New(a.fileName)
}

func NewConfig(fileName string) IApplicationConfig {
	return &appConfig{fileName: fileName}
}
