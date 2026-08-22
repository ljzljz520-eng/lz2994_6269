package service

import (
	"fmt"
	"ticketgate/internal/crypto"
	"ticketgate/internal/domain"
)

type ReplayResult struct {
	RequestID, SessionID, GateID, PayloadDigest string
	Valid                                       bool
	Reason                                      string
}

func (s *Service) ReplayValidation(requestID string) (ReplayResult, error) {
	if err := s.validateReady(); err != nil {
		return ReplayResult{}, err
	}
	req, err := s.Store.GetValidationRequest(requestID)
	if err != nil {
		return ReplayResult{}, err
	}
	gate, err := s.Store.GetGate(req.GateID)
	if err != nil {
		return ReplayResult{}, err
	}
	session, err := s.Store.GetSession(req.SessionID)
	if err != nil {
		return ReplayResult{}, err
	}
	result := ReplayResult{RequestID: req.ID, SessionID: req.SessionID, GateID: req.GateID, PayloadDigest: crypto.PayloadDigest(req.Payload)}
	if !session.State.CanValidate() {
		result.Reason = "session closed"
		return result, nil
	}
	if err := crypto.VerifyWrapped(session.Ciphertext, gate.PublicKey, gate.ID, session.SessionDate, s.Profile); err != nil {
		result.Reason = err.Error()
		return result, nil
	}
	result.Valid = true
	result.Reason = "replayed"
	return result, nil
}

func (s *Service) EnsureSessionReady(id string) error {
	if err := s.validateReady(); err != nil {
		return err
	}
	session, err := s.Store.GetSession(id)
	if err != nil {
		return err
	}
	if session.State == domain.SessionPrepared {
		_, err = s.ActivateSession(id)
		return err
	}
	if session.State == domain.SessionClosed {
		return fmt.Errorf("session closed")
	}
	return nil
}

func (s *Service) SessionParameter(id string) (string, error) {
	session, err := s.Store.GetSession(id)
	if err != nil {
		return "", err
	}
	return session.ParameterID, nil
}
