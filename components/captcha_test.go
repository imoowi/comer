package components

import "testing"

func TestGenCaptchaDrivers(t *testing.T) {
	for _, dt := range []string{`digit`, `string`, `math`, `chinese`, `audio`, ``} {
		id, b64s, err := GenCaptcha(dt)
		if err != nil {
			t.Fatalf("GenCaptcha(%q) error: %v", dt, err)
		}
		if id == `` || b64s == `` {
			t.Fatalf("GenCaptcha(%q) returned empty id/b64s", dt)
		}
	}
}

func TestVerifyCaptchaNotFound(t *testing.T) {
	if VerifyCaptcha(`nonexistent-id-12345`, `1234`) {
		t.Fatalf("VerifyCaptcha for nonexistent id should be false")
	}
}
