package itach

import (
  "fmt"
  "net"
)

// Device represents an iTach device
type Device struct {
  UUID string
  Model string
  IP net.IP
}

// Connect returns a connection to the device.
func (d *Device) Connect() (*DeviceConn, error) {
  conn, err := net.Dial("tcp", fmt.Sprintf("%v:4998", d.IP))
  if err != nil {
    return nil, err
  }

  return &DeviceConn{conn}, nil
}
