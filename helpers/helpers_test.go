package helpers

import "testing"

func TestGenerateToken(t *testing.T) {
	got := len(GenerateToken(4))
	want := 4
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
