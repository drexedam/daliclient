package daliclient

import "github.com/dropbox/godropbox/errors"

// ErrConnectionAlreadyOpen is returned if an already open connections is tried to be opened again
var ErrConnectionAlreadyOpen = errors.New("The connection is already open")

// ErrConnectionClosed is returned if an operation is called on a closed connection
var ErrConnectionClosed = errors.New("The connection is closed")

// ErrResponseLength is returned if the length of the response is not 4
var ErrResponseLength = errors.New("Invalid response length")