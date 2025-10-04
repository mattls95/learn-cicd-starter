package auth

import (
	"net/http"
	"testing"
)

func TestAuthNoHeader(t *testing.T) {
	headers := http.Header{
		"Authorization": {""},
	}
	want := ErrNoAuthHeaderIncluded
	_, got := GetAPIKey(headers)

	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
