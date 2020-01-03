package password

import (
	"testing"
)

func TestValid(t *testing.T) {

	testInputs := map[string]bool{
		"hijklmmn": false,
		"abbceffg": false,
		"abbcegjk": false,
		"abcdffaa": true,
		"ghjaabcc": true,
		"hxcbbcdd": true,
	}

	for in, expectedValid := range testInputs {
		p := Password(in)
		isValid := p.isValid()

		if expectedValid != isValid {
			t.Errorf("password validation failed. Expected %v to be %v. Got %v", in, expectedValid, isValid)
		}
	}

}
func TestIncrement(t *testing.T) {

	testInputs := map[string]string{
		"abcdefgh": "abcdffaa",
		"ghijklmn": "ghjaabcc",
		"hxbxwxba": "hxbxxyzz",
	}

	for in, expectedNext := range testInputs {
		p := Password(in)
		p.IncrementUntilValid()

		if expectedNext != string(p) {
			t.Errorf("password incrementing failed. Expected next password after %v to be %v. Got %v", in, expectedNext, string(p))
		}
	}

}
