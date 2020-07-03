package gorbit

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	block, err := aes.NewCipher(AesKey)
	if err != nil {
		t.Error(err)
	}
	msg := pad([]byte(`hello world`))
	cipherText := make([]byte, aes.BlockSize+len(msg))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		t.Error(err)

	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], msg)
	t.Log(unPadding(base64.URLEncoding.EncodeToString(cipherText)))
}

func TestAesDecrypt(t *testing.T) {
	block, err := aes.NewCipher(AesKey)
	if err != nil {
		t.Error(err)
	}
	secret := `nLK1dPZVOEajDS6tfTT_6NjvsaNzDBj1sHaVQm67-9w`
	bytes, err := base64.URLEncoding.DecodeString(padding(secret))
	if err != nil {
		t.Error(err)
	}
	if (len(bytes) % aes.BlockSize) != 0 {
		t.Error("blockSize must be multiple of decoded message length")
	}
	iv := bytes[:aes.BlockSize]
	msg := bytes[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(msg, msg)
	unPadMsg, err := unPad(msg)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(unPadMsg))
}
