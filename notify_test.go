package main

import (
	"bytes"
	"encoding/json"
	"github.com/davidbanham/notify/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHealth(t *testing.T) {
	t.Parallel()

	req := &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/health"},
	}

	rr := httptest.NewRecorder()
	healthHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestSMSRestAPI(t *testing.T) {
	t.Parallel()

	data := map[string]interface{}{
		"body": "Oh hi!",
		"to":   "+61438000000",
	}
	body, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "/v1/sms", bytes.NewReader(body))

	rr := httptest.NewRecorder()
	smsHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "from", "from field not populated")
	assert.Contains(t, rr.Body.String(), config.SmsFrom, "from field not populated")
}

func TestEmailRestAPI(t *testing.T) {
	t.Parallel()

	data := map[string]interface{}{
		"subject": "Hello",
		"body": map[string]string{
			"text": "Yes, this is dog.",
		},
		"to": map[string]string{
			"address": "cat@example.com",
			"name":    "cat",
		},
		"from": map[string]string{
			"address": "dog@example.com",
			"name":    "dog",
		},
	}
	body, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "/v1/email", bytes.NewReader(body))

	rr := httptest.NewRecorder()
	emailHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "from", "from field not populated")
	assert.Contains(t, rr.Body.String(), "html", "html field not populated")
}
