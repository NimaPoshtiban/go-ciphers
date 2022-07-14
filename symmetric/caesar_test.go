package symmetric

import "testing"



func TestEncrypt(t *testing.T){
	var exampleForEncryption Caesar = "23/AaBcbcZz12&"
	s,_:= exampleForEncryption.Encrypt()
	expectedResult := "23/DdEfefcC12&"
	if s != expectedResult{
		t.Errorf("Expected String(%s) is not same as"+
		" actual string (%s)", expectedResult,s)
	}
}

func TestDecrypt(t *testing.T){
	var exampleForEncryption Caesar = "23/DdEfefcC12&"
	s,_:= exampleForEncryption.Decrypt()
	expectedResult := "23/AaBcbcZz12&"
	if s != expectedResult{
		t.Errorf("Expected String(%s) is not same as"+
		" actual string (%s)", expectedResult,s)
	}
}