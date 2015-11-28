package dropbox_test

import (
	"io"
	"os"

	"github.com/tj/go-dropbox"
)

// Example using the Files client directly.
func Example_files() {
	files := dropbox.NewFiles(dropbox.NewConfig("<token>"))

	out, _ := files.Download(&dropbox.DownloadInput{
		Path: "Readme.md",
	})

	io.Copy(os.Stdout, out.Body)
}
