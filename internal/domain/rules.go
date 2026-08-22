package domain

import (
	"fmt"
	"strings"
)

const (
	ParameterSetV1   = "gate-session-v1"
	SessionPrefix    = "GSS1"
	ValidationEvent  = "session.validation"
	NegotiationEvent = "session.negotiation"
)

func NormalizeGateID(id string) string {
	return strings.ToUpper(strings.TrimSpace(id))
}

func NormalizeParameterSet(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func IsGateUsable(g Gate) bool {
	return g.Status == GateActive && ValidateGate(g) == nil
}

func IsSessionValidFor(s GateSession, gateID string) bool {
	if ValidateSession(s) != nil {
		return false
	}
	if NormalizeGateID(s.GateID) != NormalizeGateID(gateID) {
		return false
	}
	return s.State == SessionActive || s.State == SessionPrepared
}

func RequireParameterSet(value string) error {
	if NormalizeParameterSet(value) != ParameterSetV1 {
		return fmt.Errorf("%w: got %q", ErrInvalidParameter, value)
	}
	return nil
}

func TransitionSession(s *GateSession, next SessionState) error {
	if s == nil {
		return ErrMissingSession
	}
	if ValidateSession(*s) != nil {
		return ErrMissingSession
	}
	if s.State == SessionClosed {
		return fmt.Errorf("session %s is closed", s.ID)
	}
	if next == SessionClosed && s.State == SessionPrepared {
		return fmt.Errorf("session %s must be active before close", s.ID)
	}
	if next != SessionPrepared && next != SessionActive && next != SessionClosed {
		return fmt.Errorf("unknown session state %q", next)
	}
	s.State = next
	return nil
}

func BuildAuditID(event, subject string) string {
	return fmt.Sprintf("%s:%s", event, subject)
}
