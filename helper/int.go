package helper

import "io"

func ReadInt(reader io.Reader) (int, error) {
	rawBytes := make([]byte, 4) // Expecting 4 bytes for an int

	_, err := reader.Read(rawBytes)
	if err != nil {
		return 0, err
	}

	// Combine the 4 bytes into an int
	return int(rawBytes[0])<<24 | int(rawBytes[1])<<16 | int(rawBytes[2])<<8 | int(rawBytes[3]), nil
}
