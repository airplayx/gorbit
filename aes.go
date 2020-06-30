package gorbit

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strings"
)

func AesEncrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	msg := pad([]byte(text))
	cipherText := make([]byte, aes.BlockSize+len(msg))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], msg)
	finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(cipherText))
	return finalMsg, nil
}

func AesDecrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(text))
	if err != nil {
		return "", err
	}
	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("blockSize must be multiple of decoded message length")
	}
	iv := decodedMsg[:aes.BlockSize]
	msg := decodedMsg[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(msg, msg)
	unPadMsg, err := unPad(msg)
	if err != nil {
		return "", err
	}
	return string(unPadMsg), nil
}

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}
	return value
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	repeat := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, repeat...)
}

func unPad(src []byte) ([]byte, error) {
	length := len(src)
	unPadding := int(src[length-1])
	if unPadding > length {
		return nil, errors.New("unPad error. This could happen when incorrect encryption key is used")
	}
	return src[:(length - unPadding)], nil
}
