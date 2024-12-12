package helper

import (
	"fmt"
	"net"
)

func ReadString(conn net.Conn) string {
	length, err := ReadVarInt(conn)
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

func WriteString(data string) []byte {
	length := WriteVarInt(len(data))
	rawBytes := append(length, []byte(data)...)

	return rawBytes
}
