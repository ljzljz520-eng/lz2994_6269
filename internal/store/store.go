package store

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"go.etcd.io/bbolt"
)

var ErrNotFound = errors.New("record not found")

var bucketNames = [][]byte{[]byte("gates"), []byte("sessions"), []byte("requests"), []byte("audits")}

type Store struct {
	db   *bbolt.DB
	path string
	mu   sync.RWMutex
}

func Open(path string) (*Store, error) {
	if path == "" {
		return nil, errors.New("store path is required")
	}
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, err
		}
	}
	db, err := bbolt.Open(path, 0o600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
	if err := s.initialize(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}

func (s *Store) initialize() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		for _, name := range bucketNames {
			if _, err := tx.CreateBucketIfNotExists(name); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}

func (s *Store) Path() string { return s.path }

func (s *Store) ensureOpen() error {
	if s.db == nil {
		return errors.New("store is closed")
	}
	return nil
}

func (s *Store) put(bucket []byte, key string, value []byte) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := s.ensureOpen(); err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket(bucket).Put([]byte(key), value)
	})
}

func (s *Store) get(bucket []byte, key string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := s.ensureOpen(); err != nil {
		return nil, err
	}
	var copied []byte
	err := s.db.View(func(tx *bbolt.Tx) error {
		value := tx.Bucket(bucket).Get([]byte(key))
		if value == nil {
			return ErrNotFound
		}
		copied = append([]byte(nil), value...)
		return nil
	})
	return copied, err
}

func (s *Store) remove(bucket []byte, key string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := s.ensureOpen(); err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Delete([]byte(key)) })
}

func (s *Store) list(bucket []byte) (map[string][]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := s.ensureOpen(); err != nil {
		return nil, err
	}
	result := make(map[string][]byte)
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(bucket).ForEach(func(k, v []byte) error {
			result[string(k)] = append([]byte(nil), v...)
			return nil
		})
	})
	return result, err
}
