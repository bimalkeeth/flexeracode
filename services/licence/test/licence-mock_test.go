package test

import (
	"flexeracode/fileaccess"
	"flexeracode/fileaccess/models"
)

// mock object for file access
type MockFleAccess struct {
	fileaccess.IFileAccess
}

var mockFleAccess MockFleAccess

func init() {
	mockFleAccess = MockFleAccess{}
}

//getLicenceData
func getLicenceData() []*models.UserApplication {
	return []*models.UserApplication{
		{
			ComputerID:    "839",
			UserID:        "8809",
			ApplicationID: "374",
			ComputerType:  "laptop",
			Comment:       "Exported from System A",
		}, {
			ComputerID:    "937",
			UserID:        "9023",
			ApplicationID: "374",
			ComputerType:  "desktop",
			Comment:       "Exported from System A",
		}, {
			ComputerID:    "944",
			UserID:        "7823",
			ApplicationID: "374",
			ComputerType:  "laptop",
			Comment:       "Exported from System A",
		}, {
			ComputerID:    "2628",
			UserID:        "8076",
			ApplicationID: "374",
			ComputerType:  "LAPTOP",
			Comment:       "Exported from System A",
		}, {
			ComputerID:    "3806",
			UserID:        "5261",
			ApplicationID: "374",
			ComputerType:  "DESKTOP",
			Comment:       "Exported from System A",
		},
	}
}

//GetUsersCopiesByAppId mock function for file access
func (a MockFleAccess) GetUsersCopiesByAppId(appId string) <-chan *models.Response {

	brodCastChan := make(chan *models.Response)
	//sending read output to the caller through channel
	go func(broadChan chan *models.Response) {
		items := getLicenceData()
		for _, item := range items {
			broadChan <- &models.Response{UserCopy: item, ErrorMessage: nil}
		}
		close(broadChan)
	}(brodCastChan)
	return brodCastChan
}
