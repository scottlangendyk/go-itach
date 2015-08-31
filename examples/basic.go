package main

import(
  "log"

  "github.com/scottlangendyk/go-itach/itach"
)

func handleDevice(d *itach.Device) {
  c, err := d.Connect()
  if err != nil {
    log.Fatal(err)
  }

  defer c.Close()

  v, err := c.GetVersion()
  if err != nil {
    log.Fatal(err)
  }

  log.Println(v)
}

func main() {
  l, err := itach.ListenForDevices()
  if err != nil {
    log.Fatal(err)
  }

  defer l.Close()

  for {
    d, err := l.Accept()
    if err != nil {
      log.Fatal(err)
    }

    go handleDevice(d)
  }
}
