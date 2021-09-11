package license

import (
	"fmt"
	"strings"
	"sync"

	"flexeracode/fileaccess/models"
)

// CalculateLicense method to calculate license for the given application
func (l licenseApplication) CalculateLicense(applicationId string) (numberOfCopies int, err error) {

	if applicationId == "" {
		return -1, fmt.Errorf("application id is not defined")
	}

	return l.aggregateResult(l.fileAccess.GetUsersCopiesByAppId(applicationId))
}

//aggregating result
func (l licenseApplication) aggregateResult(chanApp <-chan *models.Response) (appCount int, err error) {

	applicationByUser := map[string][]*models.UserApplication{}

	//waiting for channel
	for appItem := range chanApp {

		application, err := appItem.UserCopy, appItem.ErrorMessage
		if err != nil {
			return 0, err
		}
		applicationByUser[application.UserID] = append(applicationByUser[application.UserID], application)
	}
	sum := 0

	//aggregating license count through channel
	sumChan := l.counter(applicationByUser)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			sumVal, success := <-sumChan
			sum += sumVal
			if !success {
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
	return sum, nil
}

//count unique required license
func (l licenseApplication) counter(applications map[string][]*models.UserApplication) (sum chan int) {
	sum = make(chan int)
	go func(responseChn chan int) {
		for _, appItems := range applications {

			desktopMap := make(map[string]bool)
			laptopMap := make(map[string]bool)

			for _, item := range appItems {
				machineType := strings.ToUpper(item.ComputerType)
				key := fmt.Sprintf("%s-%s", item.UserID, item.ComputerID)
				switch machineType {
				case Desktop:
					desktopMap[key] = true
				default:
					laptopMap[key] = true
				}
			}
			if len(desktopMap) > len(laptopMap) {
				responseChn <- len(desktopMap)
			}
			responseChn <- len(laptopMap)
		}
		defer close(responseChn)
	}(sum)
	return sum
}
