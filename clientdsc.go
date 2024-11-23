package gopen

import (
	"fmt"

	"github.com/david-gurley/gopen/pds"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type DSCClient struct {
	gRPCconn *grpc.ClientConn
}

func NewDSCClient(address string, port int) (*DSCClient, error) {
	client := DSCClient{}
	gRPCconn, err := grpc.Dial(fmt.Sprintf("%s:%v", address, port), grpc.WithInsecure())
	if err != nil {
		return &client, err
	}
	client.gRPCconn = gRPCconn
	return &client, nil
}

func (client *DSCClient) Close() error {
	client.gRPCconn.Close()
	return nil
}

func (client *DSCClient) GetDevice() (*pds.Device, error) {
	dc := pds.NewDeviceSvcClient(client.gRPCconn)
	deviceResponse, err := dc.DeviceGet(context.Background(), &pds.DeviceGetRequest{})
	if err != nil {
		return &pds.Device{}, err
	}
	return deviceResponse.Response, nil
}

func (client *DSCClient) GetDeviceStatus() (*pds.DeviceStatus, error) {
	device, err := client.GetDevice()
	if err != nil {
		return &pds.DeviceStatus{}, err
	}
	return device.Status, nil
}

func (client *DSCClient) GetDeviceSpec() (*pds.DeviceSpec, error) {
	device, err := client.GetDevice()
	if err != nil {
		return &pds.DeviceSpec{}, err
	}
	return device.Spec, nil
}

func (client *DSCClient) GetDeviceProfile() (*pds.DeviceProfile, error) {
	device, err := client.GetDevice()
	if err != nil {
		return nil, err
	}
	return &device.Spec.DeviceProfile, nil
}

func (client *DSCClient) UpdateDevice(spec *pds.DeviceSpec) error {
	dc := pds.NewDeviceSvcClient(client.gRPCconn)
	deviceRequest := pds.DeviceRequest{}
	deviceRequest.Request = spec
	_, err := dc.DeviceUpdate(context.Background(), &deviceRequest)
	if err != nil {
		return err
	}
	return nil
}

func (client *DSCClient) UpdateDeviceProfile(profileName string) error {
	deviceSpec, err := client.GetDeviceSpec()
	if err != nil {
		return err
	}
	//var deviceProfile pds.DeviceProfile = pds.DeviceProfile(pds.DeviceProfile_value[profileName])
	deviceSpec.DeviceProfile = pds.DeviceProfile(pds.DeviceProfile_value[profileName])
	return client.UpdateDevice(deviceSpec)
}

func (client *DSCClient) GetLifs() ([]*pds.Lif, error) {
	lc := pds.NewIfSvcClient(client.gRPCconn)
	lifResponse, err := lc.LifGet(context.Background(), &pds.LifGetRequest{})
	if err != nil {
		return lifResponse.Response, err
	}
	return lifResponse.Response, nil
}

func (client *DSCClient) GetInterfaces() ([]*pds.Interface, error) {
	lc := pds.NewIfSvcClient(client.gRPCconn)
	interfacesResponse, err := lc.InterfaceGet(context.Background(), &pds.InterfaceGetRequest{})
	if err != nil {
		return interfacesResponse.Response, err
	}
	return interfacesResponse.Response, nil
}
