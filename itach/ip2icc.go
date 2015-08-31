package itach

import (
  "fmt"
  "strconv"
)

// SetState sets the boolean state of an output on an IP2CC
// device.
func (conn *DeviceConn) SetState(output int, state bool) error {
  _, err := conn.SendCommand(fmt.Sprintf("setstate,1:%d,%d", output, state))

  return err
}

// GetState returns the state of an output on an IP2CC device.
func (conn *DeviceConn) GetState(output int) (state bool, err error) {
  resp, err := conn.SendCommand(fmt.Sprintf("getstate,1:%d", output))
  if err != nil {
    return state, err
  }

  return strconv.ParseBool(string(resp[10]))
}
