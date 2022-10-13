package randstr

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

func Generate(n int) string {
	var buf bytes.Buffer
	buf.Grow(n)
	l := uint32(len(letters))
	for i := 0; i < n; i++ {
		buf.WriteRune(letters[binary.BigEndian.Uint32(randBytes(4))%l])
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
