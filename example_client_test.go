package dropbox_test

import (
	"os"

	"github.com/tj/go-dropbox"
)

// Example using the Client, which provides both User and File clients.
func Example() {
	d := dropbox.New(dropbox.NewConfig("<token>"))

	file, _ := os.Open("Readme.md")

	d.Files.Upload(&dropbox.UploadInput{
		Path:   "Readme.md",
		Reader: file,
		Mute:   true,
	})
}
