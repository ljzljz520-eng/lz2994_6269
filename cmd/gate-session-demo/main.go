package main

import (
	"fmt"
	"os"
	"path/filepath"
	"ticketgate/internal/fixture"
	"ticketgate/internal/logsummary"
	"ticketgate/internal/service"
	"ticketgate/internal/store"
)

func main() {
	path := filepath.Join(os.TempDir(), "ticket-gate-session-demo.db")
	_ = os.Remove(path)
	st, err := store.Open(path)
	if err != nil {
		panic(err)
	}
	defer st.Close()
	svc := service.New(st, fixture.DefaultClock())
	if _, err := svc.RegisterGate(fixture.GateG34()); err != nil {
		panic(err)
	}
	session, negotiation, err := svc.NegotiateSession(fixture.GateID)
	if err != nil {
		panic(err)
	}
	request := fixture.RequestG34(session.ID)
	validation, err := svc.ValidateGateSession(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(logsummary.Summary(negotiation))
	fmt.Println(logsummary.Summary(validation))
}
