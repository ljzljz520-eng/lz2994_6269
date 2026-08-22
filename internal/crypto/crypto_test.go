package crypto

import "testing"

func TestWrapRoundTrip(t *testing.T) {
	p := DefaultProfile()
	ciphertext, err := WrapSecret("public", "G-34", "2026-08-22", p)
	if err != nil {
		t.Fatal(err)
	}
	secret, err := UnwrapSecret(ciphertext, "public", "G-34", "2026-08-22", p)
	if err != nil {
		t.Fatal(err)
	}
	if len(secret) != p.KeySize {
		t.Fatalf("secret length %d", len(secret))
	}
}

func TestBadCipherAndParameters(t *testing.T) {
	p := DefaultProfile()
	if _, err := UnwrapSecret("bad", "public", "G-34", "2026-08-22", p); err == nil {
		t.Fatal("expected bad cipher")
	}
	if ValidateParameters(p, "other") == nil {
		t.Fatal("expected mismatch")
	}
}
