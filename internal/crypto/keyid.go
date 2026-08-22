package crypto

import (
	"fmt"
	"strings"
)

func KeyID(gateID, publicKey string) string {
	return strings.ToUpper(strings.TrimSpace(gateID)) + ":" + Fingerprint([]byte(publicKey))
}

func KeyIDParts(value string) (string, string, error) {
	parts := strings.Split(value, ":")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid key id")
	}
	return parts[0], parts[1], nil
}

func SameKeyID(left, right string) bool {
	return strings.EqualFold(strings.TrimSpace(left), strings.TrimSpace(right))
}

func IsProfileCompatible(profile Profile, value string) bool {
	return profile.ID == strings.TrimSpace(value)
}
