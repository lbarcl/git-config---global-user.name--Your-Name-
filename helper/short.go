package helper

import (
	"io"
)

func ReadShort(reader io.Reader) (uint16, error) {
	rawBytes := make([]byte, 2) // Expecting 2 bytes for an unsigned short

	_, err := reader.Read(rawBytes)
	if err != nil {
		return 0, err
	}

	// Combine the two bytes into a uint16
	return uint16(rawBytes[0])<<8 | uint16(rawBytes[1]), nil
}
