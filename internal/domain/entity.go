package domain

import (
	"errors"
	"fmt"
	"strings"
)

type GateStatus string

const (
	GateActive   GateStatus = "active"
	GateDisabled GateStatus = "disabled"
)

type SessionState string

const (
	SessionPrepared SessionState = "prepared"
	SessionActive   SessionState = "active"
	SessionClosed   SessionState = "closed"
)

type Gate struct {
	ID        string
	Name      string
	PublicKey string
	Status    GateStatus
	CreatedAt string
}

type GateSession struct {
	ID          string
	GateID      string
	SessionDate string
	Ciphertext  string
	State       SessionState
	CreatedAt   string
	ParameterID string
}

type ValidationRequest struct {
	ID        string
	GateID    string
	SessionID string
	Payload   string
	ParamSet  string
	CreatedAt string
}

type AuditRecord struct {
	ID        string
	GateID    string
	SessionID string
	Event     string
	Result    string
	Detail    string
	CreatedAt string
}

var (
	ErrMissingGate      = errors.New("gate is missing")
	ErrMissingSession   = errors.New("session is missing")
	ErrMissingRequest   = errors.New("request is missing")
	ErrInvalidParameter = errors.New("parameter set is invalid")
	ErrInvalidCipher    = errors.New("ciphertext is invalid")
	ErrFileMissing      = errors.New("persistence file is missing")
)

func ValidateGate(g Gate) error {
	if strings.TrimSpace(g.ID) == "" {
		return fmt.Errorf("%w: id", ErrMissingGate)
	}
	if strings.TrimSpace(g.Name) == "" {
		return fmt.Errorf("%w: name", ErrMissingGate)
	}
	if strings.TrimSpace(g.PublicKey) == "" {
		return fmt.Errorf("%w: public key", ErrMissingGate)
	}
	if g.Status == "" {
		return fmt.Errorf("%w: status", ErrMissingGate)
	}
	return nil
}

func ValidateSession(s GateSession) error {
	if s.ID == "" || s.GateID == "" {
		return ErrMissingSession
	}
	if s.SessionDate == "" || s.Ciphertext == "" {
		return ErrMissingSession
	}
	if s.State != SessionPrepared && s.State != SessionActive && s.State != SessionClosed {
		return fmt.Errorf("%w: state", ErrMissingSession)
	}
	if s.ParameterID == "" {
		return fmt.Errorf("%w: parameter", ErrMissingSession)
	}
	return nil
}

func ValidateRequest(r ValidationRequest) error {
	if strings.TrimSpace(r.ID) == "" || strings.TrimSpace(r.GateID) == "" || strings.TrimSpace(r.SessionID) == "" {
		return ErrMissingRequest
	}
	if strings.TrimSpace(r.Payload) == "" {
		return fmt.Errorf("%w: payload", ErrMissingRequest)
	}
	if strings.TrimSpace(r.ParamSet) == "" {
		return fmt.Errorf("%w: parameter", ErrMissingRequest)
	}
	return nil
}

func ValidateAudit(a AuditRecord) error {
	if a.ID == "" || a.GateID == "" || a.SessionID == "" {
		return errors.New("audit identity is required")
	}
	if a.Event == "" || a.Result == "" {
		return errors.New("audit event and result are required")
	}
	return nil
}
