package fixture

import (
	"fmt"
	"ticketgate/internal/crypto"
	"ticketgate/internal/domain"
)

const GateID = "G-34"
const PublicKey = "gate-public-key-g34"
const SessionDate = "2026-08-22"

func GateG34() domain.Gate {
	return domain.Gate{ID: GateID, Name: "North concourse gate 34", PublicKey: PublicKey, Status: domain.GateActive, CreatedAt: SessionDate}
}

func RequestG34(sessionID string) domain.ValidationRequest {
	return domain.ValidationRequest{ID: "request-G-34", GateID: GateID, SessionID: sessionID, Payload: "ticket:37:accepted", ParamSet: crypto.ProfileID, CreatedAt: SessionDate}
}

func SessionID() string { return fmt.Sprintf("session-%s-%s", GateID, SessionDate) }

func InvalidCiphertext() string { return "GSS1.not-a-valid-cipher" }

func WrongParameterRequest(sessionID string) domain.ValidationRequest {
	r := RequestG34(sessionID)
	r.ParamSet = "gate-session-v0"
	return r
}

func MissingFilePath(dir string) string { return dir + "/missing-ticket-gate.db" }
