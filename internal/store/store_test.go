package store

import (
	"path/filepath"
	"testing"
	"ticketgate/internal/domain"
)

func TestStoreEntities(t *testing.T) {
	path := filepath.Join(t.TempDir(), "store.db")
	s, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	g := domain.Gate{ID: "G-34", Name: "gate", PublicKey: "key", Status: domain.GateActive, CreatedAt: "d"}
	if err := s.SaveGate(g); err != nil {
		t.Fatal(err)
	}
	got, err := s.GetGate("g-34")
	if err != nil || got.ID != "G-34" {
		t.Fatalf("gate %#v %v", got, err)
	}
	if err := s.SaveValidationRequest(domain.ValidationRequest{ID: "r", GateID: "G-34", SessionID: "s", Payload: "p", ParamSet: "x"}); err != nil {
		t.Fatal(err)
	}
}
