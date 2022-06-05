package helpers

import "testing"

func TestValidateURL(t *testing.T) {

	got := ValidateURL("http://ya.ru/")
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got2 := ValidateURL("123123")
	want2 := false
	if got2 != want2 {
		t.Errorf("got %t, wanted %t", got2, want2)
	}

}
