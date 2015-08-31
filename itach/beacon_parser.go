package itach

import (
  "bufio"
  "strings"
)

// ParseBeacon parses a Global Cache beacon and returns
// a Device.
//
// The beacon is expected to be in the following format

// AMXB<-UUID=GlobalCache_000C1E024239><-SDKClass=Utility>
// <-Make=GlobalCache><-Model=iTachWF2IR><-Revision=710-1001-05>
// <-Pkg_Level=GCPK001><-Config-URL=http://192.168.1.100.><-PCB_PN=025-0026-06>
// <-Status=Ready>
func ParseBeacon(beacon string) (d *Device, err error) {
  r := bufio.NewReader(strings.NewReader(beacon))
  d = &Device{}

  for d.Model == "" || d.UUID == "" {
    _, err := r.ReadString('<')
    if err != nil {
      return nil, err
    }

    key, err := r.ReadString('=')
    if err != nil {
      return nil, err
    }

    value, err := r.ReadString('>')
    if err != nil {
      return nil, err
    }

    key = key[0:len(key)-1]
    value = value[0:len(value)-1]

    switch key {
    case "-UUID":
      d.UUID = value
      break
    case "-Model":
      d.Model = value
      break
    }
  }

  return d, nil
}
