package service

import (
	"fmt"
	"ticketgate/internal/crypto"
	"ticketgate/internal/domain"
)

func (s *Service) NegotiateSession(gateID string) (domain.GateSession, domain.AuditRecord, error) {
	if err := s.validateReady(); err != nil {
		return domain.GateSession{}, domain.AuditRecord{}, err
	}
	gate, err := s.Store.GetGate(gateID)
	if err != nil {
		return domain.GateSession{}, domain.AuditRecord{}, fmt.Errorf("load gate: %w", err)
	}
	if !domain.IsGateUsable(gate) {
		return domain.GateSession{}, domain.AuditRecord{}, fmt.Errorf("gate %s is not usable", gate.ID)
	}
	date := s.CurrentDate()
	ciphertext, err := crypto.WrapSecret(gate.PublicKey, gate.ID, date, s.Profile)
	if err != nil {
		return domain.GateSession{}, domain.AuditRecord{}, err
	}
	session := domain.GateSession{ID: "session-" + gate.ID + "-" + date, GateID: gate.ID, SessionDate: date, Ciphertext: ciphertext, State: domain.SessionActive, CreatedAt: date, ParameterID: s.Profile.ID}
	if err := s.Store.SaveSession(session); err != nil {
		return domain.GateSession{}, domain.AuditRecord{}, fmt.Errorf("save session: %w", err)
	}
	audit := domain.AuditRecord{ID: domain.BuildAuditID(domain.NegotiationEvent, session.ID), GateID: gate.ID, SessionID: session.ID, Event: domain.NegotiationEvent, Result: "negotiated", Detail: crypto.Fingerprint([]byte(ciphertext)), CreatedAt: date}
	if err := s.Store.SaveAudit(audit); err != nil {
		return domain.GateSession{}, domain.AuditRecord{}, fmt.Errorf("save negotiation audit: %w", err)
	}
	return session, audit, nil
}

func (s *Service) ActivateSession(id string) (domain.GateSession, error) {
	if err := s.validateReady(); err != nil {
		return domain.GateSession{}, err
	}
	session, err := s.Store.GetSession(id)
	if err != nil {
		return domain.GateSession{}, err
	}
	if err := domain.TransitionSession(&session, domain.SessionActive); err != nil {
		return domain.GateSession{}, err
	}
	return session, s.Store.SaveSession(session)
}

func (s *Service) CloseSession(id string) error {
	if err := s.validateReady(); err != nil {
		return err
	}
	session, err := s.Store.GetSession(id)
	if err != nil {
		return err
	}
	if err := domain.TransitionSession(&session, domain.SessionClosed); err != nil {
		return err
	}
	return s.Store.SaveSession(session)
}
