package helper

import (
	"io"
)

func ReadLong(reader io.Reader) (int64, error) {
	rawBytes := make([]byte, 8) // Expecting 8 bytes for a long

	_, err := reader.Read(rawBytes)
	if err != nil {
		return 0, err
	}

	// Combine the 8 bytes into an int64
	return int64(rawBytes[0])<<56 | int64(rawBytes[1])<<48 | int64(rawBytes[2])<<40 | int64(rawBytes[3])<<32 | int64(rawBytes[4])<<24 | int64(rawBytes[5])<<16 | int64(rawBytes[6])<<8 | int64(rawBytes[7]), nil
}

func WriteLong(value int64) []byte {
	return []byte{
		byte(value >> 56),
		byte(value >> 48),
		byte(value >> 40),
		byte(value >> 32),
		byte(value >> 24),
		byte(value >> 16),
		byte(value >> 8),
		byte(value),
	}
}
