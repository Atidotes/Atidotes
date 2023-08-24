package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// 加密
func Encryption(target string) (bytes []byte, err error) {
	bytes, err = bcrypt.GenerateFromPassword([]byte(target), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return bytes, err
}

// 解密
func Decrypt(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
