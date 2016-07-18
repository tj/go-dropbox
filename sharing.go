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

// ListSharedFolderMembersInput request input.
type ListSharedFolderMembersInput struct {
	SharedFolderID string         `json:"shared_folder_id"`
	Actions        []MemberAction `json:"actions"`
	Limit          uint64         `json:"limit"`
}

// ListSharedFolderMembersOutput enumerates shared folder user and group membership.
type ListSharedFolderMembersOutput struct {
	Users    []UserMembershipInfo    `json:"users"`
	Groups   []GroupMembershipInfo   `json:"groups"`
	Invitees []InviteeMembershipInfo `json:"invitees"`
	Cursor   string                  `json:"cursor"`
}

// UserMembershipInfo is information about a user member of the shared folder.
type UserMembershipInfo struct {
	AccessType struct {
		Tag AccessType `json:".tag"`
	} `json:"access_type"`
	User        UserInfo           `json:"user"`
	Permissions []MemberPermission `json:"permissions,omitempty"`
	Initials    string             `json:"initials,omitempty"`
	IsInherited bool               `json:"is_inherited"`
}

// UserInfo is the account information for the membership user.
type UserInfo struct {
	AccountID    string `json:"account_id"`
	SameTeam     bool   `json:"same_team"`
	TeamMemberID string `json:"team_member_id"`
}

// MemberPermission indicates whether the user is allowed to take the action on the associated member.
type MemberPermission struct {
	Action MemberAction           `json:"action"`
	Allow  bool                   `json:"allow"`
	Reason PermissionDeniedReason `json:"reason,omitempty"`
}

// MemberAction indicates the action that the user may wish to take on the member.
type MemberAction string

// MemberAction possible values.
const (
	MemberActionMakeEditor MemberAction = "make_editor"
	MemberActionMakeOwner               = "make_owner"
	MemberActionMakeViewer              = "make_viewer"
	MemberActionRemove                  = "remove"
)

// PermissionDeniedReason is the reason the user is denied a permission.
type PermissionDeniedReason string

// PermissionDeniedReason possible values
const (
	UserNotOnSameTeamAsOwner PermissionDeniedReason = "user_not_same_team_as_owner"
	UserNotAllowedByOwner                           = "user_not_allowed_by_owner"
	TargetIsIndirectmember                          = "target_is_indirect_member"
	TargetIsOwner                                   = "target_is_owner"
	TargetIsSelf                                    = "target_is_self"
	TargetNotActive                                 = "target_not_active"
)

// GroupMembershipInfo is information about a group member of the shared folder.
type GroupMembershipInfo struct {
	AccessType struct {
		Tag AccessType `json:".tag"`
	} `json:"access_type"`
	Group       GroupInfo          `json:"group"`
	Permissions []MemberPermission `json:"permissions,omitempty"`
	Initials    *string            `json:"initials,omitempty"`
	IsInherited bool               `json:"is_inherited"`
}

// GroupInfo defines information about the membership group.
type GroupInfo struct {
	GroupName string `json:"group_name"`
	GroupID   string `json:"group_id"`
	GroupType struct {
		Tag GroupType `json:".tag"`
	} `json:"group_type"`
	IsOwner         bool    `json:"is_owner"`
	SameTeam        bool    `json:"same_team"`
	GroupExternalID string  `json:"group_external_id,omitempty"`
	MemberCount     *uint32 `json:"member_count,omitempty"`
}

// GroupType determines how a group is created and managed.
type GroupType string

// GroupType possible values.
const (
	GroupTypeTeam        GroupType = "team"
	GroupTypeUserManaged           = "user_managed"
)

// InviteeMembershipInfo is information about an invited member of a shared folder.
type InviteeMembershipInfo struct {
	AccessType struct {
		Tag AccessType `json:".tag"`
	} `json:"access_type"`
	Invitee     InviteeInfo        `json:"invitee"`
	Permissions []MemberPermission `json:"permissions,omitempty"`
	Initials    *string            `json:"initials,omitempty"`
	IsInherited bool               `json:"is_inherited"`
}

// InviteeInfo contains information about the recipient of a shared folder invitation.
type InviteeInfo struct {
	Email string `json:"email"`
}

// ListSharedFolderMembers returns shared folder membership by its folder ID.
func (c *Sharing) ListSharedFolderMembers(in *ListSharedFolderMembersInput) (out *ListSharedFolderMembersOutput, err error) {
	body, err := c.call("/sharing/list_folder_members", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// ListSharedFolderMembersInputContinue is the input for continuing listing folder members.
type ListSharedFolderMembersInputContinue struct {
	Cursor string `json:"cursor"`
}

// ListSharedFolderMembersContinue returns shared folder membership by its folder ID.
func (c *Sharing) ListSharedFolderMembersContinue(in *ListSharedFolderMembersInputContinue) (out *ListSharedFolderMembersOutput, err error) {
	body, err := c.call("/sharing/list_folder_members/continue", in)
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
	Actions []FolderAction `json:"actions,omitempty"`
}

// FolderAction defines actions that may be taken on shared folders.
type FolderAction struct {
	ChangeOptions string
}

// ListSharedFolderOutput lists metadata about shared folders with a cursor to retrieve the next page.
type ListSharedFolderOutput struct {
	Entries []SharedFolderMetadata `json:"entries"`
	Cursor  string                 `json:"cursor"`
}

// ListSharedFolders returns the list of all shared folders the current user has access to.
func (c *Sharing) ListSharedFolders(in *ListSharedFolderInput) (out *ListSharedFolderOutput, err error) {
	body, err := c.call("/sharing/list_folders", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// ListSharedFolderContinueInput request input.
type ListSharedFolderContinueInput struct {
	Cursor string `json:"cursor"`
}

// ListSharedFoldersContinue returns the list of all shared folders the current user has access to.
func (c *Sharing) ListSharedFoldersContinue(in *ListSharedFolderContinueInput) (out *ListSharedFolderOutput, err error) {
	body, err := c.call("/sharing/list_folders/continue", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// SharedFolderMetadata includes basic information about the shared folder.
type SharedFolderMetadata struct {
	AccessType struct {
		Tag AccessType `json:".tag"`
	} `json:"access_type"`
	IsTeamFolder   bool         `json:"is_team_folder"`
	Policy         FolderPolicy `json:"policy"`
	Name           string       `json:"name"`
	SharedFolderID string       `json:"shared_folder_id"`
	TimeInvited    time.Time    `json:"time_invited"`
	OwnerTeam      struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"owner_team"`
	ParentSharedFolderID string   `json:"parent_shared_folder_id"`
	PathLower            string   `json:"path_lower"`
	Permissions          []string `json:"permissions"`
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
