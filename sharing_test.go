package dropbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharing_CreateSharedLink(t *testing.T) {
	c := client()
	out, err := c.Sharing.CreateSharedLink(&CreateSharedLinkInput{
		Path:     "/hello.txt"
	})

	assert.NoError(t, err, "error sharing file")
	assert.Equal(t, "/hello.txt", out.Path)
}

func TestSharing_ListSharedFolder(t *testing.T) {
	c := client()
	out, err := c.Sharing.ListSharedFolders(&ListSharedFolderInput{
		Limit: 1,
	})

	assert.NoError(t, err, "listing shared folders")
	assert.NotEmpty(t, out.Entries, "output should be non-empty")

	for out.Cursor != "" {
		out, err = c.Sharing.ListSharedFoldersContinue(&ListSharedFolderContinueInput{
			Cursor: out.Cursor,
		})

		assert.NoError(t, err, "listing shared folders")
		assert.NotEmpty(t, out.Entries, "output should be non-empty")
	}
}
