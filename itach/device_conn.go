package itach

import (
  "bufio"
  "errors"
  "fmt"
  "net"
  "strings"
  "strconv"
)

// DeviceConn encapsulates a TCP connection to
// an iTach device and provides API specific
// functions
type DeviceConn struct {
  net.Conn
}

type NetworkSettings struct {
  ConfigLock string
  IPSettings string
  IPAddr net.IP
  Subnet net.IPMask
  Gateway net.IP
}

// ReadResponse is a convenience function for reading
// carriage return delimited responses.
func (conn *DeviceConn) ReadResponse() (resp string, err error) {
  resp, err = bufio.NewReader(conn).ReadString('\r')

  if err == nil {
    // Strip carriage return
    resp = resp[0:len(resp)-2]
  }

  return resp, err
}

// SendCommand sends the string cmd to the device, and returns
// the response.
func (conn *DeviceConn) SendCommand(cmd string) (resp string, err error) {
  _, err = fmt.Fprintf(conn, "%v\r", cmd)
  if err != nil {
    return "", err
  }

  resp, err = conn.ReadResponse()
  if err != nil {
    return "", err
  }

  if strings.HasPrefix(resp, "ERR_") {
    resp = strings.TrimPrefix(resp, "ERR_")
    code, err := strconv.ParseInt(resp, 0, 0)
    if err != nil {
      return "", err
    }

    return "", ErrorForCode(code)
  }

  if strings.HasPrefix(resp, "unknowncommand,") {
    resp = strings.TrimPrefix(resp, "unknowncommand,")
    code, err := strconv.ParseInt(resp, 0, 0)
    if err != nil {
      return "", err
    }

    return "", ErrorForCode(code)
  }

  return resp, nil
}

// GetVersion returns the device's version string.
func (conn *DeviceConn) GetVersion() (string, error) {
  return conn.SendCommand("getversion")
}

// GetNetworkSettings returns the device's network settings
func (conn *DeviceConn) GetNetworkSettings() (s *NetworkSettings, err error) {
  resp, err := conn.SendCommand("get_NET,0:1")
  if err != nil {
    return nil, err
  }

  if !strings.HasPrefix(resp, "NET,0:1,") {
    return nil, errors.New("Invalid response from device")
  }

  split := strings.Split(strings.TrimPrefix(resp, "NET,0:1,"), ",")

  if len(split) != 5 {
    return nil, errors.New("Invalid number of network settings")
  }

  s = &NetworkSettings{ConfigLock: split[0], IPSettings: split[1]}

  s.IPAddr = net.ParseIP(split[2])
  if s.IPAddr == nil {
    return nil, errors.New("Invalid IP address")
  }

  s.Gateway = net.ParseIP(split[4])
  if s.Gateway == nil {
    return nil, errors.New("Invalid gateway")
  }

  mask := net.ParseIP(split[3])
  if mask == nil {
    return nil, errors.New("Invalid subnet")
  }

  s.Subnet = net.IPv4Mask(mask[0], mask[1], mask[2], mask[3])

  return s, nil
}
