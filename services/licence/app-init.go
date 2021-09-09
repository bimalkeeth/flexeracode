package licence

import (
	"flexeracode/fileaccess"
)

// ILicenseApplication Interface abstraction to  calculate applications
type ILicenseApplication interface {

	// CalculateLicense calculate number of copies for licensing
	CalculateLicense(applicationId string) (numberOfCopies int, err error)
}

type licenseApplication struct {
	fileAccess fileaccess.IFileAccess
}

// New constructor for license service
func New(access fileaccess.IFileAccess) ILicenseApplication {
	return licenseApplication{fileAccess: access}
}
