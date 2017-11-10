package dropbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharing_CreateSharedLink(t *testing.T) {
	c := client()
	//list shares
	links, _ := c.Sharing.ListSharedLinks(&ListShareLinksInput {
			Path: "/hello.txt",
		})

	var sharedLink string
	if len(links.Links) == 0 {
		out, err := c.Sharing.CreateSharedLink(&CreateSharedLinkInput{
			Path: "/hello.txt",
		})
		
		assert.NoError(t, err, "error sharing file")
		assert.Contains(t, out.URL, "/hello.txt")
	} else {
		sharedLink = links.Links[0].URL
		
		err := c.Sharing.RevokeSharedLink(&RevokeSharedLinkInput{
			Url: sharedLink,
		})

		out, err := c.Sharing.CreateSharedLink(&CreateSharedLinkInput{
			Path: "/hello.txt",
		})
		
		assert.NoError(t, err, "error sharing file")
		assert.Contains(t, out.URL, "/hello.txt")
	}
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

func TestSharing_ListSharedLinks(t *testing.T) {
	c := client()
	out, err := c.Sharing.ListSharedLinks(&ListShareLinksInput{
		Path: "/hello.txt",
	})

	assert.NoError(t, err, "listing shared folders")
	assert.NotEmpty(t, out.Links, "output should be non-empty")
}

func TestSharing_RevokeSharedLink(t *testing.T) {
	c := client()

	links, err := c.Sharing.ListSharedLinks(&ListShareLinksInput {
		Path: "/hello.txt",
	})
	
	var sharedLink string
	if len(links.Links) == 0 {
		sl := CreateSharedLinkInput{ Path: "/hello.txt" }

		link, err := c.Sharing.CreateSharedLink(&sl)

		assert.NoError(t, err, "revoke shared link")
		sharedLink = link.URL
	} else {
		sharedLink = links.Links[0].URL
	}
	
	err = c.Sharing.RevokeSharedLink(&RevokeSharedLinkInput{
		Url: sharedLink,
	})

	assert.NoError(t, err, "revoke shared link")
}