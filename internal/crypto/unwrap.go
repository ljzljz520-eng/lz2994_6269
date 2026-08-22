package crypto

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func UnwrapSecret(ciphertext, publicKey, gateID, sessionDate string, profile Profile) ([]byte, error) {
	if err := profile.Validate(); err != nil {
		return nil, err
	}
	if ciphertext == "" || publicKey == "" || gateID == "" || sessionDate == "" {
		return nil, fmt.Errorf("missing unwrap input")
	}
	parts := strings.Split(ciphertext, ".")
	if len(parts) != 2 || parts[0] != profile.Prefix {
		return nil, fmt.Errorf("%w: prefix", fmt.Errorf("invalid ciphertext"))
	}
	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("%w: encoding", fmt.Errorf("invalid ciphertext"))
	}
	if len(raw) != 1+profile.KeySize+8 {
		return nil, fmt.Errorf("%w: length", fmt.Errorf("invalid ciphertext"))
	}
	if raw[0] != profile.VersionByte {
		return nil, fmt.Errorf("%w: version", fmt.Errorf("invalid ciphertext"))
	}
	wrapped := raw[1 : 1+profile.KeySize]
	expected := checksum(profile, gateID, sessionDate, wrapped)
	for i := range expected {
		if expected[i] != raw[1+profile.KeySize+i] {
			return nil, fmt.Errorf("%w: checksum", fmt.Errorf("invalid ciphertext"))
		}
	}
	secret := xor(wrapped, DeriveKey(publicKey))
	if string(secret) != string(DeriveSecret(gateID, sessionDate)) {
		return nil, fmt.Errorf("%w: secret", fmt.Errorf("invalid ciphertext"))
	}
	return secret, nil
}

func ValidateParameters(profile Profile, value string) error {
	if profile.ID != value {
		return fmt.Errorf("parameter set mismatch")
	}
	return nil
}

func CiphertextInfo(ciphertext string) (string, int, error) {
	parts := strings.Split(ciphertext, ".")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("ciphertext format")
	}
	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", 0, err
	}
	return parts[0], len(raw), nil
}
