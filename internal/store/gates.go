package store

import (
	"ticketgate/internal/domain"
)

var gateBucket = []byte("gates")

func (s *Store) SaveGate(g domain.Gate) error {
	if err := domain.ValidateGate(g); err != nil {
		return err
	}
	data, err := domain.EncodeGate(g)
	if err != nil {
		return err
	}
	return s.put(gateBucket, domain.NormalizeGateID(g.ID), data)
}

func (s *Store) GetGate(id string) (domain.Gate, error) {
	data, err := s.get(gateBucket, domain.NormalizeGateID(id))
	if err != nil {
		return domain.Gate{}, err
	}
	return domain.DecodeGate(data)
}

func (s *Store) DeleteGate(id string) error {
	return s.remove(gateBucket, domain.NormalizeGateID(id))
}

func (s *Store) ListGates() ([]domain.Gate, error) {
	data, err := s.list(gateBucket)
	if err != nil {
		return nil, err
	}
	result := make([]domain.Gate, 0, len(data))
	for _, raw := range data {
		value, decodeErr := domain.DecodeGate(raw)
		if decodeErr != nil {
			return nil, decodeErr
		}
		result = append(result, value)
	}
	return result, nil
}
