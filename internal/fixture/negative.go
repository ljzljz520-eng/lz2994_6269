package fixture

import "ticketgate/internal/domain"

func TamperedRequest(sessionID string) domain.ValidationRequest {
	r := RequestG34(sessionID)
	r.Payload = "ticket:37:tampered"
	return r
}

func MissingParameterRequest(sessionID string) domain.ValidationRequest {
	r := RequestG34(sessionID)
	r.ParamSet = ""
	return r
}

func ClosedSession(id string) domain.GateSession {
	return domain.GateSession{ID: id, GateID: GateID, SessionDate: SessionDate, Ciphertext: "GSS1.closed", State: domain.SessionClosed, CreatedAt: SessionDate, ParameterID: "gate-session-v1"}
}

func EmptyGate() domain.Gate { return domain.Gate{} }

func EmptyRequest() domain.ValidationRequest { return domain.ValidationRequest{} }

func ErrorCases(sessionID string) []domain.ValidationRequest {
	return []domain.ValidationRequest{TamperedRequest(sessionID), MissingParameterRequest(sessionID), EmptyRequest()}
}
