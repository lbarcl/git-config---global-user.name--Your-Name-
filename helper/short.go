package helper

import "net"

func ReadShort(conn net.Conn) (uint16, error) {
	rawBytes := make([]byte, 2) // Expecting 2 bytes for an unsigned short

	_, err := conn.Read(rawBytes)
	if err != nil {
		return 0, err
	}

	// Combine the two bytes into a uint16
	return uint16(rawBytes[0])<<8 | uint16(rawBytes[1]), nil
}
