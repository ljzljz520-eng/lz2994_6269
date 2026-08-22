package query

import (
	"sort"
	"strings"
	"ticketgate/internal/domain"
	"ticketgate/internal/store"
)

type TimelineEntry struct{ At, Event, Result, Subject string }

func (m *Manager) Timeline(gateID string) ([]TimelineEntry, error) {
	items, err := m.Store.ListAudits()
	if err != nil {
		return nil, err
	}
	result := make([]TimelineEntry, 0)
	for _, item := range items {
		if gateID == "" || strings.EqualFold(item.GateID, gateID) {
			result = append(result, TimelineEntry{At: item.CreatedAt, Event: item.Event, Result: item.Result, Subject: item.SessionID})
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].At < result[j].At })
	return result, nil
}

func (m *Manager) LatestForGate(gateID string) (domain.AuditRecord, error) {
	items, err := m.AuditsForGate(gateID)
	if err != nil {
		return domain.AuditRecord{}, err
	}
	if len(items) == 0 {
		return domain.AuditRecord{}, store.ErrNotFound
	}
	return items[len(items)-1], nil
}

func (m *Manager) Results(gateID string) (map[string]int, error) {
	items, err := m.AuditsForGate(gateID)
	if err != nil {
		return nil, err
	}
	result := make(map[string]int)
	for _, item := range items {
		result[item.Result]++
	}
	return result, nil
}
