package dropbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharing_CreateSharedLink(t *testing.T) {
	c := client()
	out, err := c.Sharing.CreateSharedLink(&CreateSharedLinkInput{
		Path:     "/Readme.md",
		ShortURL: true,
	})

	assert.NoError(t, err, "error sharing file")
	assert.Equal(t, "/Readme.md", out.Path)
}
