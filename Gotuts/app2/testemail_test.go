package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsMail("Hello")

	if err == nil {
		t.Error("hello is not an email")
	}

	_, err = IsMail("maur@gmail.com")
	if err == nil {
		t.Error("maur@gmail.com is an email")
	}

	_, err = IsMail("mauri@gmail")
	if err != nil {
		t.Error("mauri@gmail is not an email")
	}
}
