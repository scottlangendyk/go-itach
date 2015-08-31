package itach

import (
  "bytes"
  "errors"
  "fmt"
  "strconv"
  "strings"
)

// SendIR returns a string representing the IRCommand in the
// Global Cache irsend format.
func (c *IRCommand) SendIR(connectorAddress int) (s string, err error) {
  buf := new(bytes.Buffer)

  _, err = fmt.Fprintf(buf, "sendir,1:%d,0,%d,%d,%d", connectorAddress, c.Frequency, c.Repeat, c.Offset)
  if err != nil {
    return "", err
  }

  for _, pulse := range c.Pulses {
    _, err = fmt.Fprintf(buf, ",%d,%d", pulse.On, pulse.Off)
    if err != nil {
      return "", err
    }
  }

  return buf.String(), nil
}

// ParseSendIR parses a Global Cache formatted sendir string
// returns the result as an IRCommand.
//
// If the sendir string is incorrectly formatted or invalid
// an error will be returned
func ParseSendIR(raw string) (cmd *IRCommand, err error) {
  cmd = &IRCommand{}

  split := strings.Split(raw, ",")

  if len(split) < 10 || split[0] != "sendir" {
    return nil, errors.New("Invalid sendir command")
  }

  freq, err := strconv.ParseInt(split[3], 10, 0)
  if err != nil {
    return nil, err
  }

  cmd.Frequency = uint(freq)

  repeat, err := strconv.ParseInt(split[4], 10, 0)
  if err != nil {
    return nil, err
  }

  cmd.Repeat = uint(repeat)

  offset, err := strconv.ParseInt(split[5], 10, 0)
  if err != nil {
    return nil, err
  }

  cmd.Offset = uint(offset)

  p := split[6:]
  n := len(p)

  if n % 2 != 0 {
    return nil, errors.New("Uneven number of pulses")
  }

  cmd.Pulses = make([]IRPulse, n/2)

  for i := 0; i < n/2; i++ {
    pulse := IRPulse{}

    on, err := strconv.ParseInt(p[i*2], 10, 0)
    if err != nil {
      return nil, err
    }

    pulse.On = uint(on)

    off, err := strconv.ParseInt(p[i*2+1], 10, 0)
    if err != nil {
      return nil, err
    }

    pulse.Off = uint(off)

    cmd.Pulses[i] = pulse
  }

  return cmd, nil
}
