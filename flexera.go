package main

import (
	"fmt"

	"flexeracode/config"
)

func main() {

	configService := config.NewConfig("/Users/bimalkeeth/Downloads/Flexera_Code_Test-2/sample-small.csv")
	license := configService.LicenseService()
	_, err := license.CalculateLicense("374")
	fmt.Print(err)

}
