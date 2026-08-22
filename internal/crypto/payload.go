package crypto

import (
	"encoding/base64"
	"fmt"
)

func SealPayload(payload string, secret []byte) string {
	if len(secret) == 0 {
		return ""
	}
	data := []byte(payload)
	for i := range data {
		data[i] ^= secret[i%len(secret)]
	}
	return base64.RawStdEncoding.EncodeToString(data)
}

func OpenPayload(encoded string, secret []byte) (string, error) {
	if len(secret) == 0 {
		return "", fmt.Errorf("secret is empty")
	}
	data, err := base64.RawStdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	for i := range data {
		data[i] ^= secret[i%len(secret)]
	}
	return string(data), nil
}

func PayloadDigest(payload string) string { return Fingerprint([]byte(payload)) }

func IsPrintable(payload string) bool {
	for _, r := range payload {
		if r < 32 || r == 127 {
			return false
		}
	}
	return true
}
