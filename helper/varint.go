package helper

import (
	"fmt"
	"io"
)

const (
	SEGMENT_BITS byte = 0x7F
	CONTINUE_BIT byte = 0x80
)

// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Data_types

func ReadVarInt(reader io.Reader) (int, error) {
	value := 0
	position := 0

	for {
		rawBytes := make([]byte, 1) // Adjust size based on expected data length

		_, err := reader.Read(rawBytes)
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

func WriteVarInt(value int) []byte {
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

func VarIntByteLength(value int) int {
	length := 0
	for {
		length++
		value >>= 7
		if value == 0 {
			break
		}
	}
	return length
}
