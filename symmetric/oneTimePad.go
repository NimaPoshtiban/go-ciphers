// One-time pad encryption algorithm
package symmetric

import "errors"

type OneTimePad struct {
	Key       string
	PlainText string
	Hash      string
}

func (o *OneTimePad) validate() error {
	if len(o.PlainText) != len(o.Key) {
		return errors.New("key and text must have equal length")
	}
	return nil
}

func (o *OneTimePad) Encrypt() string {
	err := o.validate()
	if err != nil {
		return ""
	}

	hash := make([]byte, len(o.Key))

	for i := 0; i < len(o.PlainText); i++ {
		hash[i] = o.PlainText[i] ^ o.Key[i]
	}
	o.Hash = string(hash)
	return o.Hash
}

func (o *OneTimePad) Decrypt() string {
	err := o.validate()
	if err != nil {
		return err.Error()
	}
	text := make([]byte, len(o.Key))
	for i := 0; i < len(o.Key); i++ {
		text[i] = o.Hash[i] ^ o.Key[i]
	}
	o.PlainText = string(text)
	return o.PlainText
}
