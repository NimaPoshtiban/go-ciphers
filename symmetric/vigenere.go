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
	for i := 0; i < len(v); i++ {
		plaintext := int(v[i] - 'A')
		key := int(key[i] - 'A')
		cipher := (plaintext + key) % 26
		cipher += 'A'
		cipherText.WriteString(string(cipher))
	}
	return cipherText.String(), nil
}


// Decrypt decrypts the encrypted text 
// and returns the original text
func (v Vigenere) Decrypt(key string) (string, error) {
	if len(string(v)) > 255 {
		return "", errors.New("string is too large")
	}
	key = v.generateKey(key)
	var plaintext strings.Builder
	for i := 0; i < len(v); i++ {
		cipher := int(v[i] - 'A')
		key := int(key[i] - 'A')
		plain := (cipher - key + 26) % 26
		plain += 'A'
		plaintext.WriteString(string(plain))
	}
	return plaintext.String(), nil
}

// generateKey function generates the key in
// a cyclic manner until it's length is
// equal to the length of original text
func (v Vigenere) generateKey(key string) string {
    var genKey strings.Builder
    for i := 0; i < len(v); i++ {
        genKey.WriteString(string(key[i%len(key)]))
    }
    return genKey.String()
}
