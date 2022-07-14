package symmetric

import (
	"errors"
)

type Caesar string

// Encrypt encrypts the text using caesar algorithm
func (c Caesar) Encrypt()(string,error){
	if len(string(c)) > 255{
		return "",errors.New("string is too large")
	}
	hash := make([]rune,len(string(c)))
	for k,v := range c {
		if v >= 65 && v < 88 || (v >=97 && v < 120){
			hash[k] = v+3
		}else if v >= 120 && v <123 {
			hash[k] = v - 55
		}else if v >=88 && v <=90{
			hash[k] = v + 9
		}else{
			hash[k] = v
		}
	}
	return string(hash),nil
}

// Decrypt decrypts the text using caesar algorithm
func (c Caesar) Decrypt()(string,error){
	if len(string(c)) > 255{
		return "",errors.New("string is too large")
	}
	hash := make([]rune,len(string(c)))
	for k,v := range c{
		if v >= 68 && v <= 90 || (v <123 && v >= 100){
			hash[k] = v-3
		}else if v < 68 && v >= 65{
			hash[k] = v + 55
		}else if v <100 && v >= 97{
			hash[k] = v -9
		}else {
			hash[k]=v
		}
	}
	return string(hash),nil
}