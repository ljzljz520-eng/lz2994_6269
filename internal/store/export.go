package store

import (
	"encoding/json"
	"ticketgate/internal/domain"
)

type Snapshot struct {
	Gates    []domain.Gate
	Sessions []domain.GateSession
	Requests []domain.ValidationRequest
	Audits   []domain.AuditRecord
}

func (s *Store) Export() (Snapshot, error) {
	gates, err := s.ListGates()
	if err != nil {
		return Snapshot{}, err
	}
	sessions, err := s.ListSessions()
	if err != nil {
		return Snapshot{}, err
	}
	requestsRaw, err := s.list(requestBucket)
	if err != nil {
		return Snapshot{}, err
	}
	audits, err := s.ListAudits()
	if err != nil {
		return Snapshot{}, err
	}
	requests := make([]domain.ValidationRequest, 0, len(requestsRaw))
	for _, raw := range requestsRaw {
		value, decodeErr := domain.DecodeRequest(raw)
		if decodeErr != nil {
			return Snapshot{}, decodeErr
		}
		requests = append(requests, value)
	}
	return Snapshot{Gates: gates, Sessions: sessions, Requests: requests, Audits: audits}, nil
}

func EncodeSnapshot(snapshot Snapshot) ([]byte, error) { return json.Marshal(snapshot) }

func DecodeSnapshot(data []byte) (Snapshot, error) {
	var value Snapshot
	err := json.Unmarshal(data, &value)
	return value, err
}

func (s *Store) Import(snapshot Snapshot) error {
	for _, gate := range snapshot.Gates {
		if err := s.SaveGate(gate); err != nil {
			return err
		}
	}
	for _, session := range snapshot.Sessions {
		if err := s.SaveSession(session); err != nil {
			return err
		}
	}
	for _, request := range snapshot.Requests {
		if err := s.SaveValidationRequest(request); err != nil {
			return err
		}
	}
	for _, audit := range snapshot.Audits {
		if err := s.SaveAudit(audit); err != nil {
			return err
		}
	}
	return nil
}
