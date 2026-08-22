package domain

import "fmt"

type Outcome string

const (
	OutcomeAccepted Outcome = "accepted"
	OutcomeRejected Outcome = "rejected"
	OutcomeError    Outcome = "error"
)

func (o Outcome) IsTerminal() bool {
	return o == OutcomeAccepted || o == OutcomeRejected || o == OutcomeError
}

func SessionDisplayState(state SessionState) string {
	switch state {
	case SessionPrepared:
		return "prepared"
	case SessionActive:
		return "active"
	case SessionClosed:
		return "closed"
	default:
		return "unknown"
	}
}

func ValidateOutcome(value Outcome) error {
	if !value.IsTerminal() {
		return fmt.Errorf("unsupported outcome %q", value)
	}
	return nil
}

func (s SessionState) CanValidate() bool {
	return s == SessionPrepared || s == SessionActive
}
