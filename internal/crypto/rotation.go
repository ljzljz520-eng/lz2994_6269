package crypto

import (
	"crypto/sha256"
	"fmt"
)

type KeyMaterial struct {
	GateID, PublicKey, Fingerprint string
	Version                        int
}

func BuildKeyMaterial(gateID, publicKey string, version int) (KeyMaterial, error) {
	if gateID == "" || publicKey == "" || version < 1 {
		return KeyMaterial{}, fmt.Errorf("invalid key material")
	}
	h := sha256.Sum256([]byte(gateID + "|" + publicKey))
	return KeyMaterial{GateID: gateID, PublicKey: publicKey, Fingerprint: fmt.Sprintf("%x", h[:8]), Version: version}, nil
}

func RotateKey(old KeyMaterial, nextPublicKey string) (KeyMaterial, error) {
	if nextPublicKey == "" {
		return KeyMaterial{}, fmt.Errorf("next key is empty")
	}
	return BuildKeyMaterial(old.GateID, nextPublicKey, old.Version+1)
}

func IsNewer(current, candidate KeyMaterial) bool { return candidate.Version > current.Version }

func CompareMaterial(left, right KeyMaterial) int {
	if left.Version < right.Version {
		return -1
	}
	if left.Version > right.Version {
		return 1
	}
	if left.Fingerprint < right.Fingerprint {
		return -1
	}
	if left.Fingerprint > right.Fingerprint {
		return 1
	}
	return 0
}
