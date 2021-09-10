package license

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"flexeracode/fileaccess/models"
)

//getLicenceData
func getLicenceData() map[string][]*models.UserApplication {
	return map[string][]*models.UserApplication{
		"8809": {{
			ComputerID:    "839",
			UserID:        "8809",
			ApplicationID: "374",
			ComputerType:  "laptop",
			Comment:       "Exported from System A",
		}},
		"9023": {{
			ComputerID:    "937",
			UserID:        "9023",
			ApplicationID: "374",
			ComputerType:  "desktop",
			Comment:       "Exported from System A",
		}},
		"7823": {{
			ComputerID:    "944",
			UserID:        "7823",
			ApplicationID: "374",
			ComputerType:  "laptop",
			Comment:       "Exported from System A",
		}},
		"8076": {{
			ComputerID:    "2628",
			UserID:        "8076",
			ApplicationID: "374",
			ComputerType:  "LAPTOP",
			Comment:       "Exported from System A",
		}},
		"5261": {{
			ComputerID:    "3806",
			UserID:        "5261",
			ApplicationID: "374",
			ComputerType:  "DESKTOP",
			Comment:       "Exported from System A",
		}},
	}
}

func getUsersCopiesByAppIdForError(appId string) <-chan *models.Response {

	brodCastChan := make(chan *models.Response)
	//sending read output to the caller through channel
	go func(broadChan chan *models.Response) {
		items := []*models.UserApplication{
			{
				ComputerID:    "839",
				UserID:        "8809",
				ApplicationID: "374",
				ComputerType:  "laptop",
				Comment:       "Exported from System A",
			},
		}
		for _, item := range items {
			broadChan <- &models.Response{UserCopy: item, ErrorMessage: fmt.Errorf("error is defined")}
		}
		close(broadChan)
	}(brodCastChan)
	return brodCastChan
}

func Test_AggregateResult_When_Error_InResponse(t *testing.T) {
	licenseApplication := &licenseApplication{}
	_, err := licenseApplication.aggregateResult(getUsersCopiesByAppIdForError(""))
	assert.Error(t, err)
}

func Test_Counter_When_Success_InResponse(t *testing.T) {
	licenseApplication := &licenseApplication{}

	sumVal := 0
	chanResult := licenseApplication.counter(getLicenceData())
	for item := range chanResult {
		sumVal += item
	}
	assert.Equal(t, sumVal, 5)
}
