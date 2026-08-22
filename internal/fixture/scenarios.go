package fixture

import (
	"ticketgate/internal/domain"
)

type Scenario struct {
	Name      string
	Gate      domain.Gate
	Request   domain.ValidationRequest
	Expected  string
	Parameter string
}

func ValidScenario(sessionID string) Scenario {
	return Scenario{Name: "valid-g34", Gate: GateG34(), Request: RequestG34(sessionID), Expected: "validated", Parameter: "gate-session-v1"}
}

func InvalidCipherScenario(sessionID string) Scenario {
	r := RequestG34(sessionID)
	r.Payload = "tampered"
	return Scenario{Name: "bad-cipher", Gate: GateG34(), Request: r, Expected: "rejected", Parameter: "gate-session-v1"}
}

func ParameterScenario(sessionID string) Scenario {
	return Scenario{Name: "wrong-parameters", Gate: GateG34(), Request: WrongParameterRequest(sessionID), Expected: "rejected", Parameter: "gate-session-v0"}
}

func ScenarioNames() []string { return []string{"valid-g34", "bad-cipher", "wrong-parameters"} }
