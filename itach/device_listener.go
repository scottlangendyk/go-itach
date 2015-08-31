package itach

import "net"

type DeviceListener struct {
  conn *net.UDPConn
}

// Accept waits for and returns the next discovered Device.
func (l *DeviceListener) Accept() (d *Device, err error) {
  // Keep looping so we only return valid beacons
  for {
    var buf []byte = make([]byte, 1500)

    rlen, addr, err := l.conn.ReadFromUDP(buf)
    if err != nil {
      return nil, err
    }

    d, _ := ParseBeacon(string(buf[0:rlen]))

    d.IP = addr.IP

    switch d.Model {
    case "iTachIP2IR", "iTachWF2IR", "iTachIP2SL", "iTachWF2SL", "iTachIP2CC", "iTachWF2CC":
      return d, nil
    }
  }
}

// Close closes the listener.
func (l *DeviceListener) Close() error {
  return l.conn.Close()
}

// ListenForDevices listens for incoming device beacons
func ListenForDevices() (l *DeviceListener, err error) {
  l = &DeviceListener{}

  addr, err  := net.ResolveUDPAddr("udp4", "239.255.250.250:9131")
  if err != nil {
    return nil, err
  }

  l.conn, err = net.ListenMulticastUDP("udp4", nil, addr)
  if err != nil {
    return nil, err
  }

  return l, nil
}
