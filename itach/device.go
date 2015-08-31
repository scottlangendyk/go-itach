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

// ListenToSensors listens for sensor notification
// from an IP2IR device on the default port.
func (d *Device) ListenToSensors() (l *SensorListener, err error) {
  return d.ListenToSensorsOnPort(9132)
}

// ListenToSensors listens for sensor notifications
// from an IP2IR device on the specified port.
func (d *Device) ListenToSensorsOnPort(port int) (l *SensorListener, err error) {
  addr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("239.255.250.250:%d", port))
  if err != nil {
    return nil, err
  }

  conn, err := net.ListenMulticastUDP("udp4", nil, addr)
  if err != nil {
    return nil, err
  }

  return &SensorListener{conn: &DeviceConn{conn}}, nil
}
