package service

import (
	"ticketgate/internal/domain"
	"ticketgate/internal/query"
)

type ValidationReport struct{ RequestID, GateID, SessionID, Outcome, Fingerprint string }

func (s *Service) BuildReport(requestID string) (ValidationReport, error) {
	if err := s.validateReady(); err != nil {
		return ValidationReport{}, err
	}
	req, err := s.Store.GetValidationRequest(requestID)
	if err != nil {
		return ValidationReport{}, err
	}
	m := query.New(s.Store)
	audit, err := m.GetAudit(domain.BuildAuditID(domain.ValidationEvent, requestID))
	if err != nil {
		return ValidationReport{}, err
	}
	return ValidationReport{RequestID: req.ID, GateID: req.GateID, SessionID: req.SessionID, Outcome: string(domain.AuditOutcome(audit)), Fingerprint: domain.RequestFingerprint(req)}, nil
}

func (s *Service) ValidateBatch(requests []domain.ValidationRequest) []error {
	results := make([]error, len(requests))
	for i, request := range requests {
		_, err := s.ValidateGateSession(request)
		results[i] = err
	}
	return results
}

func CountErrors(errs []error) int {
	count := 0
	for _, err := range errs {
		if err != nil {
			count++
		}
	}
	return count
}
