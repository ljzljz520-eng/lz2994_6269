package domain

import (
	"fmt"
	"regexp"
	"strings"
)

var datePattern = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`)

func ValidateDate(value string) error {
	if !datePattern.MatchString(value) {
		return fmt.Errorf("date must use YYYY-MM-DD")
	}
	return nil
}

func SessionKey(gateID, date string) string { return NormalizeGateID(gateID) + "@" + date }

func ParseSessionKey(value string) (string, string, error) {
	parts := strings.Split(value, "@")
	if len(parts) != 2 || parts[0] == "" || ValidateDate(parts[1]) != nil {
		return "", "", fmt.Errorf("invalid session key")
	}
	return NormalizeGateID(parts[0]), parts[1], nil
}

func CloneGate(g Gate) Gate { return g }

func CloneSession(s GateSession) GateSession { return s }

func CloneRequest(r ValidationRequest) ValidationRequest { return r }

func CloneAudit(a AuditRecord) AuditRecord { return a }
