package domain

import "fmt"

type EventType string

const (
	EventGateRegistered EventType = "gate.registered"
	EventGateDisabled   EventType = "gate.disabled"
	EventSessionReady   EventType = "session.ready"
	EventSessionClosed  EventType = "session.closed"
	EventValidationDone EventType = "session.validation.done"
)

func (e EventType) Valid() bool {
	switch e {
	case EventGateRegistered, EventGateDisabled, EventSessionReady, EventSessionClosed, EventValidationDone:
		return true
	default:
		return false
	}
}

func EventForGate(status GateStatus) (EventType, error) {
	if status == GateActive {
		return EventGateRegistered, nil
	}
	if status == GateDisabled {
		return EventGateDisabled, nil
	}
	return "", fmt.Errorf("unsupported gate status")
}

func EventForSession(state SessionState) (EventType, error) {
	if state == SessionActive {
		return EventSessionReady, nil
	}
	if state == SessionClosed {
		return EventSessionClosed, nil
	}
	return "", fmt.Errorf("unsupported session state")
}

func EventDescription(event EventType) string {
	switch event {
	case EventGateRegistered:
		return "gate registration"
	case EventGateDisabled:
		return "gate disabled"
	case EventSessionReady:
		return "session ready"
	case EventSessionClosed:
		return "session closed"
	case EventValidationDone:
		return "validation complete"
	default:
		return "unknown event"
	}
}

func IsValidationEvent(event string) bool {
	return event == string(ValidationEvent) || event == string(EventValidationDone)
}
