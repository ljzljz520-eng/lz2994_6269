package logsummary

import (
	"strings"
	"testing"
	"ticketgate/internal/domain"
)

func TestLogSummary(t *testing.T) {
	a := domain.AuditRecord{Event: "event", GateID: "G-34", SessionID: "s", Result: "validated", Detail: "ok"}
	if !strings.Contains(Summary(a), "validated") {
		t.Fatal("missing result")
	}
	if !IsSuccessful(a) {
		t.Fatal("expected success")
	}
}
