package helper

import (
	"fmt"
	"io"
)

func ReadString(reader io.Reader) string {
	length, err := ReadVarInt(reader)
	if err != nil {
		fmt.Println(err)
	}

	rawBytes := make([]byte, length) // Adjust size based on expected data length

	_, err = reader.Read(rawBytes)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return ""
	}

	return string(rawBytes)
}

func WriteString(data string) []byte {
	length := WriteVarInt(int32(len(data)))
	rawBytes := append(length, []byte(data)...)

	return rawBytes
}
