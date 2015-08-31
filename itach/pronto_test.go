package itach

import (
  "testing"
)

func TestParsePronto(t *testing.T) {
  c, err := ParsePronto("0000 006C 0000 000E 00BE 004C 0025 0025 0025 0025 0025 0072 0025 0025 0025 0025 0025 0072 0025 0025 0025 0025 0025 0025 0025 0072 0025 0072 0025 0025 0025 091B")
  if err != nil {
    t.Fatal(err)
  }

  if c.Frequency > 38400 || c.Frequency < 38300 {
    t.Fatalf("Wrong frequency, expected `38380`, got `%v`'", c.Frequency)
  }

  if len(c.Pulses) != 14 {
    t.Fatalf("Expected 14 pulses, got %v", len(c.Pulses))
  }
}

func TestProntoReader(t *testing.T) {
  r := NewProntoReader("0000 006C 0000 000E 00BE 004C 0025 0025 0025 0025 0025 0072 0025 0025 0025 0025 0025 0072 0025 0025 0025 0025 0025 0025 0025 0072 0025 0072 0025 0025 0025 091B")

  w, _ := r.ReadWord()
  if (w != "0000") {
    t.Fatalf("Expected 0000 got %v", w)
  }

  w, _ = r.ReadWord()
  if (w != "006C") {
    t.Fatalf("Expected 006C got %v", w)
  }

  w, _ = r.ReadWord()
  if (w != "0000") {
    t.Fatalf("Expected 0000 got %v", w)
  }

  w, _ = r.ReadWord()
  if (w != "000E") {
    t.Fatalf("Expected 000E got %v", w)
  }

  w, _ = r.ReadWord()
  if (w != "00BE") {
    t.Fatalf("Expected 00BE got %v", w)
  }

  w, _ = r.ReadWord()
  if (w != "004C") {
    t.Fatalf("Expected 004C got %v", w)
  }

  w, _ = r.ReadWord()
  if (w != "0025") {
    t.Fatalf("Expected 0025 got %v", w)
  }
}
