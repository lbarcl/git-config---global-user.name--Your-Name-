package helper

import (
	"math/rand/v2"
)

func RandomEIDGen() int32 {
	var number int32
	var i int32
	for i = 1; i < 1000; i += 10 {
		number += i * rand.Int32N(200)
	}

	return number
}
