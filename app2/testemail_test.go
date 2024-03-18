package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T){
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("test_testy@hotmail.com")
	if err != nil {
		t.Error("test_testy@homtail.com is an valid email")
	}

	_, err = IsEmail("test_testyhotmail")
	if err != nil {
		t.Error("test_testyhotmail is not an valid email")
	}
}