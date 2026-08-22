package fixture

import (
	"encoding/json"
	"ticketgate/internal/domain"
)

type Manifest struct {
	Name     string
	Gates    []domain.Gate
	Requests []domain.ValidationRequest
	Date     string
}

func DefaultManifest() Manifest {
	return Manifest{Name: "ticket-gate-deterministic", Gates: Gates(3), Date: SessionDate}
}

func (m Manifest) GateIDs() []string {
	result := make([]string, 0, len(m.Gates))
	for _, gate := range m.Gates {
		result = append(result, gate.ID)
	}
	return result
}

func (m Manifest) Encode() ([]byte, error) { return json.Marshal(m) }

func DecodeManifest(data []byte) (Manifest, error) {
	var value Manifest
	err := json.Unmarshal(data, &value)
	return value, err
}

func BuildManifest(requests []domain.ValidationRequest) Manifest {
	return Manifest{Name: "request-batch", Gates: []domain.Gate{GateG34()}, Requests: requests, Date: SessionDate}
}
