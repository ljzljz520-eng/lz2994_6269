package service

import (
	"ticketgate/internal/domain"
)

type BatchSummary struct {
	Total, Valid, Invalid int
	FailedIDs             []string
}

func (s *Service) ValidateBatchSummary(requests []domain.ValidationRequest) (BatchSummary, error) {
	if err := s.validateReady(); err != nil {
		return BatchSummary{}, err
	}
	result := BatchSummary{Total: len(requests), FailedIDs: make([]string, 0)}
	for _, request := range requests {
		_, err := s.ValidateGateSession(request)
		if err != nil {
			result.Invalid++
			result.FailedIDs = append(result.FailedIDs, request.ID)
			continue
		}
		result.Valid++
	}
	return result, nil
}

func (s *Service) RegisterAndNegotiate(gate domain.Gate) (domain.GateSession, domain.AuditRecord, error) {
	if _, err := s.RegisterGate(gate); err != nil {
		return domain.GateSession{}, domain.AuditRecord{}, err
	}
	return s.NegotiateSession(gate.ID)
}

func (s *Service) RequestExists(id string) bool {
	if s == nil || s.Store == nil {
		return false
	}
	_, err := s.Store.GetValidationRequest(id)
	return err == nil
}

func (s *Service) AuditExists(id string) bool {
	if s == nil || s.Store == nil {
		return false
	}
	_, err := s.Store.GetAudit(id)
	return err == nil
}

func (s *Service) GateExists(id string) bool {
	if s == nil || s.Store == nil {
		return false
	}
	return s.Store.HasGate(id)
}

func (s *Service) SessionExists(id string) bool {
	if s == nil || s.Store == nil {
		return false
	}
	return s.Store.HasSession(id)
}
