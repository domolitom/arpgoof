package main

import (
	"github.com/domolitom/arpgoof/internal/network"
)

func main() {
	interfaceName := "eth0"
	localMAC, err := network.GetMACAddress(interfaceName)
	if err != nil {
		panic(err)
	}
	println(localMAC.String())
}
