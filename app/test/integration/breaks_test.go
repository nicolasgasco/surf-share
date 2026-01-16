package integration_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"surf-share/app/internal/testutil"
	"testing"

	"surf-share/app/internal/adapters"
	"surf-share/app/internal/models"
	"surf-share/app/internal/modules/breaks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BreaksTestSuite struct {
	suite.Suite
	dbAdapter *adapters.DatabaseAdapter
	server    *httptest.Server
}

func (s *BreaksTestSuite) SetupTest() {
	connStr, err := testutil.GetDbConnectionString()
	require.NoError(s.T(), err)

	ctx := context.Background()
	s.dbAdapter = &adapters.DatabaseAdapter{}
	err = s.dbAdapter.Connect(ctx, connStr)
	require.NoError(s.T(), err)

	mux := http.NewServeMux()
	breaksHandler := breaks.NewBreaksHandler(s.dbAdapter)
	mux.HandleFunc("GET /breaks", breaksHandler.HandleBreaks)
	mux.HandleFunc("GET /breaks/{slug}", breaksHandler.HandleBreakBySlug)

	s.server = httptest.NewServer(mux)
}

func (s *BreaksTestSuite) TearDownTest() {
	if s.server != nil {
		s.server.Close()
	}

	if s.dbAdapter != nil {
		s.dbAdapter.Close()
	}
}

func (s *BreaksTestSuite) TestGetBreaks() {
	resp, err := http.Get(s.server.URL + "/breaks")
	require.NoError(s.T(), err)
	defer resp.Body.Close()

	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(s.T(), "application/json", resp.Header.Get("Content-Type"))

	var response struct {
		Count  int                     `json:"count"`
		Breaks []breaks.BreaksResponse `json:"breaks"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(s.T(), err)

	require.NotEmpty(s.T(), response.Breaks)
	assert.Equal(s.T(), len(response.Breaks), response.Count)

	require.GreaterOrEqual(s.T(), len(response.Breaks), 2)
	assert.LessOrEqual(s.T(), response.Breaks[0].Name, response.Breaks[1].Name)
}

func (s *BreaksTestSuite) TestGetBreakBySlug() {
	// Get the slug of the first break
	resp, err := http.Get(s.server.URL + "/breaks")
	require.NoError(s.T(), err)
	defer resp.Body.Close()

	var response struct {
		Breaks []breaks.BreaksResponse `json:"breaks"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), response.Breaks)

	firstBreak := response.Breaks[0]

	resp, err = http.Get(s.server.URL + "/breaks/" + firstBreak.Slug)
	require.NoError(s.T(), err)
	defer resp.Body.Close()

	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(s.T(), "application/json", resp.Header.Get("Content-Type"))

	var breakResponse breaks.BreakResponse
	err = json.NewDecoder(resp.Body).Decode(&breakResponse)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), firstBreak.Slug, breakResponse.Slug)
	assert.Equal(s.T(), firstBreak.Name, breakResponse.Name)
}

func (s *BreaksTestSuite) TestGetBreakByInvalidSlug() {
	resp, err := http.Get(s.server.URL + "/breaks/invalid-slug")
	require.NoError(s.T(), err)
	defer resp.Body.Close()

	assert.Equal(s.T(), http.StatusNotFound, resp.StatusCode)
	assert.Equal(s.T(), "application/json", resp.Header.Get("Content-Type"))

	var errorResponse models.ErrorResponse

	err = json.NewDecoder(resp.Body).Decode(&errorResponse)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), "Break with slug 'invalid-slug' not found", errorResponse.Message)
}

func TestBreaksTestSuite(t *testing.T) {
	suite.Run(t, new(BreaksTestSuite))
}
