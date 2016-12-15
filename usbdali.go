package daliclient

import (
	"bufio"
	"fmt"
	"net"
)

// The version of the protocol to be used for commands
const (
	daliVersion = 0x02
)

// Usbdali is a connection to the connected dali devices
// to be used to send commands and receive messages
type Usbdali struct {
	conn net.Conn
}

// Connect establishes a connection to the dali system
func (u *Usbdali) Connect(host string) error {
	if u.isSet() {
		return ErrConnectionAlreadyOpen
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, "55825"))
	u.conn = conn
	return err
}

// Close terminates the connection
func (u *Usbdali) Close() error {
	if !u.isSet() {
		return ErrConnectionClosed
	}
	err := u.conn.Close()
	// TODO Should we check for an error first?
	u.conn = nil
	return err
}

// Send sends a command to the system
func (u *Usbdali) Send(cmd []byte) error {
	if !u.isSet() {
		return ErrConnectionClosed
	}

	// We ignore length because we will get an error if len < len(cmd)
	// and if we don't send everything it's an error for us
	writer := bufio.NewWriter(u.conn)
	_, err := writer.Write(cmd)
	if err != nil {
		return err
	}
	return writer.Flush()
}

// Receive waits blocking until data is received from the dali system
func (u *Usbdali) Receive() ([]byte, error) {
	if !u.isSet() {
		return nil, ErrConnectionClosed
	}

	buffer := make([]byte, 4)
	_, err := u.conn.Read(buffer)

	return buffer, err
}

// isSet returns true if the object already contains a connection
func (u Usbdali) isSet() bool {
	return u.conn != nil
}
