package query

import (
	"strings"
	"ticketgate/internal/domain"
)

type AuditFilter struct{ GateID, Event, Result, SessionID string }

func (m *Manager) FilterAudits(filter AuditFilter) ([]domain.AuditRecord, error) {
	items, err := m.Store.ListAudits()
	if err != nil {
		return nil, err
	}
	result := make([]domain.AuditRecord, 0)
	for _, item := range items {
		if filter.GateID != "" && !strings.EqualFold(filter.GateID, item.GateID) {
			continue
		}
		if filter.Event != "" && filter.Event != item.Event {
			continue
		}
		if filter.Result != "" && filter.Result != item.Result {
			continue
		}
		if filter.SessionID != "" && filter.SessionID != item.SessionID {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

func (m *Manager) HasResult(filter AuditFilter) (bool, error) {
	items, err := m.FilterAudits(filter)
	if err != nil {
		return false, err
	}
	return len(items) > 0, nil
}

func GroupByEvent(items []domain.AuditRecord) map[string][]domain.AuditRecord {
	result := make(map[string][]domain.AuditRecord)
	for _, item := range items {
		result[item.Event] = append(result[item.Event], item)
	}
	return result
}
