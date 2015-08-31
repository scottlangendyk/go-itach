package itach

import (
  "errors"
  "fmt"
  "io"
  "net"
  "strconv"
  "strings"
)

// ConnectSerial creates a TCP connection for communicating
// with an IP2SL serial port.
func (d *Device) ConnectSerial() (io.ReadWriteCloser, error) {
  return net.Dial("tcp", fmt.Sprintf("%v:4999", d.IP))
}

// SetSerial sets the baud rate, flow control, and parity settings for the
// serial port on an IP2SL device.
func (conn *DeviceConn) SetSerial(baud int, flow string, parity string) error {
  _, err := conn.SendCommand(fmt.Sprintf("set_SERIAL,1:1, %d, %v,%v", baud, flow, parity))

  return err
}

// GetSerial returns the baud rate, flow control, and parity settings for the
// serial port on an IP2SL device.
func (conn *DeviceConn) GetSerial() (baud int, flow string, parity string, err error) {
  resp, err := conn.SendCommand("get_SERIAL,1:1")
  if err != nil {
    return 0, "", "", err
  }

  if resp[0:11] != "SERIAL,1:1," {
    return 0, "", "", errors.New("Invalid response")
  }

  split := strings.Split(resp[11:], ",")

  if len(split) != 3 {
    return 0, "", "", errors.New("Invalid response")
  }

  b, err := strconv.ParseInt(split[0], 10, 0)
  if err != nil {
    return 0, "", "", err
  }

  return int(b), split[1], split[2], nil;
}
