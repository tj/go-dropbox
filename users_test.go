package dropbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_GetCurrentAccount(t *testing.T) {
	c := client()
	_, err := c.Users.GetCurrentAccount()
	assert.NoError(t, err)
}
