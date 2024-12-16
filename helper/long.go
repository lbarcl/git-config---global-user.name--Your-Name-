package helper

import (
	"encoding/binary"
	"fmt"
	"io"
)

func ReadLong(reader io.Reader) (int64, error) {
	rawBytes := make([]byte, 8) // Buffer for 8 bytes

	// Read 8 bytes into the buffer
	_, err := io.ReadFull(reader, rawBytes)
	if err != nil {
		return 0, fmt.Errorf("failed to read 8 bytes: %w", err)
	}

	// Convert the bytes to int64 using big-endian
	return int64(binary.BigEndian.Uint64(rawBytes)), nil
}

func WriteLong(value int64) []byte {
	val := uint64(value)

	return []byte{
		byte(val >> 56),
		byte(val >> 48),
		byte(val >> 40),
		byte(val >> 32),
		byte(val >> 24),
		byte(val >> 16),
		byte(val >> 8),
		byte(val),
	}
}
