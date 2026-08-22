package service

import (
	"fmt"
	"ticketgate/internal/domain"
)

func (s *Service) RecordManualAudit(gateID, sessionID, event, result, detail string) (domain.AuditRecord, error) {
	if err := s.validateReady(); err != nil {
		return domain.AuditRecord{}, err
	}
	if gateID == "" || sessionID == "" || event == "" {
		return domain.AuditRecord{}, fmt.Errorf("audit identity is required")
	}
	if result == "" {
		return domain.AuditRecord{}, fmt.Errorf("audit result is required")
	}
	audit := domain.AuditRecord{ID: domain.BuildAuditID(event, sessionID), GateID: domain.NormalizeGateID(gateID), SessionID: sessionID, Event: event, Result: result, Detail: detail, CreatedAt: s.CurrentDate()}
	if err := s.Store.SaveAudit(audit); err != nil {
		return domain.AuditRecord{}, err
	}
	return audit, nil
}

func (s *Service) AuditForSession(sessionID string) ([]domain.AuditRecord, error) {
	if err := s.validateReady(); err != nil {
		return nil, err
	}
	items, err := s.Store.ListAudits()
	if err != nil {
		return nil, err
	}
	result := make([]domain.AuditRecord, 0)
	for _, item := range items {
		if item.SessionID == sessionID {
			result = append(result, item)
		}
	}
	return result, nil
}

func (s *Service) AuditCount(sessionID string) (int, error) {
	items, err := s.AuditForSession(sessionID)
	if err != nil {
		return 0, err
	}
	return len(items), nil
}
