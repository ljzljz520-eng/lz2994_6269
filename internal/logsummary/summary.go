package logsummary

import (
	"fmt"
	"strings"
	"ticketgate/internal/domain"
)

func Summary(a domain.AuditRecord) string {
	return fmt.Sprintf("event=%s gate=%s session=%s result=%s detail=%s", a.Event, a.GateID, a.SessionID, a.Result, a.Detail)
}

func Compact(a domain.AuditRecord) string {
	parts := []string{a.Event, a.GateID, a.SessionID, a.Result}
	return strings.Join(parts, "|")
}

func Lines(records []domain.AuditRecord) []string {
	lines := make([]string, 0, len(records))
	for _, record := range records {
		lines = append(lines, Summary(record))
	}
	return lines
}

func IsSuccessful(a domain.AuditRecord) bool {
	return a.Result == "negotiated" || a.Result == "validated"
}
