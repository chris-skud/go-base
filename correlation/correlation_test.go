package correlation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chris-skud/go-base/correlation"
	"github.com/stretchr/testify/assert"
)

func TestCorrelation(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	tVal := "correlation1"

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(tVal, correlation.ID(r.Context()))
		w.Write([]byte("OK"))
	})
	ts := httptest.NewServer(correlation.InitID(ctx, handler))
	defer ts.Close()

	// Make a request with headers to ensure correlation id added from header.
	req, _ := http.NewRequest("GET", ts.URL, nil)
	req.Header.Set("CORRELATION-ID", tVal)
	_, err := http.DefaultClient.Do(req)
	assert.NoError(err)
}

func TestCorrelationNoHeader(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEmpty(correlation.ID(r.Context()))
		w.Write([]byte("OK"))
	})
	ts := httptest.NewServer(correlation.InitID(ctx, handler))
	defer ts.Close()

	// Make a request with no headers to ensure default correlation id set.
	req, _ := http.NewRequest("GET", ts.URL, nil)
	_, err := http.DefaultClient.Do(req)
	assert.NoError(err)
}
