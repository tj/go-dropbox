package dropbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func client() *Client {
	return New(NewConfig("KfNzQkz4taMAAAAAAAAosMzpPAFNp6ZJHpr3jpXZJrjkQUWHPhZpaciQW40Oxzfb"))
}

func TestUsers_GetCurrentAccount(t *testing.T) {
	c := client()
	_, err := c.Users.GetCurrentAccount()
	assert.NoError(t, err)
}
