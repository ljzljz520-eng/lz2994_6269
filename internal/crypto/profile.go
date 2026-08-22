package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const ProfileID = "gate-session-v1"

type Profile struct {
	ID          string
	KeySize     int
	Digest      string
	Prefix      string
	VersionByte byte
}

func DefaultProfile() Profile {
	return Profile{ID: ProfileID, KeySize: 32, Digest: "sha256", Prefix: "GSS1", VersionByte: 1}
}

func (p Profile) Validate() error {
	if p.ID == "" || p.KeySize < 16 || p.Digest != "sha256" || p.Prefix == "" {
		return fmt.Errorf("invalid crypto profile")
	}
	return nil
}

func DeriveSecret(gateID, sessionDate string) []byte {
	h := sha256.Sum256([]byte("ticket-gate|" + gateID + "|" + sessionDate + "|secret"))
	result := make([]byte, len(h))
	copy(result, h[:])
	return result
}

func DeriveKey(publicKey string) []byte {
	h := sha256.Sum256([]byte("ticket-gate|public-key|" + publicKey))
	result := make([]byte, len(h))
	copy(result, h[:])
	return result
}

func Fingerprint(value []byte) string {
	h := sha256.Sum256(value)
	return hex.EncodeToString(h[:8])
}
