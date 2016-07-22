package dropbox

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiles_Upload(t *testing.T) {
	c := client()

	file, err := os.Open("Readme.md")
	assert.NoError(t, err, "error opening file")
	defer file.Close()

	out, err := c.Files.Upload(&UploadInput{
		Mute:   true,
		Mode:   WriteModeOverwrite,
		Path:   "/Readme.md",
		Reader: file,
	})

	assert.NoError(t, err, "error uploading file")
	assert.Equal(t, "/readme.md", out.PathLower)
}

func TestFiles_Download(t *testing.T) {
	c := client()

	out, err := c.Files.Download(&DownloadInput{"/Readme.md"})

	assert.NoError(t, err, "error downloading")
	defer out.Body.Close()

	fi, err := os.Lstat("Readme.md")
	assert.NoError(t, err, "error getting local file info")
	assert.Equal(t, fi.Size(), out.Length, "Readme.md length mismatch")
	assert.NotEmpty(t, out.APIResult)

	remote, err := ioutil.ReadAll(out.Body)
	assert.NoError(t, err, "error reading remote")

	local, err := ioutil.ReadFile("Readme.md")
	assert.NoError(t, err, "error reading local")

	assert.Equal(t, local, remote, "Readme.md mismatch")
}

func TestFiles_GetMetadata(t *testing.T) {
	c := client()

	out, err := c.Files.GetMetadata(&GetMetadataInput{
		Path: "/Readme.md",
	})
	assert.NoError(t, err)
	assert.Equal(t, "file", out.Tag)
}

func TestFiles_ListFolder(t *testing.T) {
	t.Parallel()
	c := client()

	out, err := c.Files.ListFolder(&ListFolderInput{
		Path: "/list",
	})

	assert.NoError(t, err)
	assert.Equal(t, 2000, len(out.Entries))
	assert.True(t, out.HasMore)
}

func TestFiles_ListFolder_root(t *testing.T) {
	t.Parallel()
	c := client()

	_, err := c.Files.ListFolder(&ListFolderInput{
		Path: "/",
	})

	assert.NoError(t, err)
}

func TestFiles_Search(t *testing.T) {
	c := client()

	out, err := c.Files.Search(&SearchInput{
		Path:  "/",
		Query: "hello",
	})

	assert.NoError(t, err)
	assert.Equal(t, 2, len(out.Matches))
}

func TestFiles_Delete(t *testing.T) {
	c := client()

	out, err := c.Files.Delete(&DeleteInput{
		Path: "/Readme.md",
	})

	assert.NoError(t, err)
	assert.Equal(t, "/readme.md", out.PathLower)
}

// A gray, 64 by 64 px PNG
var grayPng = []byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49,
	0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x40, 0x08, 0x02,
	0x00, 0x00, 0x00, 0x25, 0x0b, 0xe6, 0x89, 0x00, 0x00, 0x00, 0x04, 0x67, 0x41,
	0x4d, 0x41, 0x00, 0x00, 0xb1, 0x8f, 0x0b, 0xfc, 0x61, 0x05, 0x00, 0x00, 0x00,
	0x01, 0x73, 0x52, 0x47, 0x42, 0x00, 0xae, 0xce, 0x1c, 0xe9, 0x00, 0x00, 0x00,
	0x09, 0x70, 0x48, 0x59, 0x73, 0x00, 0x00, 0x0e, 0xc3, 0x00, 0x00, 0x0e, 0xc3,
	0x01, 0xc7, 0x6f, 0xa8, 0x64, 0x00, 0x00, 0x00, 0x18, 0x74, 0x45, 0x58, 0x74,
	0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x00, 0x70, 0x61, 0x69, 0x6e,
	0x74, 0x2e, 0x6e, 0x65, 0x74, 0x20, 0x34, 0x2e, 0x30, 0x2e, 0x36, 0xfc, 0x8c,
	0x63, 0xdf, 0x00, 0x00, 0x00, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x68, 0xde, 0xed,
	0xcf, 0x31, 0x0d, 0x00, 0x00, 0x08, 0x03, 0x30, 0xe6, 0x7e, 0xb2, 0x91, 0xc0,
	0x4d, 0xd2, 0x3a, 0x68, 0xda, 0xce, 0x67, 0x11, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0xb8, 0x2c, 0x4a, 0x27, 0x66, 0x41, 0xb9, 0xd3, 0xef, 0xa3, 0x00,
	0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
}

func TestFiles_GetThumbnail(t *testing.T) {
	c := client()
	// REVIEW(bg): This feels a bit sloppy...
	{
		buf := bytes.NewBuffer(grayPng)
		_, err := c.Files.Upload(&UploadInput{
			Mute:   true,
			Mode:   WriteModeOverwrite,
			Path:   "/gray.png",
			Reader: buf,
		})
		assert.NoError(t, err, "error uploading file")
	}
	out, err := c.Files.GetThumbnail(&GetThumbnailInput{"/gray.png", GetThumbnailFormatJPEG, GetThumbnailSizeW32H32})
	assert.NoError(t, err)
	if err != nil {
		return
	}
	defer out.Body.Close()

	assert.NotEmpty(t, out.Length, "length should not be 0")

	buf := make([]byte, 11)
	_, err = out.Body.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, []byte{
		0xff, 0xd8, // JPEG SOI marker
		0xff, 0xe0, 0x00, 0x10, 0x4a, 0x46, 0x49, 0x46, 0x00, // JFIF tag
	}, buf, "should have jpeg header")
	assert.NotEmpty(t, out.APIResult)
}

func TestFiles_GetPreview(t *testing.T) {
	c := client()

	out, err := c.Files.GetPreview(&GetPreviewInput{"/sample.ppt"})
	defer out.Body.Close()

	assert.NoError(t, err)

	assert.NotEmpty(t, out.Length, "length should not be 0")

	buf := make([]byte, 4)
	_, err = out.Body.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x25, 0x50, 0x44, 0x46}, buf, "should have pdf magic number")
	assert.NotEmpty(t, out.APIResult)
}

func TestFiles_ListRevisions(t *testing.T) {
	c := client()

	out, err := c.Files.ListRevisions(&ListRevisionsInput{Path: "/sample.ppt"})

	assert.NoError(t, err)
	assert.NotEmpty(t, out.Entries)
	assert.False(t, out.IsDeleted)
}

func TestFiles_UnmarshalDropboxApiResultNotSet(t *testing.T) {
	hdr := http.Header{
		"Date":                   []string{"Wed, 13 Jan 2016 03:14:59 GMT"},
		"Content-Length":         []string{"808"},
		"Server":                 []string{"nginx"},
		"X-Server-Response-Time": []string{"740"},
		"X-Dropbox-Request-Id":   []string{"2053baf502dc4d4d64f2201a12f0ce26"},
		"X-Robots-Tag":           []string{"noindex", "nofollow", "noimageindex"},
	}

	v := unmarshalDropboxApiResult(hdr)

	assert.Empty(t, v)
}

func TestFiles_UnmarshalDropboxApiResultSet(t *testing.T) {
	hdr := http.Header{
		"Date":                   []string{"Wed, 13 Jan 2016 03:14:59 GMT"},
		"Content-Length":         []string{"808"},
		"Server":                 []string{"nginx"},
		"Dropbox-Api-Result":     []string{`{"name": "Readme.md", "path_lower": "/readme.md", "id": "id:dxYX4OTicUAAAAAAAAATzA", "client_modified": "2001-01-01T00:00:00Z", "server_modified": "2016-01-13T21:42:37Z", "rev": "14a4406950f6", "size": 808}`},
		"X-Server-Response-Time": []string{"740"},
		"X-Dropbox-Request-Id":   []string{"2053baf502dc4d4d64f2201a12f0ce26"},
		"X-Robots-Tag":           []string{"noindex", "nofollow", "noimageindex"},
	}

	v := unmarshalDropboxApiResult(hdr)

	assert.NotEmpty(t, v)
	assert.Equal(t, "id:dxYX4OTicUAAAAAAAAATzA", v.ID)
}
