package dropbox

import (
	"encoding/json"
)

// Users client for user accounts.
type Users struct {
	*Client
}

// NewUsers client.
func NewUsers(config *Config) *Users {
	return &Users{
		Client: &Client{
			Config: config,
		},
	}
}

// GetAccountInput request input.
type GetAccountInput struct {
	AccountID string `json:"account_id"`
}

// GetAccountOutput request output.
type GetAccountOutput struct {
	AccountID string `json:"account_id"`
	Name      struct {
		GivenName    string `json:"given_name"`
		Surname      string `json:"surname"`
		FamiliarName string `json:"familiar_name"`
		DisplayName  string `json:"display_name"`
	} `json:"name"`
}

// GetAccount returns information about a user's account.
func (c *Users) GetAccount(in *GetAccountInput) (out *GetAccountOutput, err error) {
	body, err := c.call("/users/get_account", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetCurrentAccountOutput request output.
type GetCurrentAccountOutput struct {
	AccountID string `json:"account_id"`
	Name      struct {
		GivenName    string `json:"given_name"`
		Surname      string `json:"surname"`
		FamiliarName string `json:"familiar_name"`
		DisplayName  string `json:"display_name"`
	} `json:"name"`
	Email        string `json:"email"`
	Locale       string `json:"locale"`
	ReferralLink string `json:"referral_link"`
	IsPaired     bool   `json:"is_paired"`
	AccountType  struct {
		Tag string `json:".tag"`
	} `json:"account_type"`
	Country string `json:"country"`
}

// GetCurrentAccount returns information about the current user's account.
func (c *Users) GetCurrentAccount() (out *GetCurrentAccountOutput, err error) {
	body, err := c.call("/users/get_current_account", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// GetSpaceUsageOutput request output.
type GetSpaceUsageOutput struct {
	Used       uint64 `json:"used"`
	Allocation struct {
		Used      uint64 `json:"used"`
		Allocated uint64 `json:"allocated"`
	} `json:"allocation"`
}

// GetSpaceUsage returns space usage information for the current user's account.
func (c *Users) GetSpaceUsage() (out *GetSpaceUsageOutput, err error) {
	body, err := c.call("/users/get_space_usage", nil)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
