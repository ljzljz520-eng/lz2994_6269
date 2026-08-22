package store

import (
	"sort"
	"ticketgate/internal/domain"
)

var auditBucket = []byte("audits")

func (s *Store) SaveAudit(value domain.AuditRecord) error {
	if err := domain.ValidateAudit(value); err != nil {
		return err
	}
	data, err := domain.EncodeAudit(value)
	if err != nil {
		return err
	}
	return s.put(auditBucket, value.ID, data)
}

func (s *Store) GetAudit(id string) (domain.AuditRecord, error) {
	data, err := s.get(auditBucket, id)
	if err != nil {
		return domain.AuditRecord{}, err
	}
	return domain.DecodeAudit(data)
}

func (s *Store) ListAudits() ([]domain.AuditRecord, error) {
	data, err := s.list(auditBucket)
	if err != nil {
		return nil, err
	}
	result := make([]domain.AuditRecord, 0, len(data))
	for _, raw := range data {
		value, decodeErr := domain.DecodeAudit(raw)
		if decodeErr != nil {
			return nil, decodeErr
		}
		result = append(result, value)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].CreatedAt < result[j].CreatedAt })
	return result, nil
}

func (s *Store) DeleteAudit(id string) error { return s.remove(auditBucket, id) }
