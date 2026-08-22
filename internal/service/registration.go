package service

import (
	"fmt"
	"ticketgate/internal/domain"
)

func (s *Service) RegisterGate(g domain.Gate) (domain.Gate, error) {
	if err := s.validateReady(); err != nil {
		return domain.Gate{}, err
	}
	g.ID = domain.NormalizeGateID(g.ID)
	if g.Status == "" {
		g.Status = domain.GateActive
	}
	if g.CreatedAt == "" {
		g.CreatedAt = s.CurrentDate()
	}
	if err := domain.ValidateGate(g); err != nil {
		return domain.Gate{}, err
	}
	if err := s.Store.SaveGate(g); err != nil {
		return domain.Gate{}, fmt.Errorf("save gate: %w", err)
	}
	return s.Store.GetGate(g.ID)
}

func (s *Service) DisableGate(id string) error {
	if err := s.validateReady(); err != nil {
		return err
	}
	g, err := s.Store.GetGate(id)
	if err != nil {
		return err
	}
	g.Status = domain.GateDisabled
	return s.Store.SaveGate(g)
}

func (s *Service) EnableGate(id string) error {
	if err := s.validateReady(); err != nil {
		return err
	}
	g, err := s.Store.GetGate(id)
	if err != nil {
		return err
	}
	g.Status = domain.GateActive
	return s.Store.SaveGate(g)
}
