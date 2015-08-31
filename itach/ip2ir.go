package itach

import (
  "errors"
  "fmt"
  "strings"
)

// IRPulse is a single on/off IR pulse.
type IRPulse struct {
  On uint
  Off uint
}

type IRCommand struct {
  Frequency uint
  Pulses []IRPulse
  Repeat uint
  Offset uint
}

// IRLearner encapsulates IR learning capabilities
// for IP2IR devices.
type IRLearner struct {
  conn *DeviceConn
}

// Close closes the IRLearner.
func (l *IRLearner) Close() error {
  defer l.conn.Close()

  _, err := l.conn.SendCommand("stop_IRL")
  if err != nil {
    return err
  }

  return nil
}

// Accept waits for and returns the next learned IRCommand.
func (l *IRLearner) Accept() (cmd *IRCommand, err error) {
  resp, err := l.conn.ReadResponse()
  if err != nil {
    return nil, err
  }

  cmd, err = ParseSendIR(resp)
  if err != nil {
    return nil, err
  }

  return cmd, nil
}

// GetIRLearner puts the IP2IR device in learning mode and
// returns an IRLearner.
func (d *Device) GetIRLearner() (l *IRLearner, err error) {
 l = &IRLearner{}

 l.conn, err = d.Connect()
 if err != nil {
   return nil, err
 }

 resp, err := l.conn.SendCommand("get_IRL")
 if err != nil {
    l.conn.Close()
    return nil, err
  }

  if resp != "IR Learner Enabled" || resp == "IR Learner Unavailable" {
    l.conn.Close()
    return nil, errors.New("Couldn't enable IR Learner")
  }

 return l, nil
}

// StopIR halts IR transmission for a given conenctorAddress on an
// IP2IR device.
func (conn *DeviceConn) StopIR(connectorAddress int) error {
  _, err := conn.SendCommand(fmt.Sprintf("stopir,1:%d", connectorAddress))

  return err
}

// SetIRMode sets the ir mode for a given connectorAddress on an
// IP2IR device.
func (conn *DeviceConn) SetIRMode(connectorAddress int, mode string) error {
  _, err := conn.SendCommand(fmt.Sprintf("set_IR,1:%d,%v", connectorAddress, mode))

  return err
}

// GetIRMode returns a string with current ir mode for a given
// connectorAddress on an IP2IR device.
func (conn *DeviceConn) GetIRMode(connectorAddress int) (string, error) {
  resp, err := conn.SendCommand(fmt.Sprintf("get_IR,1:%d", connectorAddress))
  if err != nil {
    return "", err
  }

  return resp[7:len(resp)-1], nil
}

// SendIR sends an IRCommand to the specified connectorAddress on an
// IP2IR device.
func (conn *DeviceConn) SendIR(connectorAddress int, command *IRCommand) error {
  s, err := command.SendIR(connectorAddress)
  if err != nil {
    return err
  }

  resp, err := conn.SendCommand(s)
  if err != nil {
    return err
  }

  if strings.Contains(resp, "busyir") {
    return errors.New("IR connector is busy")
  }

  return nil
}
