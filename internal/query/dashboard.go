package query

import (
	"sort"
	"ticketgate/internal/domain"
)

type Dashboard struct {
	Stats          map[string]int
	ActiveSessions []string
	Recent         []domain.AuditRecord
}

func (m *Manager) BuildDashboard(gateID string) (Dashboard, error) {
	sessions, err := m.ListSessions(gateID, string(domain.SessionActive))
	if err != nil {
		return Dashboard{}, err
	}
	audits, err := m.AuditsForGate(gateID)
	if err != nil {
		return Dashboard{}, err
	}
	stats := make(map[string]int)
	for _, audit := range audits {
		stats[audit.Result]++
	}
	active := make([]string, 0, len(sessions))
	for _, session := range sessions {
		active = append(active, session.ID)
	}
	if len(audits) > 5 {
		audits = audits[len(audits)-5:]
	}
	return Dashboard{Stats: stats, ActiveSessions: active, Recent: audits}, nil
}

func SortSessionsByDate(items []domain.GateSession) []domain.GateSession {
	result := append([]domain.GateSession(nil), items...)
	sort.Slice(result, func(i, j int) bool { return result[i].SessionDate < result[j].SessionDate })
	return result
}

func SessionStates(items []domain.GateSession) map[domain.SessionState]int {
	result := make(map[domain.SessionState]int)
	for _, item := range items {
		result[item.State]++
	}
	return result
}

func IsHealthyDashboard(d Dashboard) bool {
	return len(d.ActiveSessions) > 0 && d.Stats["negotiated"] >= 1
}
