package helper

import (
	"bytes"
	"io"
)

func ReadLong(reader io.Reader) (int64, error) {
	rawBytes := make([]byte, 8) // Expecting 8 bytes for a long

	_, err := reader.Read(rawBytes)
	if err != nil {
		return 0, err
	}

	// Combine the 8 bytes into an int64
	return int64(rawBytes[0])<<56 | int64(rawBytes[1])<<48 | int64(rawBytes[2])<<40 | int64(rawBytes[3])<<32 | int64(rawBytes[4])<<24 | int64(rawBytes[5])<<16 | int64(rawBytes[6])<<8 | int64(rawBytes[7]), nil
}

func WriteLong(value int64) []byte {
	buf := &bytes.Buffer{}
	val := uint64(value)
	const maxBytes = 10 // Minecraft protokolünde VarInt maksimum boyut
	for i := 0; i < maxBytes; i++ {
		// İlk 7 bitlik kısmı al
		b := byte(val) & 0b01111111

		// İşaretli kaydırma işlemi: sağa doğru kaydır ve işareti koru
		val >>= 7

		// Eğer hala verimiz varsa veya negatif değerle uğraşıyorsak devam biti ekle
		if val != 0 {
			b |= 0b10000000
		}

		// Baytı buffer'a yaz
		buf.WriteByte(b)

		// Eğer devam biti (`MSB`) set değilse işlem biter
		if val == 0 {
			break
		}
	}

	return buf.Bytes()
}
