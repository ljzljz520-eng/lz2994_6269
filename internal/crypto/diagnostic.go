package crypto

import (
	"fmt"
	"strings"
)

type CipherReport struct {
	Prefix       string
	EncodedBytes int
	Profile      string
	Valid        bool
	Reason       string
}

func ExplainCiphertext(ciphertext string, profile Profile) CipherReport {
	prefix, size, err := CiphertextInfo(ciphertext)
	report := CipherReport{Prefix: prefix, EncodedBytes: size, Profile: profile.ID}
	if err != nil {
		report.Reason = err.Error()
		return report
	}
	if prefix != profile.Prefix {
		report.Reason = "prefix mismatch"
		return report
	}
	if size != 1+profile.KeySize+8 {
		report.Reason = "length mismatch"
		return report
	}
	report.Valid = true
	report.Reason = "shape accepted"
	return report
}

func VerifyWrapped(ciphertext, publicKey, gateID, date string, profile Profile) error {
	if strings.TrimSpace(publicKey) == "" {
		return fmt.Errorf("public key is empty")
	}
	if _, err := UnwrapSecret(ciphertext, publicKey, gateID, date, profile); err != nil {
		return err
	}
	return nil
}

func Proof(ciphertext, gateID string) string {
	return Fingerprint([]byte(ciphertext + "|" + gateID))
}
