package password

import "testing"

func TestGeneratePassword(t *testing.T) {
	cases := map[string]string{
		"test": "3f5a4d8fe75574feca89226e7d53331b9ba235a992b6c7748ecbe0a277764ccb",
		"abc":  "1904813fec3d9ab7f4d3ab60f6d950bff69d52f0a853d28367d81a375cebda0f",
	}
	for in, want := range cases {
		got := GeneratePassword(in)
		if got != want {
			t.Errorf("GeneratePassword(%q) = %q, want %q", in, got, want)
		}
		if len(got) != 64 {
			t.Errorf("GeneratePassword(%q) length = %d, want 64", in, len(got))
		}
	}
}

func TestResetPassword(t *testing.T) {
	got := ResetPassword()
	want := "5b37b751fd3ccee62d95da9b9ed9e08a850e9833d6c2b4110f99daa9219dcff0"
	if got != want {
		t.Errorf("ResetPassword() = %q, want %q", got, want)
	}
}

func TestGenerateLoginToken(t *testing.T) {
	got := GenerateLoginToken(1)
	want := "1ba4686338ee7354395ad13f7d3466c1"
	if got != want {
		t.Errorf("GenerateLoginToken(1) = %q, want %q", got, want)
	}
	if GenerateLoginToken(1) == GenerateLoginToken(2) {
		t.Errorf("GenerateLoginToken should differ for different ids")
	}
}

func TestGenerateMd5Str(t *testing.T) {
	cases := map[string]string{
		"hello":  "5d41402abc4b2a76b9719d911017c592",
		"123456": "e10adc3949ba59abbe56e057f20f883e",
	}
	for in, want := range cases {
		if got := GenerateMd5Str(in); got != want {
			t.Errorf("GenerateMd5Str(%q) = %q, want %q", in, got, want)
		}
	}
}
