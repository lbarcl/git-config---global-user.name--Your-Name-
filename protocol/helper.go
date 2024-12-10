package protocol

import (
	"fmt"
	"net"
)

const (
	SEGMENT_BITS byte = 0x7F
	CONTINUE_BIT byte = 0x80
)

// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Data_types

func readShort(conn net.Conn) (uint16, error) {
	rawBytes := make([]byte, 2) // Expecting 2 bytes for an unsigned short

	_, err := conn.Read(rawBytes)
	if err != nil {
		return 0, err
	}

	// Combine the two bytes into a uint16
	return uint16(rawBytes[0])<<8 | uint16(rawBytes[1]), nil
}

func readLong(conn net.Conn) (int64, error) {
	rawBytes := make([]byte, 8) // Expecting 8 bytes for a long

	_, err := conn.Read(rawBytes)
	if err != nil {
		return 0, err
	}

	// Combine the 8 bytes into an int64
	return int64(rawBytes[0])<<56 | int64(rawBytes[1])<<48 | int64(rawBytes[2])<<40 | int64(rawBytes[3])<<32 | int64(rawBytes[4])<<24 | int64(rawBytes[5])<<16 | int64(rawBytes[6])<<8 | int64(rawBytes[7]), nil
}

func readVarInt(conn net.Conn) (int, error) {
	value := 0
	position := 0

	for {
		rawBytes := make([]byte, 1) // Adjust size based on expected data length

		_, err := conn.Read(rawBytes)
		if err != nil {
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

func writeVarInt(value int) []byte {
	var encoded []byte

	for {
		if (value & ^int(SEGMENT_BITS)) == 0 {
			encoded = append(encoded, byte(value))
			break
		} else {
			encoded = append(encoded, byte(value&int(SEGMENT_BITS))|CONTINUE_BIT)
			value = value >> 7
		}
	}

	return encoded
}

func writeLong(value int64) []byte {
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
