package itach

import (
  "bytes"
  "errors"
  "strconv"
  "strings"
  "unicode"
)

// ProntoReader is a Reader with convenience functions for reading
// Pronto hex formatted ir command strings.
type ProntoReader struct {
  *strings.Reader
}

// NewProntoReader returns a ProntoReader for the s string.
func NewProntoReader(s string) *ProntoReader {
  return &ProntoReader{strings.NewReader(s)}
}

// ReadWord reads a 4 character hex word and returns the raw
// string value.
func (r *ProntoReader) ReadWord() (w string, err error) {
  buf := new(bytes.Buffer)

  for buf.Len() < 4 {
    ch, _, err := r.ReadRune()
    if err != nil {
      return "", err
    }

    if !unicode.IsSpace(ch) {
      buf.WriteRune(ch)
    }
  }

  return buf.String(), nil
}

// ReadHexWord reads a 4 character hex word and returns the
// integer value.
func (r *ProntoReader) ReadHexWord() (w int, err error) {
  s, err := r.ReadWord()
  if err != nil {
    return 0, err
  }

  i, err := strconv.ParseInt(s, 16, 0)
  if err != nil {
    return 0, err
  }

  return int(i), nil
}

// ParsePronto converts an IR command string in the Pronto hex
// to an IRCommand.
//
// https://www.remotecentral.com/features/irdisp2.htm
func ParsePronto(code string) (c *IRCommand, err error) {
  c = &IRCommand{Repeat: 1, Offset: 1}
  r := NewProntoReader(code)

  w, err := r.ReadHexWord()
  if err != nil {
    return nil, err
  }

  if w != 0 {
    return nil, errors.New("Invalid pronto code")
  }

  w, err = r.ReadHexWord()
  if err != nil {
    return nil, err
  }

  c.Frequency = uint(1000000 / (float64(w) * 0.241246))

  n, err := r.ReadHexWord()
  if err != nil {
    return nil, err
  }

  w, err = r.ReadHexWord()
  if err != nil {
    return nil, err
  }

  n = n + w

  c.Pulses = make([]IRPulse, n)

  for i := 0; i < n; i++ {
    c.Pulses[i] = IRPulse{}

    w, err = r.ReadHexWord()
    if err != nil {
      return nil, err
    }

    c.Pulses[i].On = uint(w)

    w, err = r.ReadHexWord()
    if err != nil {
      return nil, err
    }

    c.Pulses[i].Off = uint(w)
  }

  return c, nil
}
