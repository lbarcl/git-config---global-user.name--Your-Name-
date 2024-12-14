package helper

import (
	"bytes"
	"fmt"
	"io"
)

const (
	SEGMENT_BITS byte = 0x7F
	CONTINUE_BIT byte = 0x80
)

// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Data_types

func ReadVarInt(reader io.Reader) (int32, error) {
	var value int32 = 0
	position := 0

	for {
		rawBytes := make([]byte, 1) // Adjust size based on expected data length

		_, err := reader.Read(rawBytes)
		if err != nil {
			return 0, err
		}

		value |= int32(rawBytes[0]&SEGMENT_BITS) << position

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

func WriteVarInt(value int32) []byte {
	buf := &bytes.Buffer{}
	val := uint32(value)
	const maxBytes = 5 // Minecraft protokolünde VarInt maksimum boyut
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

func VarIntByteLength(value int32) int {
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
