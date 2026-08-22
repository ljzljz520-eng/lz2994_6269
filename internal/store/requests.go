package store

import "ticketgate/internal/domain"

var requestBucket = []byte("requests")

func (s *Store) SaveValidationRequest(value domain.ValidationRequest) error {
	if err := domain.ValidateRequest(value); err != nil {
		return err
	}
	data, err := domain.EncodeRequest(value)
	if err != nil {
		return err
	}
	return s.put(requestBucket, value.ID, data)
}

func (s *Store) GetValidationRequest(id string) (domain.ValidationRequest, error) {
	data, err := s.get(requestBucket, id)
	if err != nil {
		return domain.ValidationRequest{}, err
	}
	return domain.DecodeRequest(data)
}

func (s *Store) DeleteValidationRequest(id string) error {
	return s.remove(requestBucket, id)
}
