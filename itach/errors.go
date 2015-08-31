package itach

import "errors"

// ErrorForCode returns an error for the specified iTach
// error code.
func ErrorForCode(code int64) (err error) {
  switch code {
  case 1:
    err = errors.New("Invalid command. Command not found.")
  case 2:
    err = errors.New("Invalid module address (does not exist).")
  case 3:
    err = errors.New("Invalid connector address (does not exist).")
  case 4:
    err = errors.New("Invalid ID value.")
  case 5:
    err = errors.New("Invalid frequency value.")
  case 6:
    err = errors.New("Invalid repeat value.")
  case 7:
    err = errors.New("Invalid offset value.")
  case 8:
    err = errors.New("Invalid pulse count.")
  case 9:
    err = errors.New("Invalid pulse data.")
  case 10:
    err = errors.New("Uneven amount of <on|off> statements.")
  case 11:
    err = errors.New("No carriage return found.")
  case 12:
    err = errors.New("Repeat count exceeded.")
  case 13:
    err = errors.New("IR command sent to input connector.")
  case 14:
    err = errors.New("Blaster command sent to non-blaster connector.")
  case 15:
    err = errors.New("No carriage return before buffer full.")
  case 16:
    err = errors.New("No carriage return.")
  case 17:
    err = errors.New("Bad command syntax.")
  case 18:
    err = errors.New("Sensor command sent to non-input connector.")
  case 19:
    err = errors.New("Repeated IR transmission failure.")
  case 20:
    err = errors.New("Above designated IR <on|off> pair limit.")
  case 21:
    err = errors.New("Symbol odd boundary.")
  case 22:
    err = errors.New("Undefined symbol.")
  case 23:
    err = errors.New("Unknown option.")
  case 24:
    err = errors.New("Invalid baud rate setting.")
  case 25:
    err = errors.New("Invalid flow control setting.")
  case 26:
    err = errors.New("Invalid parity setting.")
  case 27:
    err = errors.New("Settings are locked.")
  default:
    err = errors.New("Unknown error code")
  }

  return err
}
