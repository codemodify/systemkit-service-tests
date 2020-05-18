// +build windows

package main

import (
	"fmt"
	"os"

	service "github.com/codemodify/systemkit-service"
	spec "github.com/codemodify/systemkit-service-spec"
)

func main() {

	// install ourselves as service
	if len(os.Args) > 1 {
		execPath, err := os.Executable()
		if err != nil {
			fmt.Println("Could not get executable path!")
			return
		}

		myService := spec.NewEmptySERVICE()
		myService.Name = "abracadabra-service"
		myService.Executable = execPath

		s := service.NewServiceFromSERVICE(myService)

		if os.Args[1] == "install" {
			s.Install()
		} else if os.Args[1] == "uninstall" {
			s.Uninstall()
		}

		return
	}

	// GETTING here means we are runnning in service mode

	// run the service loop inside our app, required by the Win SVC
	myService := spec.NewEmptySERVICE()
	myService.Name = "abracadabra-service"

	winServiceQueryHandler := service.NewWindowsServiceQueryHandler(&myService)
	winServiceQueryHandler.RunServiceQueryLoop()
}
