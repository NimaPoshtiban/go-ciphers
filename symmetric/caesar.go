package symmetric

import (
	"errors"
)

type Caesar string

// Encrypt encrypts the text using caesar algorithm
func (c Caesar) Encrypt()(string,error){
    
	hash := make([]rune,len(string(c)))
	for k,v := range c {
		if v >= 65 && v < 68 { // for A-C
			hash[k] = v+23
		} else if v >= 97 && v < 100 { // for a-c
			hash[k] = v+23
		} else {
			hash[k] = v + 3
		}
	}
	return string(hash),nil
}

// Decrypt decrypts the text using caesar algorithm
func (c Caesar) Decrypt()(string,error){
    
	hash := make([]rune,len(string(c)))
	for k,v := range c{
		if v > 87 && v <= 90 { // for X-Z
			hash[k] = v-23
		} else if v > 119 && v <= 122 { // for x-z
			hash[k] = v-23
		} else {
			hash[k] = v - 3
		}
	}
	return string(hash),nil
}
