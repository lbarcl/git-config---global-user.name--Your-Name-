package helper

import (
	"encoding/hex"
	"io"
	"strings"
)

func ReadUUID(reader io.Reader) string {
	rawBytes := make([]byte, 16)
	_, err := reader.Read(rawBytes)

	if err != nil {
		return ""
	}

	// Convert rawBytes to UUID format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	uuid := strings.Join([]string{
		hex.EncodeToString(rawBytes[0:4]),
		hex.EncodeToString(rawBytes[4:6]),
		hex.EncodeToString(rawBytes[6:8]),
		hex.EncodeToString(rawBytes[8:10]),
		hex.EncodeToString(rawBytes[10:16]),
	}, "-")

	return uuid
}

func WriteUUID(uuid string) []byte {

	// Convert UUID format to rawBytes: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	uuidParts := strings.Split(uuid, "-")
	rawBytes := []byte{}

	for _, part := range uuidParts {
		partBytes, _ := hex.DecodeString(part)
		rawBytes = append(rawBytes, partBytes...)
	}

	return rawBytes
}
