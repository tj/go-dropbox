package dropbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestSharing_CreateSharedLink(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.dropboxapi.com/2/sharing/create_shared_link_with_settings",
		httpmock.NewStringResponder(200, `{"path":"/hello.txt"}`))

	c := client()
	out, err := c.Sharing.CreateSharedLink(&CreateSharedLinkInput{
		Path: "/hello.txt",
	})

	assert.NoError(t, err, "error sharing file")
	assert.Equal(t, "/hello.txt", out.Path)
}

func TestSharing_ListSharedFolder(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.dropboxapi.com/2/sharing/list_folders",
		httpmock.NewStringResponder(200, `{"entries":[{"name":"foo", "shared_folder_id":"321321"}]}`))

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
