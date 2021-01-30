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

func TestAesBase64Encrypt(t *testing.T) {
	t.Parallel()
	for _, v := range testCases {
		v := v
		t.Run(v.name, func(t *testing.T) {
			if v.expectSuccess {
				result, err := AesBase64Encrypt(v.plaintext, v.key)
				assert.Nil(t, err)
				assert.Equal(t, result, v.ciphertext)
			} else {
				_, err := AesBase64Encrypt(v.plaintext, v.key)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestAesBase64Decrypt(t *testing.T) {
	t.Parallel()
	for _, v := range testCases {
		v := v
		t.Run(v.name, func(t *testing.T) {
			if v.expectSuccess {
				result, err := AesBase64Decrypt(v.ciphertext, v.key)
				assert.Nil(t, err)
				assert.Equal(t, result, v.plaintext)
			} else {
				_, err := AesBase64Decrypt(v.ciphertext, v.key)
				assert.NotNil(t, err)
			}
		})
	}
}

/* AES128-CBC-Base64 */
var testCases = []struct {
	name          string
	key           []byte
	plaintext     string
	ciphertext    string
	expectSuccess bool
} {
	{
		name:          "SimpleCase",
		key:           []byte("0123456789012345"),
		plaintext:     "MyAwesomeProject",
		ciphertext:    "d/deUMZPRnAm7so0HWu+oLhg0n5v+m3qCuoyWCmFp8g=",
		expectSuccess: true,
	}, {
		name:          "AnotherSimpleCase",
		key:           []byte("9876543210987654"),
		plaintext:     "=8wsm3OIi8z8UQiq012oz]",
		ciphertext:    "xSydQakavgSXZEITdX73KKnAqGuv98wJiLNB15FqBzM=",
		expectSuccess: true,
	}, {
		name:          "24ByteKeyWhichWillCauseError",
		key:           []byte("012345678901234567890123"),
		plaintext:     "JustATest",
		ciphertext:    "MQB3u9YsQV/PWzlHCA0mfU98GBpG9yTG",
		expectSuccess: false,
	}, {
		name:          "32ByteKeyWhichWillCauseError",
		key:           []byte("01234567890123456789012345678901"),
		plaintext:     "AnotherTest",
		ciphertext:    "Bui0poiD75ZlBGkAMA9QutXNZ50owiUVp6ZdpiSXxAc=",
		expectSuccess: false,
	}, {
		name:          "AnotherInvalidKeyLengthWhichWillCauseError",
		key:           []byte("0123456789"),
		plaintext:     "AnotherTest",
		ciphertext:    "Bui0poiD75ZlBGkAMA9QutXNZ50owiUVp6ZdpiSXxAc=",
		expectSuccess: false,
	},
}
