package jira

import (
	"errors"
	"github.com/zdz1715/ghttp"
)

var (
	ErrCredential = errors.New("invalid credential")
)

type Credential interface {
	GetEndpoint() string
	GenerateCallOptions() (*ghttp.CallOptions, error)
	Valid() error
}

// BasicAuth
// Jira docs: https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/
// Create a new API token: https://id.atlassian.com/manage-profile/security/api-tokens
type BasicAuth struct {
	Endpoint string `json:"endpoint" xml:"endpoint"`
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

func (ba *BasicAuth) GetEndpoint() string {
	return ba.Endpoint
}

func (ba *BasicAuth) GenerateCallOptions() (*ghttp.CallOptions, error) {
	return &ghttp.CallOptions{
		Username: ba.Username,
		Password: ba.Password,
	}, nil
}

func (ba *BasicAuth) Valid() error {
	if ba.Endpoint == "" || ba.Username == "" || ba.Password == "" {
		return ErrCredential
	}
	return nil
}
