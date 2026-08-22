package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type GatePolicy struct {
	RequireActiveGate bool
	AllowedParameters []string
	MaxPayloadLength  int
	RequireTicketTag  bool
}

func DefaultGatePolicy() GatePolicy {
	return GatePolicy{RequireActiveGate: true, AllowedParameters: []string{ParameterSetV1}, MaxPayloadLength: 256, RequireTicketTag: true}
}

func (p GatePolicy) ValidateGate(g Gate) error {
	if err := ValidateGate(g); err != nil {
		return err
	}
	if p.RequireActiveGate && g.Status != GateActive {
		return fmt.Errorf("gate %s is not active", g.ID)
	}
	return nil
}

func (p GatePolicy) ValidateRequest(r ValidationRequest) error {
	if err := ValidateRequest(r); err != nil {
		return err
	}
	if p.MaxPayloadLength > 0 && len(r.Payload) > p.MaxPayloadLength {
		return fmt.Errorf("payload exceeds policy limit")
	}
	if p.RequireTicketTag && !strings.Contains(r.Payload, "ticket:") {
		return fmt.Errorf("ticket tag is required")
	}
	for _, allowed := range p.AllowedParameters {
		if NormalizeParameterSet(r.ParamSet) == NormalizeParameterSet(allowed) {
			return nil
		}
	}
	return ErrInvalidParameter
}

func RequestFingerprint(r ValidationRequest) string {
	h := sha256.Sum256([]byte(r.ID + "|" + r.GateID + "|" + r.SessionID + "|" + r.Payload + "|" + r.ParamSet))
	return hex.EncodeToString(h[:])
}

func AuditOutcome(a AuditRecord) Outcome {
	switch a.Result {
	case "validated", "negotiated":
		return OutcomeAccepted
	case "rejected":
		return OutcomeRejected
	default:
		return OutcomeError
	}
}
