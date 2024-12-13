package helper

import (
	"io"
)

func ReadBytes(reader io.Reader, length int) ([]byte, error) {
	rawBytes := make([]byte, length)

	_, err := reader.Read(rawBytes)
	if err != nil {
		return nil, err
	}

	return rawBytes, nil
}
