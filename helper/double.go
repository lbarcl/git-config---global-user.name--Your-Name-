package helper

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func DoubleToBytes(value float64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert double to bytes: %v", err))
	}
	return buf.Bytes()
}

func FloatToBytes(value float32) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert double to bytes: %v", err))
	}
	return buf.Bytes()
}
