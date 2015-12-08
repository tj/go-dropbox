package dropbox

import (
	"encoding/json"
	"time"
)

// Sharing client.
type Sharing struct {
	*Client
}

//NewSharing client.
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
	ShortUrl bool   `json:"short_url"`
}

// CreateSharedLinkOutput request output.
type CreateSharedLinkOutput struct {
	Url             string     `json:"url"`
	Path            string     `json:"path"`
	VisibilityModel Visibility `json:"visibility"`
	Expires         time.Time  `json:"expires,omitempty"`
}

type VisibilityType string

// VisibilityType  Who can access the link.
const (
	Public           VisibilityType = "public"
	TeamOnly                        = "team_only"
	Password                        = "password"
	TeamAndPassword                 = "team_and_password"
	SharedFolderOnly                = "shared_folder_only"
)

type Visibility struct {
	Tag VisibilityType `json:".tag"`
}

// Create a shared link.
func (c *Sharing) CreateSharedLink(in *CreateSharedLinkInput) (out *CreateSharedLinkOutput, err error) {

	body, err := c.call("/sharing/create_shared_link", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
