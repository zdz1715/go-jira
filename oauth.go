package jira

import (
	"fmt"
	"github.com/zdz1715/ghttp"
)

type OAuthService struct {
	client     *Client
	credential Credential
}

func (o *OAuthService) generateCallOptions() (*ghttp.CallOptions, error) {
	if o.credential == nil {
		return nil, fmt.Errorf("nil Credential")
	}
	if err := o.credential.Valid(); err != nil {
		return nil, err
	}
	return o.credential.GenerateCallOptions()
}
