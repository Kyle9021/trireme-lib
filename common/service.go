package common

import "github.com/aporeto-inc/trireme-lib/utils/portspec"

// Service is a protocol/port service of interest - used to pass user requests
type Service struct {
	// Ports are the corresponding ports
	Ports *portspec.PortSpec `json:"Ports,omitempty"`

	// Port is the service port. This has been deprecated and will be removed in later releases 01/13/2018
	Port uint16

	// Protocol is the protocol number
	Protocol uint8
}

// ConvertServicesToPortList converts an array of services to a port list
func ConvertServicesToPortList(services []Service) string {

	portlist := ""
	for _, s := range services {
		portlist = portlist + s.Ports.String() + ","
	}

	if len(portlist) == 0 {
		portlist = "0"
	} else {
		portlist = string(portlist[:len(portlist)-1])
	}

	return portlist
}