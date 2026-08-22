package crypto

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func WrapSecret(publicKey, gateID, sessionDate string, profile Profile) (string, error) {
	if err := profile.Validate(); err != nil {
		return "", err
	}
	if publicKey == "" || gateID == "" || sessionDate == "" {
		return "", fmt.Errorf("missing wrapping input")
	}
	secret := DeriveSecret(gateID, sessionDate)
	key := DeriveKey(publicKey)
	wrapped := xor(secret, key)
	check := checksum(profile, gateID, sessionDate, wrapped)
	payload := append([]byte{profile.VersionByte}, wrapped...)
	payload = append(payload, check...)
	return profile.Prefix + "." + base64.RawURLEncoding.EncodeToString(payload), nil
}

func xor(left, right []byte) []byte {
	out := make([]byte, len(left))
	for i := range left {
		out[i] = left[i] ^ right[i%len(right)]
	}
	return out
}

func checksum(profile Profile, gateID, sessionDate string, wrapped []byte) []byte {
	h := sha256.New()
	h.Write([]byte(profile.ID))
	h.Write([]byte("|"))
	h.Write([]byte(gateID))
	h.Write([]byte("|"))
	h.Write([]byte(sessionDate))
	h.Write(wrapped)
	sum := h.Sum(nil)
	return sum[:8]
}
