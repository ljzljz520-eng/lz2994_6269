package store

import (
	"strings"
	"ticketgate/internal/domain"
)

func (s *Store) SaveSessionBundle(g domain.Gate, session domain.GateSession, request domain.ValidationRequest, audit domain.AuditRecord) error {
	return s.SaveAll(g, session, request, audit)
}

func (s *Store) DeleteByGate(gateID string) error {
	sessions, err := s.FindSessionsByGate(gateID)
	if err != nil {
		return err
	}
	requests, err := s.FindRequestsByGate(gateID)
	if err != nil {
		return err
	}
	audits, err := s.ListAudits()
	if err != nil {
		return err
	}
	for _, item := range sessions {
		if err := s.DeleteSession(item.ID); err != nil {
			return err
		}
	}
	for _, item := range requests {
		if err := s.DeleteValidationRequest(item.ID); err != nil {
			return err
		}
	}
	for _, item := range audits {
		if strings.EqualFold(item.GateID, gateID) {
			if err := s.DeleteAudit(item.ID); err != nil {
				return err
			}
		}
	}
	return s.DeleteGate(gateID)
}

func (s *Store) CountSessionsForGate(gateID string) (int, error) {
	items, err := s.FindSessionsByGate(gateID)
	return len(items), err
}

func (s *Store) CountAuditsForGate(gateID string) (int, error) {
	items, err := s.ListAudits()
	if err != nil {
		return 0, err
	}
	count := 0
	for _, item := range items {
		if strings.EqualFold(item.GateID, gateID) {
			count++
		}
	}
	return count, nil
}

func (s *Store) ReplaceAudit(audit domain.AuditRecord) error { return s.SaveAudit(audit) }
