package store

import (
	"errors"
	"ticketgate/internal/domain"
)

type Stats struct{ Gates, Sessions, Requests, Audits int }

func (s *Store) Statistics() (Stats, error) {
	g, err := s.ListGates()
	if err != nil {
		return Stats{}, err
	}
	ss, err := s.ListSessions()
	if err != nil {
		return Stats{}, err
	}
	r, err := s.list(requestBucket)
	if err != nil {
		return Stats{}, err
	}
	a, err := s.ListAudits()
	if err != nil {
		return Stats{}, err
	}
	return Stats{Gates: len(g), Sessions: len(ss), Requests: len(r), Audits: len(a)}, nil
}

func (s *Store) HasGate(id string) bool { _, err := s.GetGate(id); return err == nil }

func (s *Store) HasSession(id string) bool { _, err := s.GetSession(id); return err == nil }

func (s *Store) SaveAll(g domain.Gate, session domain.GateSession, request domain.ValidationRequest, audit domain.AuditRecord) error {
	if err := s.SaveGate(g); err != nil {
		return err
	}
	if err := s.SaveSession(session); err != nil {
		return err
	}
	if err := s.SaveValidationRequest(request); err != nil {
		return err
	}
	if err := s.SaveAudit(audit); err != nil {
		return err
	}
	return nil
}

func (s *Store) RemoveAll(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	for _, operation := range []func() error{func() error { return s.DeleteGate(id) }, func() error { return s.DeleteSession(id) }, func() error { return s.DeleteValidationRequest(id) }, func() error { return s.DeleteAudit(id) }} {
		if err := operation(); err != nil && !errors.Is(err, ErrNotFound) {
			return err
		}
	}
	return nil
}
