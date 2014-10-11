package co

import "bytes"
import "crypto/hmac"
import "encoding/base64"
import "hash"

// Messenger interface
type Messenger interface {
	Message() ([]byte, error)
}

type shaNewFunc func() hash.Hash

// Signature type
type Signature []byte

// Sign returns an hmac Signature
func Sign(m Messenger, h shaNewFunc, key []byte) (Signature, error) {
	return SignSum(m, h, key, nil)
}

// SignSum allows you to add sum bytes to the signature
func SignSum(m Messenger, h shaNewFunc, key, sum []byte) (Signature, error) {
	mac := hmac.New(h, key)
	msg, err := m.Message()
	if err != nil {
		return nil, err
	}

	mac.Write(msg)
	return mac.Sum(sum), nil
}

// Base64 encodes a Signature to Base64, using StdEncoding
func (s Signature) Base64() ([]byte, error) {
	return s.Base64Encoding(base64.StdEncoding)
}

// Base64Encoding encodes a Signture to Base64 with the given encoding
func (s Signature) Base64Encoding(e *base64.Encoding) ([]byte, error) {
	var b []byte
	w := bytes.NewBuffer(b)

	enc := base64.NewEncoder(e, w)
	enc.Write(s)
	if err := enc.Close(); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}
