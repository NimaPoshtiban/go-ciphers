package symmetric

import (
	"errors"
	"strings"
)
// Vigenere cipher 
// for more info take a
// look at https://www.geeksforgeeks.org/vigenere-cipher
type Vigenere string


// Encrypt encrypts text with the help of provided key
func (v Vigenere) Encrypt(key string) (string, error) {
	if len(string(v)) > 255 {
		return "", errors.New("string is too large")
	}
	key = v.generateKey(key)
	var cipherText strings.Builder
	for i:=0;i<len(v);i++ {
		x := (v[i] + key[i]) % 26
		x += 'A'
		// cipherText +=string(x)
		cipherText.WriteString(string(x))
	}
	return cipherText.String(),nil
}


// Decrypt decrypts the encrypted text 
// and returns the original text
func (v Vigenere) Decrypt(key string)(string,error){
	if len(string(v)) > 255 {
		return "", errors.New("string is too large")
	}
	key = v.generateKey(key)
	var originalText strings.Builder 
	for i := 0; i < len(v)&&i<len(key); i++ {
		x:=(v[i] - key[i]+26) % 26
		x += 'A'
		// originalText += string(x)
		originalText.WriteString(string(x))
	}
	return originalText.String(),nil
}
// generateKey function generates the key in
// a cyclic manner until it's length is
// equal to the length of original text
func (v Vigenere) generateKey(key string) string{
	x := len(key)
	for i:=0;;i++{
		if x==i{
			i = 0
		}
		if len(key)==len(v){
			break
		}
		key +=(string(key[i]))
	}
	return key
}