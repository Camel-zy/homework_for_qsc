package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func pkcs5Padding(plainText []byte, blockSize int) []byte{
	padding := blockSize - len(plainText) % blockSize           // size of the free space in a block
	paddingData := bytes.Repeat([]byte{byte(padding)}, padding) // we need to fill the up free space
	return append(plainText, paddingData...)
}

func pkcs5TransPadding(dataWithPadding []byte) []byte{
	dataLength := len(dataWithPadding)
	paddingLength := int(dataWithPadding[dataLength- 1])
	return dataWithPadding[:(dataLength - paddingLength)] // only return the text before padding text
}

func aesEncrypt(plainText, key []byte) ([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plainText = pkcs5Padding(plainText,blockSize)
	blockMode := cipher.NewCBCEncrypter(block,key[:blockSize])
	encryptedByteData := make([]byte, len(plainText))
	blockMode.CryptBlocks(encryptedByteData, plainText)
	return encryptedByteData, nil
}

func aesDecrypt(encryptedByteText, key []byte) (plainText string, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	buffer := make([]byte, len(encryptedByteText))
	blockMode.CryptBlocks(buffer, encryptedByteText)
	plainText = string(pkcs5TransPadding(buffer))
	return
}

/* Check the validity of a AES-256 key. */
func checkAes256Key(aesKey []byte) error {
	if len(aesKey) != 24 {
		return errors.New("the length of a AES-256 key needs to be 24 bytes")
	}
	return nil
}

// TODO: call these functions in functions related to JWT generating and parsing
func Aes256Base64Encrypt(plainText string, aesKey []byte) (base64String string, err error) {
	err = checkAes256Key(aesKey)
	if err != nil {
		return
	}
	encryptedByteData, err := aesEncrypt([]byte(plainText), aesKey)
	if err != nil {
		return
	}
	base64String = base64.StdEncoding.EncodeToString(encryptedByteData)
	return
}

func Aes256Base64Decrypt(base64String string, aesKey []byte) (plainText string, err error) {
	err = checkAes256Key(aesKey)
	if err != nil {
		return
	}
	encryptedByteData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return
	}
	plainText, err = aesDecrypt(encryptedByteData, aesKey)
	if err != nil {
		return
	}
	return
}
