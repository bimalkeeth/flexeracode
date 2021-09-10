package main

import (
	"flag"
	"fmt"
	"log"

	"flexeracode/config"
)

func main() {

	fmt.Print("\033[H\033[2J")

	fmt.Println(`
(  __)(  )  (  __)( \/ )  (  )  (  )/ __)(  __)(  ( \/ ___)(  __)
 ) _) / (_/\ ) _)  )  (   / (_/\ )(( (__  ) _) /    /\___ \ ) _) 
(__)  \____/(____)(_/\_)  \____/(__)\___)(____)\_)__)(____/(____)
     `)

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
		log.Printf("error in clculating application license %v", err)
		return
	}

	colorYellow := "\033[33m"
	log.Println(string(colorYellow))

	log.Printf("Required license this application Number %v is %v", *applicationId, licenseCount)

}
