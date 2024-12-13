package helper

import "io"

func ReadBoolean(reader io.Reader) (bool, error) {
	rawByte, err := ReadBytes(reader, 1)

	if err != nil {
		return false, err
	}

	return rawByte[0] == 1, nil

}
