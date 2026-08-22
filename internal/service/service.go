package service

import (
	"errors"
	"ticketgate/internal/crypto"
	"ticketgate/internal/store"
)

type Clock interface{ Today() string }

type StaticClock struct{ Date string }

func (c StaticClock) Today() string { return c.Date }

type Service struct {
	Store   *store.Store
	Clock   Clock
	Profile crypto.Profile
}

func New(st *store.Store, clock Clock) *Service {
	if clock == nil {
		clock = StaticClock{Date: "2026-08-22"}
	}
	return &Service{Store: st, Clock: clock, Profile: crypto.DefaultProfile()}
}

func (s *Service) validateReady() error {
	if s == nil || s.Store == nil {
		return errors.New("service store is not configured")
	}
	if s.Clock == nil {
		return errors.New("service clock is not configured")
	}
	return s.Profile.Validate()
}

func (s *Service) CurrentDate() string {
	if s == nil || s.Clock == nil {
		return ""
	}
	return s.Clock.Today()
}
