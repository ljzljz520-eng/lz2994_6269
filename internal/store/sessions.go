package store

import "ticketgate/internal/domain"

var sessionBucket = []byte("sessions")

func (s *Store) SaveSession(value domain.GateSession) error {
	if err := domain.ValidateSession(value); err != nil {
		return err
	}
	data, err := domain.EncodeSession(value)
	if err != nil {
		return err
	}
	return s.put(sessionBucket, value.ID, data)
}

func (s *Store) GetSession(id string) (domain.GateSession, error) {
	data, err := s.get(sessionBucket, id)
	if err != nil {
		return domain.GateSession{}, err
	}
	return domain.DecodeSession(data)
}

func (s *Store) ListSessions() ([]domain.GateSession, error) {
	data, err := s.list(sessionBucket)
	if err != nil {
		return nil, err
	}
	result := make([]domain.GateSession, 0, len(data))
	for _, raw := range data {
		value, decodeErr := domain.DecodeSession(raw)
		if decodeErr != nil {
			return nil, decodeErr
		}
		result = append(result, value)
	}
	return result, nil
}

func (s *Store) DeleteSession(id string) error { return s.remove(sessionBucket, id) }
