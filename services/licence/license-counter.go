package licence

import (
	"fmt"
	"strings"
	"sync"
	"time"

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

	//aggregating licence count through channel
	sumChan := l.counter(applicationByUser)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			sumVal, more := <-sumChan
			sum += sumVal
			if !more {
				wg.Done()
				return
			}
			time.Sleep(1000 * time.Nanosecond)
		}
	}()
	wg.Wait()
	return sum, nil
}

//count unique required license
func (l licenseApplication) counter(applications map[string][]*models.UserApplication) (sum chan int) {
	sum = make(chan int)
	go func() {
		for _, appItems := range applications {

			desktopMap := make(map[string]bool)
			laptopMap := make(map[string]bool)

			for _, item := range appItems {

				machineType := strings.ToUpper(item.ComputerType)
				key := item.UserID + "-" + item.ComputerID
				switch machineType {
				case "DESKTOP":
					desktopMap[key] = true
				default:
					laptopMap[key] = true
				}
			}
			if len(desktopMap) > len(laptopMap) {
				sum <- len(desktopMap)
			}
			sum <- len(laptopMap)
		}
		close(sum)
	}()
	return sum
}
