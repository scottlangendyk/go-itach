package itach

import (
  "testing"
)

func TestParseSendIR(t *testing.T) {
  cmd, err := ParseSendIR("sendir,1:2,2445,40000,6,9,4,5,4,5,8,9,4,5,8,9,8,9")
  if err != nil {
    t.Fatal(err)
  }

  if cmd == nil {
    t.Fatal("No command returned")
  }

  if cmd.Frequency != 40000 {
    t.Fatalf("Wrong frequency, expected 40000 got %v", cmd.Frequency)
  }

  if cmd.Repeat != 6 {
    t.Fatalf("Wrong repeat, expected 6 got %v", cmd.Repeat)
  }

  if cmd.Offset != 9 {
    t.Fatalf("Wrong offset, expected 9 got %v", cmd.Offset)
  }

  if len(cmd.Pulses) != 6 {
    t.Fatalf("Wrong number of pulses, expected 6 got %v", len(cmd.Pulses))
  }

  if cmd.Pulses[0].On != 4 {
    t.Fatalf("Wrong on value, expected 4 got %v", cmd.Pulses[0].On)
  }

  if cmd.Pulses[0].Off != 5 {
    t.Fatalf("Wrong on value, expected 5 got %v", cmd.Pulses[0].Off)
  }

  if cmd.Pulses[1].On != 4 {
    t.Fatalf("Wrong on value, expected 4 got %v", cmd.Pulses[1].On)
  }

  if cmd.Pulses[1].Off != 5 {
    t.Fatalf("Wrong on value, expected 5 got %v", cmd.Pulses[1].Off)
  }

  if cmd.Pulses[2].On != 8 {
    t.Fatalf("Wrong on value, expected 8 got %v", cmd.Pulses[2].On)
  }

  if cmd.Pulses[2].Off != 9 {
    t.Fatalf("Wrong on value, expected 9 got %v", cmd.Pulses[2].Off)
  }
}
