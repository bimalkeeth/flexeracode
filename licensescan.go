package main

import (
	"flag"
	"fmt"
	"log"

	"flexeracode/config"
)

func main() {

	fmt.Print("\033[H\033[2J")

	fmt.Println(`___________.__                  .____    .__                                   
\_   _____/|  |   ____ ___  ___ |    |   |__| ____  ____   ____   ______ ____  
 |    __)  |  | _/ __ \\  \/  / |    |   |  |/ ___\/ __ \ /    \ /  ___// __ \ 
 |     \   |  |_\  ___/ >    <  |    |___|  \  \__\  ___/|   |  \\___ \\  ___/ 
 \___  /   |____/\___  >__/\_ \ |_______ \__|\___  >___  >___|  /____  >\___  >
     \/              \/      \/         \/       \/    \/     \/     \/     \/ `)

	filename := flag.String("file", "sample-small-test.csv", "filename")
	applicationId := flag.String("app-id", "374", "application id")

	flag.Parse()

	if *filename == "" || *applicationId == "" {
		flag.Usage()
		return
	}

	configService := config.NewConfig(*filename)

	license := configService.LicenseService()
	licenseCount, err := license.CalculateLicense(*applicationId)

	if err != nil {
		log.Printf("error in clculating application licence %v", err)
		return
	}

	colorYellow := "\033[33m"
	log.Println(string(colorYellow))

	log.Printf("Required licence this application Number %v is %v", *applicationId, licenseCount)

}
