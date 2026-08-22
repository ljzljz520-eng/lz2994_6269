package logsummary

import (
	"encoding/json"
	"sort"
	"strings"
	"ticketgate/internal/domain"
)

func JSONLine(a domain.AuditRecord) string {
	data, _ := json.Marshal(struct{ Event, Gate, Session, Result string }{a.Event, a.GateID, a.SessionID, a.Result})
	return string(data)
}

func Table(records []domain.AuditRecord) string {
	copyRecords := append([]domain.AuditRecord(nil), records...)
	sort.Slice(copyRecords, func(i, j int) bool { return copyRecords[i].ID < copyRecords[j].ID })
	rows := make([]string, 0, len(copyRecords))
	for _, record := range copyRecords {
		rows = append(rows, strings.Join([]string{record.Event, record.GateID, record.SessionID, record.Result}, "\t"))
	}
	return strings.Join(rows, "\n")
}

func ResultCode(a domain.AuditRecord) string {
	switch a.Result {
	case "validated":
		return "OK"
	case "negotiated":
		return "READY"
	case "37":
		return "MISMATCH"
	default:
		return "ERR"
	}
}
