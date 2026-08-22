package domain

import (
	"encoding/json"
)

func EncodeGate(g Gate) ([]byte, error) { return json.Marshal(g) }

func DecodeGate(data []byte) (Gate, error) {
	var value Gate
	if err := json.Unmarshal(data, &value); err != nil {
		return Gate{}, err
	}
	return value, nil
}

func EncodeSession(s GateSession) ([]byte, error) { return json.Marshal(s) }

func DecodeSession(data []byte) (GateSession, error) {
	var value GateSession
	if err := json.Unmarshal(data, &value); err != nil {
		return GateSession{}, err
	}
	return value, nil
}

func EncodeRequest(r ValidationRequest) ([]byte, error) { return json.Marshal(r) }

func DecodeRequest(data []byte) (ValidationRequest, error) {
	var value ValidationRequest
	if err := json.Unmarshal(data, &value); err != nil {
		return ValidationRequest{}, err
	}
	return value, nil
}

func EncodeAudit(a AuditRecord) ([]byte, error) { return json.Marshal(a) }

func DecodeAudit(data []byte) (AuditRecord, error) {
	var value AuditRecord
	if err := json.Unmarshal(data, &value); err != nil {
		return AuditRecord{}, err
	}
	return value, nil
}
