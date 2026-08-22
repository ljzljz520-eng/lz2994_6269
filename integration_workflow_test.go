package ticketgate_test

import (
	"path/filepath"
	"testing"
	"ticketgate/internal/fixture"
	"ticketgate/internal/service"
	"ticketgate/internal/store"
)

func TestWorkflowRegisterGate(t *testing.T) {
	s, err := store.Open(filepath.Join(t.TempDir(), "workflow.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	svc := service.New(s, fixture.DefaultClock())
	g, err := svc.RegisterGate(fixture.GateG34())
	if err != nil || g.ID != fixture.GateID {
		t.Fatalf("gate=%#v err=%v", g, err)
	}
}

func TestWorkflowNegotiateSession(t *testing.T) {
	s, err := store.Open(filepath.Join(t.TempDir(), "workflow.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	svc := service.New(s, fixture.DefaultClock())
	if _, err := svc.RegisterGate(fixture.GateG34()); err != nil {
		t.Fatal(err)
	}
	ss, audit, err := svc.NegotiateSession(fixture.GateID)
	if err != nil || ss.State != "active" || audit.Result != "negotiated" {
		t.Fatalf("session=%#v audit=%#v err=%v", ss, audit, err)
	}
}
