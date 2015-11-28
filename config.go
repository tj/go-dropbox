package dropbox

import (
	"net/http"
)

// Config for the Dropbox clients.
type Config struct {
	HTTPClient  *http.Client
	AccessToken string
}

// NewConfig with the given access token.
func NewConfig(accessToken string) *Config {
	return &Config{
		HTTPClient:  http.DefaultClient,
		AccessToken: accessToken,
	}
}
