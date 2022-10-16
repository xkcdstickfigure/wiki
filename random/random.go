package random

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"strconv"
)

func String(length int) string {
	return generate(length, []rune("0123456789abcdefghijklmnopqrstuvwxyz"))
}

func Number(length int) int {
	n, _ := strconv.Atoi(generate(length, []rune("0123456789")))
	return n
}

func generate(length int, chars []rune) string {
	var buf bytes.Buffer
	buf.Grow(length)
	l := uint32(len(chars))
	for i := 0; i < length; i++ {
		buf.WriteRune(chars[binary.BigEndian.Uint32(randBytes(4))%l])
	}
	return buf.String()
}

func randBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
