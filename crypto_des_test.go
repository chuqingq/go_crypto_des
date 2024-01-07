package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDes(t *testing.T) {
	data := [][2][]byte{
		{[]byte("123"), []byte("abcdef")},
		{[]byte("12345678"), []byte("abcdefghi")},
		{[]byte("123456789012"), []byte("abcdefghijk")},
		{[]byte("12345678901234567890123"), []byte("abcdefghijklmn11222333445556677888990")},
	}

	for _, d := range data {
		key := d[0]
		plaintext := d[1]

		ciphertext, err := DESEncrypt(key, plaintext)
		assert.NoError(t, err)

		// fmt.Printf("Ciphertext: %x\n", ciphertext)

		decrypted, err := DESDecrypt(key, ciphertext)
		assert.NoError(t, err)

		assert.Equal(t, plaintext, decrypted)
		// fmt.Printf("plain %v -> %v\n", len(plaintext), len(ciphertext))

		// fmt.Printf("Decrypted: %s\n", decrypted)
	}
}
