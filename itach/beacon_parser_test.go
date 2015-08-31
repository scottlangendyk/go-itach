package itach

import (
  "testing"
)

func TestParseBeacon(t *testing.T) {
  beacon := "AMXB<-UUID=GlobalCache_000C1E024239><-SDKClass=Utility> <-Make=GlobalCache><-Model=iTachWF2IR><-Revision=710-1001-05> <-Pkg_Level=GCPK001><-Config-URL=http://192.168.1.100.><-PCB_PN=025-0026-06> <-Status=Ready>"

  device, err := ParseBeacon(beacon)
  if err != nil {
    t.Fatal(err)
  }

  if device.Model != "iTachWF2IR" {
    t.Fatalf("Wrong model, expected `iTachWF2IR`, got `%v`", device.Model)
  }

  if device.UUID != "GlobalCache_000C1E024239" {
    t.Fatalf("Wrong UUID, expected `GlobalCache_000C1E024239`, got `%v`", device.UUID)
  }
}
