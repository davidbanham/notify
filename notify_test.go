package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestOrganisationCreateOrUpdateHandler(t *testing.T) {
	t.Parallel()

	req := &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/health"},
	}

	rr := httptest.NewRecorder()
	healthHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
