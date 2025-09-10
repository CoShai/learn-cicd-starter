package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		headers http.Header
	}{
		{headers: http.Header{}},
		{headers: http.Header{}},
		{headers: http.Header{}},
	}

	data := make([]byte, 32)

	for _, test := range tests {
		rand.Read(data)
		test.headers.Set("Authorization", "ApiKey "+hex.EncodeToString(data))
	}

	for _, tc := range tests {
		_, err := GetAPIKey(tc.headers)
		if err != nil {
			t.Error(err)
		}
	}

	fail := http.Header{}
	fail2 := http.Header{}
	fail.Set("Aution", "ApiKey "+hex.EncodeToString(data))
	fail2.Set("Authorization", "ApvvaKey "+hex.EncodeToString(data))

	_, err := GetAPIKey(fail)
	if err == nil {
		t.Errorf("key not equal to Authorization")
	}

	_, err = GetAPIKey(fail2)
	if err == nil {
		t.Errorf("value not containing ApiKey")
	}
}
