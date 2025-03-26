package network

import (
	"fmt"
	"net"
)

func GetMACAddress(ifaceName string) (net.HardwareAddr, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range interfaces {
		if iface.Name == ifaceName {
			return iface.HardwareAddr, nil
		}
	}
	return nil, fmt.Errorf("interface %s not found", ifaceName)
}
