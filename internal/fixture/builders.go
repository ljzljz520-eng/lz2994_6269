package fixture

import (
	"fmt"
	"ticketgate/internal/crypto"
	"ticketgate/internal/domain"
)

func GateWithID(id string) domain.Gate {
	return domain.Gate{ID: id, Name: "fixture gate " + id, PublicKey: "public-" + id, Status: domain.GateActive, CreatedAt: SessionDate}
}

func Gates(count int) []domain.Gate {
	result := make([]domain.Gate, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, GateWithID(fmt.Sprintf("G-%02d", i+1)))
	}
	return result
}

func RequestFor(gateID, sessionID, payload string) domain.ValidationRequest {
	return domain.ValidationRequest{ID: "request-" + gateID + "-" + sessionID, GateID: gateID, SessionID: sessionID, Payload: payload, ParamSet: crypto.ProfileID, CreatedAt: SessionDate}
}

func RequestsForSession(gateID, sessionID string, payloads []string) []domain.ValidationRequest {
	result := make([]domain.ValidationRequest, 0, len(payloads))
	for _, payload := range payloads {
		result = append(result, RequestFor(gateID, sessionID, payload))
	}
	return result
}

func AuditFixture(event, gateID, sessionID, result string) domain.AuditRecord {
	return domain.AuditRecord{ID: domain.BuildAuditID(event, sessionID), GateID: gateID, SessionID: sessionID, Event: event, Result: result, Detail: "fixture", CreatedAt: SessionDate}
}
