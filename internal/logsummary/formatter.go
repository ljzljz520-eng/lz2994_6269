package logsummary

import (
	"fmt"
	"ticketgate/internal/domain"
)

type Formatter struct {
	IncludeDetail bool
	Prefix        string
}

func DefaultFormatter() Formatter { return Formatter{IncludeDetail: true, Prefix: "ticket-gate"} }

func (f Formatter) Format(a domain.AuditRecord) string {
	base := fmt.Sprintf("%s %s %s %s", f.Prefix, a.Event, a.GateID, a.Result)
	if f.IncludeDetail && a.Detail != "" {
		return base + " " + a.Detail
	}
	return base
}

func (f Formatter) FormatAll(records []domain.AuditRecord) []string {
	result := make([]string, 0, len(records))
	for _, record := range records {
		result = append(result, f.Format(record))
	}
	return result
}

func (f Formatter) WithDetail(enabled bool) Formatter { f.IncludeDetail = enabled; return f }

func FormatOutcome(outcome domain.Outcome) string {
	if err := domain.ValidateOutcome(outcome); err != nil {
		return "invalid"
	}
	return string(outcome)
}
