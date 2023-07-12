package logic

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

var key = []byte("1234567890123456")

// 对二进制数据进行 AES 加密
func aESEncrypt(src []byte) []byte {
	cipherBlock, _ := aes.NewCipher(generateKey(key))
	s := cipher.NewCBCEncrypter(cipherBlock, key)
	src = pkcs7Padding(src, aes.BlockSize)
	encrypted := make([]byte, len(src))
	s.CryptBlocks(encrypted, src)
	return encrypted
}

// 对二进制数据进行 AES 解密
func aESDecrypt(encrypted []byte) ([]byte, error) {
	cipherBlock, _ := aes.NewCipher(generateKey(key))
	s := cipher.NewCBCDecrypter(cipherBlock, key)
	decrypted := make([]byte, len(encrypted))
	s.CryptBlocks(decrypted, encrypted)
	decrypted, err := pkcs7UnPadding(decrypted)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

// generateKey 生成16字节密钥
func generateKey(key []byte) []byte {
	genKey := make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// pkcs7Padding 对数据进行 pkcs7 填充
func pkcs7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	pt := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, pt...)
}

// pkcs7UnPadding 对 pkcs7 填充的数据进行去除填充
func pkcs7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	up := int(src[length-1])
	if up > length {
		return nil, errors.New("invalid padding")
	}
	return src[:(length - up)], nil
}
