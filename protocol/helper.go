package protocol

import (
	"fmt"
	"net"
)

const (
	SEGMENT_BITS byte = 0x7F
	CONTINUE_BIT byte = 0x80
)

func readShort(conn net.Conn) (uint16, error) {
	rawBytes := make([]byte, 2) // Expecting 2 bytes for an unsigned short

	_, err := conn.Read(rawBytes)
	if err != nil {
		return 0, err
	}

	// Combine the two bytes into a uint16
	return uint16(rawBytes[0])<<8 | uint16(rawBytes[1]), nil
}

func readVarInt(conn net.Conn) (int, error) {
	value := 0
	position := 0

	for {
		rawBytes := make([]byte, 1) // Adjust size based on expected data length

		_, err := conn.Read(rawBytes)
		if err != nil {
			fmt.Println("Error reading data:", err)
			return 0, err
		}

		value |= int(rawBytes[0]&SEGMENT_BITS) << position

		if (rawBytes[0] & CONTINUE_BIT) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return 0, fmt.Errorf("VarInt is too big")
		}
	}

	return value, nil
}

func readString(conn net.Conn) string {
	length, err := readVarInt(conn)
	if err != nil {
		fmt.Println(err)
	}

	rawBytes := make([]byte, length) // Adjust size based on expected data length

	_, err = conn.Read(rawBytes)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return ""
	}

	return string(rawBytes)
}
