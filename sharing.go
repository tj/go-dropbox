package dropbox

import (
	"encoding/json"
	"time"
)

// Sharing client.
type Sharing struct {
	*Client
}

// NewSharing client.
func NewSharing(config *Config) *Sharing {
	return &Sharing{
		Client: &Client{
			Config: config,
		},
	}
}

// CreateSharedLinkInput request input.
type CreateSharedLinkInput struct {
	Path     string `json:"path"`
	ShortURL bool   `json:"short_url"`
}

// CreateSharedLinkOutput request output.
type CreateSharedLinkOutput struct {
	URL             string `json:"url"`
	Path            string `json:"path"`
	VisibilityModel struct {
		Tag VisibilityType `json:".tag"`
	} `json:"visibility"`
	Expires time.Time `json:"expires,omitempty"`
}

// VisibilityType determines who can access the link.
type VisibilityType string

// Visibility types supported.
const (
	Public           VisibilityType = "public"
	TeamOnly                        = "team_only"
	Password                        = "password"
	TeamAndPassword                 = "team_and_password"
	SharedFolderOnly                = "shared_folder_only"
)

// CreateSharedLink returns a shared link.
func (c *Sharing) CreateSharedLink(in *CreateSharedLinkInput) (out *CreateSharedLinkOutput, err error) {
	body, err := c.call("/sharing/create_shared_link", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
