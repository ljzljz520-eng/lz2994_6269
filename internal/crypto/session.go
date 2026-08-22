package crypto

import (
	"fmt"
	"strings"
)

type SessionEnvelope struct{ GateID, Date, Ciphertext, ProfileID string }

func NewEnvelope(gateID, date, publicKey string, profile Profile) (SessionEnvelope, error) {
	ciphertext, err := WrapSecret(publicKey, gateID, date, profile)
	if err != nil {
		return SessionEnvelope{}, err
	}
	return SessionEnvelope{GateID: gateID, Date: date, Ciphertext: ciphertext, ProfileID: profile.ID}, nil
}

func (e SessionEnvelope) Verify(publicKey string, profile Profile) error {
	if e.ProfileID != profile.ID {
		return fmt.Errorf("profile mismatch")
	}
	if strings.TrimSpace(e.GateID) == "" || strings.TrimSpace(e.Date) == "" {
		return fmt.Errorf("envelope identity missing")
	}
	return VerifyWrapped(e.Ciphertext, publicKey, e.GateID, e.Date, profile)
}

func (e SessionEnvelope) IsEmpty() bool { return e.GateID == "" && e.Date == "" && e.Ciphertext == "" }

func (e SessionEnvelope) Summary() string {
	return e.GateID + "|" + e.Date + "|" + Proof(e.Ciphertext, e.GateID)
}

func CompareEnvelopes(a, b SessionEnvelope) bool {
	return a.GateID == b.GateID && a.Date == b.Date && a.Ciphertext == b.Ciphertext && a.ProfileID == b.ProfileID
}
