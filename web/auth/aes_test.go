package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckAesKey(t *testing.T) {
	assert.Nil(t, checkAesKey(make([]byte, 16)))     // AES-128
	assert.NotNil(t, checkAesKey(make([]byte, 24)))  // AES-192
	assert.NotNil(t, checkAesKey(make([]byte, 32)))  // AES-256
}

/* AES128-CBC-Base64 */
var aesBase64Tests = []struct {
	key        []byte
	plaintext  string
	ciphertext string
} {
	{
		key:        []byte("0123456789012345"),
		plaintext:  "MyAwesomeProject",
		ciphertext: "d/deUMZPRnAm7so0HWu+oLhg0n5v+m3qCuoyWCmFp8g=",
	}, {
		key:        []byte("9876543210987654"),
		plaintext:  "=8wsm3OIi8z8UQiq012oz]",
		ciphertext: "xSydQakavgSXZEITdX73KKnAqGuv98wJiLNB15FqBzM=",
	},
}

func TestAesBase64Encrypt(t *testing.T) {
	for _, testCase := range aesBase64Tests {
		result, err := AesBase64Encrypt(testCase.plaintext, testCase.key)
		assert.Nil(t, err)
		assert.Equal(t, result, testCase.ciphertext)
	}
}

func TestAesBase64Decrypt(t *testing.T) {
	for _, testCase := range aesBase64Tests {
		result, err := AesBase64Decrypt(testCase.ciphertext, testCase.key)
		assert.Nil(t, err)
		assert.Equal(t, result, testCase.plaintext)
	}
}
