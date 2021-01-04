package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func pkcs5Padding(plaintext []byte, blockSize int) []byte{
	padding := blockSize-len(plaintext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)},padding)
	return append(plaintext, paddingText...)
}

func pkcs5Unpadding(origData []byte) []byte{
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func aesEncrypt(origData, key []byte) ([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pkcs5Padding(origData,blockSize)
	blockMode := cipher.NewCBCEncrypter(block,key[:blockSize])
	encryptedByteText := make([]byte, len(origData))
	blockMode.CryptBlocks(encryptedByteText,origData)
	return encryptedByteText, nil
}

func aesDecrypt(encryptedByteText, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(encryptedByteText))
	blockMode.CryptBlocks(origData, encryptedByteText)
	origData = pkcs5Unpadding(origData)
	return string(origData), nil
}

// TODO: force AES-256, and call these functions in functions related to JWT generating and parsing
// TODO: add documents for the implemented AES encryption
func AesBase64Encrypt(plainText string, aesKey []byte) (base64String string, err error)  {
	encryptedByteText, err := aesEncrypt([]byte(plainText), aesKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	base64String = base64.StdEncoding.EncodeToString(encryptedByteText)
	return
}

func AesBase64Decrypt(base64String string, aesKey []byte) (plainText string, err error) {
	encryptedByteText, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println(err)
		return
	}
	plainText, err = aesDecrypt(encryptedByteText, aesKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
