package helper

import "net"

func ReadBytes(conn net.Conn, length int) ([]byte, error) {
	rawBytes := make([]byte, length)

	_, err := conn.Read(rawBytes)
	if err != nil {
		return nil, err
	}

	return rawBytes, nil
}
