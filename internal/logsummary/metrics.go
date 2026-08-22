package logsummary

import (
	"strconv"
	"ticketgate/internal/domain"
)

type Metrics struct{ Total, Accepted, Rejected, Errors int }

func Measure(records []domain.AuditRecord) Metrics {
	result := Metrics{Total: len(records)}
	for _, record := range records {
		switch domain.AuditOutcome(record) {
		case domain.OutcomeAccepted:
			result.Accepted++
		case domain.OutcomeRejected:
			result.Rejected++
		default:
			result.Errors++
		}
	}
	return result
}

func (m Metrics) SuccessRate() float64 {
	if m.Total == 0 {
		return 0
	}
	return float64(m.Accepted) / float64(m.Total)
}

func (m Metrics) String() string {
	return "total=" + strconv.Itoa(m.Total) + " accepted=" + strconv.Itoa(m.Accepted) + " rejected=" + strconv.Itoa(m.Rejected) + " errors=" + strconv.Itoa(m.Errors)
}

func SummarizeResults(records []domain.AuditRecord) string { return Measure(records).String() }
