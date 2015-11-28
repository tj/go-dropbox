package dropbox_test

import (
	"fmt"
	"io"
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

// Example using the Files client directly.
func Example_files() {
	files := dropbox.NewFiles(dropbox.NewConfig("<token>"))

	out, _ := files.Download(&dropbox.DownloadInput{
		Path: "Readme.md",
	})

	io.Copy(os.Stdout, out.Body)
}

// Example using the Users client directly.
func Example_users() {
	users := dropbox.NewUsers(dropbox.NewConfig("<token>"))
	out, _ := users.GetCurrentAccount()
	fmt.Printf("%v\n", out)
}
