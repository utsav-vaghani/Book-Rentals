package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

//Encrypt using bcrypt
func Encrypt(str string) (string, error) {
	Str := []byte(str)

	hashedStr, er := bcrypt.GenerateFromPassword(Str, bcrypt.DefaultCost)
	log.Println(string(hashedStr))
	return string(hashedStr), er
}

//Decrypt using bcrypt
func Decrypt(str, hashedStr string) bool {
	er := bcrypt.CompareHashAndPassword([]byte(str), []byte(hashedStr))
	if er != nil {
		log.Println(er)
		return false
	}

	return true
}
