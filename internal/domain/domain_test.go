package domain

import "testing"

func TestDomainValidation(t *testing.T) {
	if err := ValidateGate(Gate{ID: "G-34", Name: "gate", PublicKey: "key", Status: GateActive}); err != nil {
		t.Fatal(err)
	}
	if ValidateGate(Gate{}) == nil {
		t.Fatal("expected gate error")
	}
	if err := RequireParameterSet(ParameterSetV1); err != nil {
		t.Fatal(err)
	}
	if RequireParameterSet("wrong") == nil {
		t.Fatal("expected parameter error")
	}
}

func TestSessionTransitions(t *testing.T) {
	s := GateSession{ID: "s", GateID: "g", SessionDate: "d", Ciphertext: "c", State: SessionActive, ParameterID: ParameterSetV1}
	if err := TransitionSession(&s, SessionClosed); err != nil {
		t.Fatal(err)
	}
	if s.State != SessionClosed {
		t.Fatal("state not closed")
	}
}
