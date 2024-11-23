package gopen

import (
	"fmt"
	"testing"

	"github.com/david-gurley/gopen/pds"
)

var DSC_CLIENT_ADDRESSES = []string{
	"172.31.7.23",
	"172.31.7.24",
}

var (
	CLIENT_PORT = 11357
)

func NoTestGetDeviceStatus(t *testing.T) {
	for _, clientAddress := range DSC_CLIENT_ADDRESSES {
		c, err := NewDSCClient(clientAddress, CLIENT_PORT)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("successfully connected to gRPC server on DSC: %s\n", clientAddress)
		status, err := c.GetDeviceStatus()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("device status: %v\n", status)
		c.Close()
	}
}

func NoTestUpdateDeviceProfile(t *testing.T) {
	for _, clientAddress := range DSC_CLIENT_ADDRESSES {
		c, err := NewDSCClient(clientAddress, CLIENT_PORT)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("successfully connected to gRPC server on DSC: %s\n", clientAddress)
		err = c.UpdateDeviceProfile("DEVICE_PROFILE_64VF")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("succesfully updated the device profile")
		c.Close()
	}
}

func NoTestGetDeviceProfile(t *testing.T) {
	for _, clientAddress := range DSC_CLIENT_ADDRESSES {
		c, err := NewDSCClient(clientAddress, CLIENT_PORT)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("successfully connected to gRPC server on DSC: %s\n", clientAddress)
		profile, err := c.GetDeviceProfile()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("device profile: %v\n", profile)
		c.Close()
	}
}
func NoTestGetLifs(t *testing.T) {
	for _, clientAddress := range DSC_CLIENT_ADDRESSES {
		c, err := NewDSCClient(clientAddress, CLIENT_PORT)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("successfully connected to gRPC server on DSC: %s\n", clientAddress)
		lifs, err := c.GetLifs()
		if err != nil {
			fmt.Println(err)
		}
		for _, lif := range lifs {
			var lifType pds.LifType = pds.LifType(pds.LifType_value["LIF_TYPE_HOST_VF"])
			if lif.Spec.Type == lifType {
				fmt.Printf("id: %v type: %v MAC: %v\n", lif.Spec.Id, lif.Spec.Type, lif.Spec.MacAddress)
			}
		}
		c.Close()
	}
}

func NoTestGetInterfaces(t *testing.T) {
	for _, clientAddress := range DSC_CLIENT_ADDRESSES {
		c, err := NewDSCClient(clientAddress, CLIENT_PORT)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("successfully connected to gRPC server on DSC: %s\n", clientAddress)
		interfaces, err := c.GetInterfaces()
		if err != nil {
			fmt.Println(err)
		}
		for _, intf := range interfaces {
			switch i := intf.Spec.Ifinfo.(type) {
			case *pds.InterfaceSpec_HostIfSpec:
				s := intf.Status.Ifstatus.(*pds.InterfaceStatus_HostIfStatus)
				fmt.Printf("Vf: %v MAC: %#v Name: %v Name: %#v\n", i.HostIfSpec.Vf, i.HostIfSpec.MACAddress, i.HostIfSpec.Name, s.HostIfStatus.Name)
			default:
				fmt.Println("does not match")
			}
		}
		c.Close()
	}
}
