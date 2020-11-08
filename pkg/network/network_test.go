package network

import (
	"fmt"
	"testing"
)

func TestAvailableNetworkPort(t *testing.T) {
	port, err := ObtainAvailablePort(20000, 30000)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("port => ", port)
	}
}

func TestGetOutboundIP(t *testing.T) {
	ip := GetOutboundIP()
	fmt.Println("\noutbound ip => ", ip.String())
}
