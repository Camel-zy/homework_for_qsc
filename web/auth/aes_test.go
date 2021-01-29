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

type aesBase64Test struct {
	key        []byte
	plaintext  string
	ciphertext string
}

/* AES128-CBC-Base64 */
var successTests = []aesBase64Test {
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
var failedTests = []aesBase64Test {
	{       // 24 byte key
		key: []byte("012345678901234567890123"),
		plaintext: "JustATest",
		ciphertext: "MQB3u9YsQV/PWzlHCA0mfU98GBpG9yTG",
	}, {    // 32 byte key
		key: []byte("01234567890123456789012345678901"),
		plaintext: "AnotherTest",
		ciphertext: "Bui0poiD75ZlBGkAMA9QutXNZ50owiUVp6ZdpiSXxAc=",
	}, {    // 10 byte key
		key: []byte("0123456789"),
		plaintext: "AnotherTest",
		ciphertext: "Bui0poiD75ZlBGkAMA9QutXNZ50owiUVp6ZdpiSXxAc=",
	},
}

func TestAesBase64Encrypt(t *testing.T) {
	for _, testCase := range successTests {
		result, err := AesBase64Encrypt(testCase.plaintext, testCase.key)
		assert.Nil(t, err)
		assert.Equal(t, result, testCase.ciphertext)
	}
	for _, testCase := range  failedTests {
		_, err := AesBase64Encrypt(testCase.plaintext, testCase.key)
		assert.NotNil(t, err)
	}
}

func TestAesBase64Decrypt(t *testing.T) {
	for _, testCase := range successTests {
		result, err := AesBase64Decrypt(testCase.ciphertext, testCase.key)
		assert.Nil(t, err)
		assert.Equal(t, result, testCase.plaintext)
	}
	for _, testCase := range  failedTests {
		_, err := AesBase64Decrypt(testCase.plaintext, testCase.key)
		assert.NotNil(t, err)
	}
}
