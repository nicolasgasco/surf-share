package integration_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"surf-share/app/internal/adapters"
	"surf-share/app/internal/modules/auth"
	"surf-share/app/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const testEmail = "test@example.com"
const testPassword = "password123"
const testUsername = "testuser"

type AuthTestSuite struct {
	suite.Suite
	dbAdapter *adapters.DatabaseAdapter
	server    *httptest.Server
}

func (s *AuthTestSuite) SetupTest() {
	connStr, err := testutil.GetDbConnectionString()
	require.NoError(s.T(), err)

	ctx := context.Background()
	s.dbAdapter = &adapters.DatabaseAdapter{}
	err = s.dbAdapter.Connect(ctx, connStr)
	require.NoError(s.T(), err)

	mux := http.NewServeMux()

	userRepository := auth.NewRepository(s.dbAdapter)
	passwordHasher := auth.NewBcryptHasher()
	tokenGenerator := auth.NewJWTGenerator([]byte("test-secret"))
	authService := auth.NewAuthService(userRepository, passwordHasher, tokenGenerator)
	httpHandler := auth.NewHTTPHandler(authService)
	mux.HandleFunc("POST /auth/register", httpHandler.HandleRegister)
	mux.HandleFunc("POST /auth/login", httpHandler.HandleLogin)

	s.server = httptest.NewServer(mux)
}

func (s *AuthTestSuite) TearDownTest() {
	if s.dbAdapter != nil {
		s.dbAdapter.Exec(context.Background(), "DELETE FROM app.users WHERE email = $1", testEmail)
		s.dbAdapter.Close()
	}

	if s.server != nil {
		s.server.Close()
	}
}

func (s *AuthTestSuite) TestRegister() {
	data := strings.NewReader("username=" + testUsername + "&email=" + testEmail + "&password=" + testPassword)
	resp, err := http.Post(s.server.URL+"/auth/register", "application/x-www-form-urlencoded", data)
	require.NoError(s.T(), err)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		s.T().Fatalf("Expected 201, got %d: %s", resp.StatusCode, string(body))
	}

	var response struct {
		User  *auth.User `json:"user"`
		Token string     `json:"token"`
	}

	err = json.Unmarshal(body, &response)
	require.NoError(s.T(), err)

	assert.NotEmpty(s.T(), response.User.ID)
	assert.Equal(s.T(), testUsername, response.User.Username)
	assert.Equal(s.T(), testEmail, response.User.Email)
	assert.NotEmpty(s.T(), response.Token)
}

func (s *AuthTestSuite) TestLogin() {
	registerData := strings.NewReader("username=" + testUsername + "&email=" + testEmail + "&password=" + testPassword)
	resp, err := http.Post(s.server.URL+"/auth/register", "application/x-www-form-urlencoded", registerData)
	require.NoError(s.T(), err)
	resp.Body.Close()

	loginData := strings.NewReader("email=" + testEmail + "&password=" + testPassword)
	resp, err = http.Post(s.server.URL+"/auth/login", "application/x-www-form-urlencoded", loginData)
	require.NoError(s.T(), err)
	defer resp.Body.Close()

	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(s.T(), "application/json", resp.Header.Get("Content-Type"))

	var response struct {
		User  *auth.User `json:"user"`
		Token string     `json:"token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(s.T(), err)

	assert.NotEmpty(s.T(), response.User.ID)
	assert.Equal(s.T(), testUsername, response.User.Username)
	assert.Equal(s.T(), testEmail, response.User.Email)
	assert.NotEmpty(s.T(), response.Token)
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
