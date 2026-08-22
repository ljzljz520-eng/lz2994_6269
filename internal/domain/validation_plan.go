package domain

import "fmt"

type ValidationPlan struct {
	GateID      string
	SessionID   string
	ParameterID string
	Checks      []string
}

func NewValidationPlan(gateID, sessionID, parameterID string) (ValidationPlan, error) {
	if gateID == "" || sessionID == "" {
		return ValidationPlan{}, fmt.Errorf("plan identities are required")
	}
	if parameterID == "" {
		return ValidationPlan{}, ErrInvalidParameter
	}
	return ValidationPlan{GateID: NormalizeGateID(gateID), SessionID: sessionID, ParameterID: parameterID, Checks: []string{"gate", "session", "parameter", "ciphertext", "payload"}}, nil
}

func (p ValidationPlan) HasCheck(name string) bool {
	for _, check := range p.Checks {
		if check == name {
			return true
		}
	}
	return false
}

func (p ValidationPlan) Complete() bool {
	for _, required := range []string{"gate", "session", "parameter", "ciphertext", "payload"} {
		if !p.HasCheck(required) {
			return false
		}
	}
	return true
}

func (p ValidationPlan) StepCount() int { return len(p.Checks) }

func BuildValidationPlan(r ValidationRequest) (ValidationPlan, error) {
	return NewValidationPlan(r.GateID, r.SessionID, r.ParamSet)
}
