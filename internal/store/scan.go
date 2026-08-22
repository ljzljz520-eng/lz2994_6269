package store

import (
	"strings"
	"ticketgate/internal/domain"
)

func (s *Store) FindSessionsByDate(date string) ([]domain.GateSession, error) {
	items, err := s.ListSessions()
	if err != nil {
		return nil, err
	}
	result := make([]domain.GateSession, 0)
	for _, item := range items {
		if item.SessionDate == date {
			result = append(result, item)
		}
	}
	return result, nil
}

func (s *Store) FindSessionsByGate(gateID string) ([]domain.GateSession, error) {
	items, err := s.ListSessions()
	if err != nil {
		return nil, err
	}
	result := make([]domain.GateSession, 0)
	for _, item := range items {
		if strings.EqualFold(item.GateID, gateID) {
			result = append(result, item)
		}
	}
	return result, nil
}

func (s *Store) FindRequestsByGate(gateID string) ([]domain.ValidationRequest, error) {
	data, err := s.list(requestBucket)
	if err != nil {
		return nil, err
	}
	result := make([]domain.ValidationRequest, 0)
	for _, raw := range data {
		value, decodeErr := domain.DecodeRequest(raw)
		if decodeErr != nil {
			return nil, decodeErr
		}
		if strings.EqualFold(value.GateID, gateID) {
			result = append(result, value)
		}
	}
	return result, nil
}

func (s *Store) FindAuditsByEvent(event string) ([]domain.AuditRecord, error) {
	items, err := s.ListAudits()
	if err != nil {
		return nil, err
	}
	result := make([]domain.AuditRecord, 0)
	for _, item := range items {
		if item.Event == event {
			result = append(result, item)
		}
	}
	return result, nil
}
