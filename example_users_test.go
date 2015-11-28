package dropbox_test

import (
	"fmt"

	"github.com/tj/go-dropbox"
)

// Example using the Users client directly.
func Example_users() {
	users := dropbox.NewUsers(dropbox.NewConfig("<token>"))
	out, _ := users.GetCurrentAccount()
	fmt.Printf("%v\n", out)
}
