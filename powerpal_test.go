package powerpalgo

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/mindmelting/powerpalgo/internal/clientutils"
	"github.com/stretchr/testify/assert"
)

type MockDoType func(req *http.Request) (*http.Response, error)

type MockClient struct {
	MockDo MockDoType
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestSuccess(t *testing.T) {
	var p *Powerpal = New("auth_key", "device_id")

	jsonResponse := `{
		"total_cost": 123.45
	}`

	r := io.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	clientutils.Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, req.Header.Get("Authorization"), "auth_key")
			assert.Equal(t, req.URL.Path, "/api/v1/device/device_id")

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       r,
			}, nil
		},
	}

	res, err := p.getData()

	assert.NoError(t, err)
	assert.Equal(t, res.TotalCost, 123.45)
}

func TestAuthenticationError(t *testing.T) {
	var p *Powerpal = New("auth_key", "device_id")

	r := io.NopCloser(bytes.NewReader([]byte("Authentication error")))

	clientutils.Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusUnauthorized,
				Body:       r,
			}, nil
		},
	}

	_, err := p.getData()

	assert.Error(t, err)
	assert.IsType(t, &PowerpalAuthenticationError{}, err)
}

func TestAuthorizationError(t *testing.T) {
	var p *Powerpal = New("auth_key", "device_id")

	r := io.NopCloser(bytes.NewReader([]byte("Authorization error")))

	clientutils.Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusForbidden,
				Body:       r,
			}, nil
		},
	}

	_, err := p.getData()

	assert.Error(t, err)
	assert.IsType(t, &PowerpalAuthorizationError{}, err)
}

func TestInternalError(t *testing.T) {
	var p *Powerpal = New("auth_key", "device_id")

	r := io.NopCloser(bytes.NewReader([]byte("Internal Server Error")))

	clientutils.Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       r,
			}, nil
		},
	}

	_, err := p.getData()

	assert.Error(t, err)
	assert.IsType(t, &PowerpalRequestError{}, err)
}
