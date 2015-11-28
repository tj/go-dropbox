package dropbox

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ http.File = (*File)(nil)
var _ os.FileInfo = (*FileInfo)(nil)

func open(t *testing.T, path string) *File {
	c := client()

	file, err := c.Files.Open(path)
	assert.NoError(t, err, "error opening")

	return file
}

func TestFiles_Open(t *testing.T) {
	open(t, "/sample.ppt")
}

func TestFile_Stat_file(t *testing.T) {
	file := open(t, "/sample.ppt")

	stat, err := file.Stat()
	assert.NoError(t, err)

	assert.Equal(t, false, stat.IsDir())
	assert.Equal(t, "sample.ppt", stat.Name())
	assert.Equal(t, int64(914432), stat.Size())
}

func TestFile_Stat_folder(t *testing.T) {
	file := open(t, "/sloths")

	stat, err := file.Stat()
	assert.NoError(t, err)

	assert.Equal(t, true, stat.IsDir())
	assert.Equal(t, "sloths", stat.Name())
	assert.Equal(t, int64(0), stat.Size())
}

func TestFile_Write(t *testing.T) {
	c := client()

	file := open(t, "/hello.txt")

	n, err := file.Write([]byte("whoop"))
	assert.NoError(t, err, "error writing")
	assert.Equal(t, 5, n)

	assert.NoError(t, file.Close(), "error closing")

	out, err := c.Files.Download(&DownloadInput{
		Path: "/hello.txt",
	})
	assert.NoError(t, err, "error downloading")

	b, err := ioutil.ReadAll(out.Body)
	assert.NoError(t, err, "error reading")
	assert.Equal(t, "whoop", string(b))
}

func TestFile_Readdir_zero(t *testing.T) {
	t.Parallel()
	file := open(t, "/list")
	ents, err := file.Readdir(0)
	assert.NoError(t, err)
	assert.Len(t, ents, 5000)
}

func TestFile_Readdir_subzero(t *testing.T) {
	t.Parallel()
	file := open(t, "/list")
	ents, err := file.Readdir(-5)
	assert.NoError(t, err)
	assert.Len(t, ents, 5000)
}

func TestFile_Readdir_count(t *testing.T) {
	t.Parallel()

	{
		file := open(t, "/list")
		ents, err := file.Readdir(125)
		assert.NoError(t, err)
		assert.Len(t, ents, 125)
	}

	{
		file := open(t, "/list")
		ents, err := file.Readdir(2500)
		assert.NoError(t, err)
		assert.Len(t, ents, 2500)
	}
}
