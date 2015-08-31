package itach

import (
  "strconv"
  "strings"
)

type SensorListener struct {
  conn *DeviceConn
}

// Accept waits for and returns the next sensor notification.
func (l *SensorListener) Accept() (connectorAddress string, state bool, err error) {
  // Keep looping until we get a notification message
  for {
    resp, err := l.conn.ReadResponse()
    if err != nil {
      return "", false, err
    }

    if strings.HasPrefix(resp, "sensornotify,") {
      resp = strings.TrimPrefix(resp, "sensornotify,")
      split := strings.Split(resp, ",")

      if len(split) >= 2 {
        state, err = strconv.ParseBool(split[1])
        if err == nil {
          return split[0], state, nil
        }
      }
    }
  }
}

// Close closes the listener.
func (l *SensorListener) Close() error {
  return l.conn.Close()
}
