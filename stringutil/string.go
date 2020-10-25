package stringutil

import (
	"math/rand"
	"time"
)

var (
	_rand  = rand.New(rand.NewSource(time.Now().UnixNano()))
	_chars = []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'v', 'x', 'y', 'z', 'A', 'B', 'C', 'D',
		'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
		'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
		'Y', 'Z', '_', '-',
	}
	_charCount = len(_chars)
)

func Rand(length int) string {
	bs := make([]byte, length)

	for i := 0; i < length; i++ {
		bs[i] = _chars[_rand.Intn(_charCount)]
	}

	return string(bs)
}
