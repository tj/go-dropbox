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

// ListSharedFolderInput request input.
type ListSharedFolderInput struct {
	Limit   uint64         `json:"limit"`
	Actions []FolderAction `json:"actions"`
}

type FolderAction struct {
	ChangeOptions string
}

// ListSharedFolderOutput lists metadata about shared folders with a cursor to retrieve the next page.
type ListSharedFolderOutput struct {
	Entries []SharedFolderMetadata `json:"entries"`
	Cursor  string                 `json:"cursor"`
}

// SharedFolderMetadata includes basic information about the shared folder.
type SharedFolderMetadata struct {
	AccessType     AccessType   `json:"access_type"`
	IsTeamFolder   bool         `json:"is_team_folder"`
	Policy         FolderPolicy `json:"policy"`
	Name           string       `json:"name"`
	SharedFolderID string       `json:"shared_folder_id"`
	TimeInvited    time.Time    `json:"time_invited"`
	OwnerTeam      struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"owner_team"`
	ParentSharedFolderID string `json:"parent_shared_folder_id"`
	PathLower            string `json:"path_lower"`
	Permissions          string `json:"permissions"`
}

// FolderPolicy enumerates the policies governing this shared folder.
type FolderPolicy struct {
	ACLUpdatePolicy struct {
		Tag ACLUpdatePolicy `json:".tag"`
	} `json:"acl_update_policy"`
	SharedLinkPolicy struct {
		Tag SharedLinkPolicy `json:".tag"`
	} `json:"shared_link_policy"`
	MemberPolicy struct {
		Tag MemberPolicy `json:".tag"`
	} `json:"member_policy"`
	ResolvedMemberPolicy struct {
		Tag MemberPolicy `json:".tag"`
	} `json:"resolved_member_policy"`
}

// AccessType determines the level of access to a shared folder.
type AccessType string

// Access types supported.
const (
	Owner           AccessType = "owner"
	Editor                     = "editor"
	Viewer                     = "viewer"
	ViewerNoComment            = "viewer_no_comment"
)

// ACLUpdatePolicy determines who can add and remove members from this shared folder.
type ACLUpdatePolicy string

// ACLUpdatePolicy types supported.
const (
	ACLUpdatePolicyOwner   ACLUpdatePolicy = "owner"
	ACLUpdatePolicyEditors                 = "editors"
)

// SharedLinkPolicy governs who can view shared links.
type SharedLinkPolicy string

// SharedLinkPolicy types supported.
const (
	SharedLinkPolicyAnyone  SharedLinkPolicy = "anyone"
	SharedLinkPolicyMembers                  = "members"
)

// MemberPolicy determines who can be a member of this shared folder, as set on the folder itself.
type MemberPolicy string

// MemberPolicy types supported.
const (
	MemberPolicyTeam   MemberPolicy = "team"
	MemberPolicyAnyone              = "anyone"
)
