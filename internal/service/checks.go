package service

import (
	"ticketgate/internal/crypto"
	"ticketgate/internal/domain"
)

type CheckResult struct {
	Name   string
	Passed bool
	Detail string
}

func (s *Service) RunChecks(req domain.ValidationRequest) ([]CheckResult, error) {
	if err := s.validateReady(); err != nil {
		return nil, err
	}
	gate, err := s.Store.GetGate(req.GateID)
	if err != nil {
		return nil, err
	}
	session, err := s.Store.GetSession(req.SessionID)
	if err != nil {
		return nil, err
	}
	policy := domain.DefaultGatePolicy()
	checks := []CheckResult{{Name: "gate", Passed: policy.ValidateGate(gate) == nil, Detail: gate.ID}, {Name: "session", Passed: domain.IsSessionValidFor(session, gate.ID), Detail: string(session.State)}, {Name: "parameter", Passed: crypto.ValidateParameters(s.Profile, req.ParamSet) == nil, Detail: req.ParamSet}}
	_, unwrapErr := crypto.UnwrapSecret(session.Ciphertext, gate.PublicKey, gate.ID, session.SessionDate, s.Profile)
	checks = append(checks, CheckResult{Name: "ciphertext", Passed: unwrapErr == nil, Detail: "wrapped-secret"})
	checks = append(checks, CheckResult{Name: "payload", Passed: s.ValidatePayload(req) == nil, Detail: req.Payload})
	return checks, nil
}

func ChecksPassed(checks []CheckResult) bool {
	for _, check := range checks {
		if !check.Passed {
			return false
		}
	}
	return true
}

func FailedChecks(checks []CheckResult) []string {
	result := make([]string, 0)
	for _, check := range checks {
		if !check.Passed {
			result = append(result, check.Name)
		}
	}
	return result
}

func (s *Service) Decision(req domain.ValidationRequest) (domain.AccessDecision, error) {
	if err := s.validateReady(); err != nil {
		return domain.AccessDecision{}, err
	}
	gate, err := s.Store.GetGate(req.GateID)
	if err != nil {
		return domain.AccessDecision{}, err
	}
	session, err := s.Store.GetSession(req.SessionID)
	if err != nil {
		return domain.AccessDecision{}, err
	}
	return domain.DecideGateAccess(gate, session, req, domain.DefaultGatePolicy()), nil
}
