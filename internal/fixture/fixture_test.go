package fixture

import "testing"

func TestFixtures(t *testing.T) {
	if GateG34().ID != GateID {
		t.Fatal("gate id mismatch")
	}
	if len(ScenarioNames()) != 3 {
		t.Fatal("scenario count")
	}
	if RequestG34(SessionID()).ParamSet == "" {
		t.Fatal("parameter missing")
	}
}
