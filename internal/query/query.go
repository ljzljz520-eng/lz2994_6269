package query

import (
	"sort"
	"strings"
	"ticketgate/internal/domain"
	"ticketgate/internal/store"
)

type Manager struct{ Store *store.Store }

func New(st *store.Store) *Manager { return &Manager{Store: st} }

func (m *Manager) ListSessions(gateID, state string) ([]domain.GateSession, error) {
	items, err := m.Store.ListSessions()
	if err != nil {
		return nil, err
	}
	filtered := make([]domain.GateSession, 0, len(items))
	for _, item := range items {
		if gateID != "" && !strings.EqualFold(item.GateID, gateID) {
			continue
		}
		if state != "" && string(item.State) != state {
			continue
		}
		filtered = append(filtered, item)
	}
	sort.Slice(filtered, func(i, j int) bool { return filtered[i].ID < filtered[j].ID })
	return filtered, nil
}

func (m *Manager) GetAudit(id string) (domain.AuditRecord, error) { return m.Store.GetAudit(id) }

func (m *Manager) AuditsForGate(gateID string) ([]domain.AuditRecord, error) {
	items, err := m.Store.ListAudits()
	if err != nil {
		return nil, err
	}
	result := make([]domain.AuditRecord, 0, len(items))
	for _, item := range items {
		if strings.EqualFold(item.GateID, gateID) {
			result = append(result, item)
		}
	}
	return result, nil
}

func (m *Manager) CountByResult(gateID, result string) (int, error) {
	items, err := m.AuditsForGate(gateID)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, item := range items {
		if item.Result == result {
			count++
		}
	}
	return count, nil
}
