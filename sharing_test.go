package dropbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharing_CreateSharedLink(t *testing.T) {
	c := client()
	out, err := c.Sharing.CreateSharedLink(&CreateSharedLinkInput{
		Path:     "/hello.txt",
		ShortURL: true,
	})

	assert.NoError(t, err, "error sharing file")
	assert.Equal(t, "/hello.txt", out.Path)
}
