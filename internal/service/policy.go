package service

import (
	"ticketgate/internal/domain"
)

type PolicyReport struct {
	GateID, RequestID string
	Decision          domain.AccessDecision
	Checks            []CheckResult
}

func (s *Service) EvaluatePolicy(req domain.ValidationRequest) (PolicyReport, error) {
	if err := s.validateReady(); err != nil {
		return PolicyReport{}, err
	}
	decision, err := s.Decision(req)
	if err != nil {
		return PolicyReport{}, err
	}
	checks, err := s.RunChecks(req)
	if err != nil {
		return PolicyReport{}, err
	}
	return PolicyReport{GateID: req.GateID, RequestID: req.ID, Decision: decision, Checks: checks}, nil
}

func (p PolicyReport) Passed() bool { return p.Decision.Allowed && ChecksPassed(p.Checks) }

func (p PolicyReport) FailedNames() []string { return FailedChecks(p.Checks) }

func (p PolicyReport) Score() int {
	score := 0
	for _, check := range p.Checks {
		if check.Passed {
			score++
		}
	}
	return score
}

func (p PolicyReport) RequiredScore() int { return len(p.Checks) }

func (p PolicyReport) Complete() bool { return p.RequiredScore() > 0 && p.Score() == p.RequiredScore() }

func (s *Service) PolicyAllows(req domain.ValidationRequest) (bool, error) {
	report, err := s.EvaluatePolicy(req)
	if err != nil {
		return false, err
	}
	return report.Passed(), nil
}

func (s *Service) PolicyScore(req domain.ValidationRequest) (int, error) {
	report, err := s.EvaluatePolicy(req)
	if err != nil {
		return 0, err
	}
	return report.Score(), nil
}
