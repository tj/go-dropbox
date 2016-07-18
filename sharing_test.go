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

func TestSharing_ListSharedFolder(t *testing.T) {
	c := client()
	out, err := c.Sharing.ListSharedFolders(&ListSharedFolderInput{
		Limit: 1,
	})

	shared := out.Entries

	assert.NoError(t, err, "listing shared folders")
	assert.NotEmpty(t, out.Entries, "output should be non-empty")

	for out.Cursor != "" {
		out, err = c.Sharing.ListSharedFoldersContinue(&ListSharedFolderContinueInput{
			Cursor: out.Cursor,
		})

		shared = append(shared, out.Entries...)

		assert.NoError(t, err, "listing shared folders")
		assert.NotEmpty(t, out.Entries, "output should be non-empty")
	}

	for _, sharedFolder := range shared {
		out, err := c.Sharing.ListSharedFolderMembers(&ListSharedFolderMembersInput{
			SharedFolderID: sharedFolder.SharedFolderID,
			Limit:          1,
		})

		assert.NoError(t, err, "listing shared folder members")
		assert.Equal(t, 1, len(out.Users)+len(out.Groups)+len(out.Invitees), "there should be 1 item present")

		for out.Cursor != "" {
			out, err = c.Sharing.ListSharedFolderMembersContinue(&ListSharedFolderMembersInputContinue{
				Cursor: out.Cursor,
			})

			assert.NoError(t, err, "listing shared folder members")
			assert.Equal(t, 1, len(out.Users)+len(out.Groups)+len(out.Invitees), "there should be 1 item present")
		}
	}
}
