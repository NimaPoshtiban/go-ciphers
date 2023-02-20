package symmetric

import (
	"testing"
)

func TestEncryptVigenere(t *testing.T) {
	var exampleForEncryption Vigenere = "GEEKSFORGEEKS"
	key := "AYUSH"
	s, _ := exampleForEncryption.Encrypt(key)
	expectedResult := "GCYCZFMLYLEIM"
	if expectedResult != s {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedResult, s)
	}
}

func TestDecryptVigenere(t *testing.T) {
	var exampleForDecryption Vigenere = "GCYCZFMLYLEIM"
	key := "AYUSH"
	s, _ := exampleForDecryption.Decrypt(key)
	expectedResult := "GEEKSFORGEEKS"
	if expectedResult != s {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedResult, s)
	}
}

func Test_VignereGenerateKey(t *testing.T) {
	var example Vigenere = "SOMETEXT"
	key := "APM"
	key = example.generateKey(key)
	if len(key) != len(example) {
		t.Errorf("Expected String(%s) is not long as"+
			" actual string (%s)", key, example)
	}
}
