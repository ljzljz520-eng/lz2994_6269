package domain

import (
	"fmt"
	"strings"
)

type AccessDecision struct {
	Allowed bool
	Code    string
	Reason  string
}

func DecideGateAccess(g Gate, session GateSession, request ValidationRequest, policy GatePolicy) AccessDecision {
	if err := policy.ValidateGate(g); err != nil {
		return AccessDecision{Code: "GATE_REJECTED", Reason: err.Error()}
	}
	if !strings.EqualFold(g.ID, session.GateID) {
		return AccessDecision{Code: "GATE_MISMATCH", Reason: "session gate differs"}
	}
	if !session.State.CanValidate() {
		return AccessDecision{Code: "SESSION_CLOSED", Reason: "session cannot validate"}
	}
	if err := policy.ValidateRequest(request); err != nil {
		return AccessDecision{Code: "REQUEST_REJECTED", Reason: err.Error()}
	}
	return AccessDecision{Allowed: true, Code: "ALLOW", Reason: "all policy checks passed"}
}

func (d AccessDecision) Error() error {
	if d.Allowed {
		return nil
	}
	return fmt.Errorf("%s: %s", d.Code, d.Reason)
}

func (d AccessDecision) Terminal() bool { return d.Code != "" }

func (d AccessDecision) Summary() string { return d.Code + "|" + d.Reason }

func CompareDecisions(a, b AccessDecision) bool {
	return a.Allowed == b.Allowed && a.Code == b.Code && a.Reason == b.Reason
}

func AcceptedDecision() AccessDecision {
	return AccessDecision{Allowed: true, Code: "ALLOW", Reason: "all policy checks passed"}
}

func RejectedDecision(code, reason string) AccessDecision {
	return AccessDecision{Allowed: false, Code: code, Reason: reason}
}

func (d AccessDecision) CodeIs(code string) bool { return d.Code == code }

func (d AccessDecision) ReasonContains(fragment string) bool {
	return strings.Contains(d.Reason, fragment)
}

func (d AccessDecision) AllowedCode() string {
	if d.Allowed {
		return "ALLOW"
	}
	return d.Code
}

func (d AccessDecision) IsRejected() bool { return !d.Allowed }

func (d AccessDecision) IsAccepted() bool { return d.Allowed && d.Code == "ALLOW" }

func DecisionCode(allowed bool) string {
	if allowed {
		return "ALLOW"
	}
	return "DENY"
}

func (d AccessDecision) Detail() string { return d.Reason }

func (d AccessDecision) Normalize() AccessDecision {
	d.Code = strings.ToUpper(strings.TrimSpace(d.Code))
	d.Reason = strings.TrimSpace(d.Reason)
	return d
}

func DecisionsAccepted(items []AccessDecision) int {
	count := 0
	for _, item := range items {
		if item.IsAccepted() {
			count++
		}
	}
	return count
}

func DecisionsRejected(items []AccessDecision) int {
	count := 0
	for _, item := range items {
		if item.IsRejected() {
			count++
		}
	}
	return count
}
