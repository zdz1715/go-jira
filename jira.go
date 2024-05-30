package jira

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/zdz1715/ghttp"
)

type service struct {
	client *Client
}

type Options struct {
	ClientOpts []ghttp.ClientOption
}

type Client struct {
	cc   *ghttp.Client
	opts *Options

	common service
	// Services used for talking to different parts of the Jira API.
	OAuth   *OAuthService
	User    *UsersService
	Issue   *IssuesService
	Project *ProjectsService
}

func NewClient(credential Credential, opts *Options) (*Client, error) {
	if opts == nil {
		opts = &Options{}
	}

	clientOptions := make([]ghttp.ClientOption, 0)

	if len(opts.ClientOpts) > 0 {
		clientOptions = append(clientOptions, opts.ClientOpts...)
	}

	clientOptions = append(clientOptions, ghttp.WithNot2xxError(func() ghttp.Not2xxError {
		return new(Error)
	}))

	cc := ghttp.NewClient(clientOptions...)

	c := &Client{
		cc:   cc,
		opts: opts,
	}

	c.common.client = c

	c.OAuth = &OAuthService{client: c.common.client}
	c.User = (*UsersService)(&c.common)
	c.Issue = (*IssuesService)(&c.common)
	c.Project = (*ProjectsService)(&c.common)

	if credential != nil {
		if err := c.SetCredential(credential); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) SetCredential(credential Credential) error {
	if credential == nil {
		return ErrCredential
	}

	if err := credential.Valid(); err != nil {
		return err
	}

	c.cc.SetEndpoint(credential.GetEndpoint())

	if c.OAuth != nil {
		c.OAuth.credential = credential
	}

	return nil
}

func (c *Client) Invoke(ctx context.Context, method, path string, args interface{}, reply interface{}) error {
	callOpts, err := c.OAuth.generateCallOptions()
	if err != nil {
		return err
	}

	if method == http.MethodGet && args != nil {
		callOpts.Query = args
		args = nil
	}

	_, err = c.cc.Invoke(ctx, method, path, args, reply, callOpts)
	return err
}

type Error struct {
	ErrorMessages []string    `json:"errorMessages"`
	Errors        interface{} `json:"errors"`
	Messages      string      `json:"message"`
}

func (e *Error) String() string {
	if e.Messages != "" {
		return e.Messages
	}

	if len(e.ErrorMessages) > 0 {
		return strings.Join(e.ErrorMessages, ",")
	}
	if e.Errors != nil {
		if b, err := json.Marshal(e.Errors); err == nil {
			return string(b)
		}
	}
	return ""
}

func (e *Error) Reset() {
	e.Errors = nil
	e.Messages = ""
	e.ErrorMessages = nil
}

// SearchOptions specifies the optional parameters to various List methods that
// support pagination.
// Pagination is used for the Jira REST APIs to conserve server resources and limit
// response size for resources that return potentially large collection of items.
// A request to a pages API will result in a values array wrapped in a JSON object with some paging metadata
// Default Pagination options
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/intro/#expansion
type SearchOptions struct {
	// StartAt: The starting index of the returned projects. Base index: 0.
	StartAt int `query:"startAt,omitempty"`
	// MaxResults: The maximum number of projects to return per page. Default: 50.
	MaxResults int `query:"maxResults,omitempty"`
	// Expand: Expand specific sections in the returned issues
	Expand string `query:"expand,omitempty"`
}

type Pagination[T any] struct {
	StartAt    int  `json:"startAt"`
	MaxResults int  `json:"maxResults"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []*T `json:"values"`
}
