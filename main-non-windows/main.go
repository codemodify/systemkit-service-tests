// +build !windows

package main

import (
	service "github.com/codemodify/systemkit-service"
	spec "github.com/codemodify/systemkit-service-spec"
)

func main() {
	myService := spec.NewEmptySERVICE()
	myService.Name = "abracadabra-service"
	myService.Executable = "abracadabra-service"

	s := service.NewServiceFromSERVICE(myService)
	s.Install()
}
