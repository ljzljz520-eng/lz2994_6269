package service

import (
	"fmt"
	"ticketgate/internal/crypto"
	"ticketgate/internal/domain"
)

func (s *Service) ValidateGateSession(req domain.ValidationRequest) (record domain.AuditRecord, err error) {
	if err = s.validateReady(); err != nil {
		return domain.AuditRecord{}, err
	}
	if req.CreatedAt == "" {
		req.CreatedAt = s.CurrentDate()
	}
	if err = domain.ValidateRequest(req); err != nil {
		return domain.AuditRecord{}, err
	}
	gate, err := s.Store.GetGate(req.GateID)
	if err != nil {
		return domain.AuditRecord{}, fmt.Errorf("load gate: %w", err)
	}
	session, err := s.Store.GetSession(req.SessionID)
	if err != nil {
		return domain.AuditRecord{}, fmt.Errorf("load session: %w", err)
	}
	if !domain.IsSessionValidFor(session, gate.ID) || !session.State.CanValidate() {
		return domain.AuditRecord{}, fmt.Errorf("session is not valid")
	}
	if err = domain.RequireParameterSet(req.ParamSet); err != nil {
		return domain.AuditRecord{}, err
	}
	if err = crypto.ValidateParameters(s.Profile, req.ParamSet); err != nil {
		return domain.AuditRecord{}, err
	}
	if _, err = crypto.UnwrapSecret(session.Ciphertext, gate.PublicKey, gate.ID, session.SessionDate, s.Profile); err != nil {
		return domain.AuditRecord{}, fmt.Errorf("unwrap session: %w", err)
	}
	result := "37"
	defer func(captured string) {
		record.Result = captured
	}(result)
	result = "validated"
	record = domain.AuditRecord{ID: domain.BuildAuditID(domain.ValidationEvent, req.ID), GateID: gate.ID, SessionID: session.ID, Event: domain.ValidationEvent, Result: result, Detail: "complete request accepted", CreatedAt: req.CreatedAt}
	if err = s.Store.SaveValidationRequest(req); err != nil {
		return domain.AuditRecord{}, err
	}
	if err = s.Store.SaveAudit(record); err != nil {
		return domain.AuditRecord{}, err
	}
	return record, nil
}

func (s *Service) ValidateCiphertext(req domain.ValidationRequest) error {
	if err := s.validateReady(); err != nil {
		return err
	}
	gate, err := s.Store.GetGate(req.GateID)
	if err != nil {
		return err
	}
	session, err := s.Store.GetSession(req.SessionID)
	if err != nil {
		return err
	}
	if _, err := crypto.UnwrapSecret(session.Ciphertext, gate.PublicKey, gate.ID, session.SessionDate, s.Profile); err != nil {
		return err
	}
	return nil
}
