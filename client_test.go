package dropbox

import (
	"testing"

	"github.com/segmentio/go-env"
	"github.com/stretchr/testify/assert"
)

func client() *Client {
	token := env.MustGet("DROPBOX_ACCESS_TOKEN")
	return New(NewConfig(token))
}

func TestClient_error_text(t *testing.T) {
	c := client()

	_, err := c.Files.Download(&DownloadInput{
		Path: "asdfasdfasdf",
	})

	assert.Error(t, err)

	e := err.(*Error)
	assert.Contains(t, e.Error(), "Error in call")
	assert.Equal(t, "Bad Request", e.Status)
	assert.Equal(t, 400, e.StatusCode)
}

func TestClient_error_json(t *testing.T) {
	c := client()

	_, err := c.Files.Download(&DownloadInput{"/nothing"})
	assert.Error(t, err)

	e := err.(*Error)
	assert.Contains(t, e.Error(), "path/not_found")
	assert.Equal(t, "Conflict", e.Status)
	assert.Equal(t, 409, e.StatusCode)
}
