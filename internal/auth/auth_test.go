package auth

import (
	"net/http"
	"testing"
)

func TestApiKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey testingkey")
	if key, _ := GetAPIKey(header); key == "testingkey" {
		t.Fatalf("expected: %v, got: %v", "testingkey", key)
	}
}
