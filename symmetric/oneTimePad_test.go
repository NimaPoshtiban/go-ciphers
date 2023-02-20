package symmetric

import "testing"

func TestOneTimePad(t *testing.T) {
	c := OneTimePad{Key: "1234567", PlainText: "ILoveGo"}
	enc := c.Encrypt()
	expectedResult := "x~\\BPqX"
	if enc != expectedResult {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedResult, enc)
	}
	dec := c.Decrypt()
	if dec != c.PlainText {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", c.PlainText, dec)
	}
}
