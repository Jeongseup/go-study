package myapp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	// mockup server create
	ts := httptest.NewServer(NewHandler())
	defer ts.Close() // server close

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.Code)
}
