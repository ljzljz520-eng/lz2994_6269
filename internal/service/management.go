package service

import (
	"fmt"
	"ticketgate/internal/domain"
)

func (s *Service) RegisterMany(gates []domain.Gate) ([]domain.Gate, error) {
	if err := s.validateReady(); err != nil {
		return nil, err
	}
	result := make([]domain.Gate, 0, len(gates))
	for _, gate := range gates {
		saved, err := s.RegisterGate(gate)
		if err != nil {
			return nil, err
		}
		result = append(result, saved)
	}
	return result, nil
}

func (s *Service) RevokeSession(id string) error {
	if err := s.validateReady(); err != nil {
		return err
	}
	session, err := s.Store.GetSession(id)
	if err != nil {
		return err
	}
	if session.State == domain.SessionClosed {
		return fmt.Errorf("session already closed")
	}
	if err := domain.TransitionSession(&session, domain.SessionClosed); err != nil {
		return err
	}
	return s.Store.SaveSession(session)
}

func (s *Service) SessionHealth(id string) (string, error) {
	if err := s.validateReady(); err != nil {
		return "", err
	}
	session, err := s.Store.GetSession(id)
	if err != nil {
		return "", err
	}
	if session.State == domain.SessionActive {
		return "healthy", nil
	}
	if session.State == domain.SessionPrepared {
		return "pending", nil
	}
	return "closed", nil
}

func (s *Service) ValidatePayload(req domain.ValidationRequest) error {
	policy := domain.DefaultGatePolicy()
	return policy.ValidateRequest(req)
}
