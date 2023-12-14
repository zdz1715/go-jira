package jira

import (
	"context"
	"net/http"
)

type UsersService service

// User represents a Jira user.
type User struct {
	Self             string           `json:"self,omitempty" structs:"self,omitempty"`
	AccountID        string           `json:"accountId,omitempty" structs:"accountId,omitempty"`
	AccountType      string           `json:"accountType,omitempty" structs:"accountType,omitempty"`
	Name             string           `json:"name,omitempty" structs:"name,omitempty"`
	Key              string           `json:"key,omitempty" structs:"key,omitempty"`
	Password         string           `json:"-"`
	EmailAddress     string           `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	AvatarUrls       AvatarUrls       `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
	DisplayName      string           `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active           bool             `json:"active,omitempty" structs:"active,omitempty"`
	TimeZone         string           `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	Locale           string           `json:"locale,omitempty" structs:"locale,omitempty"`
	Groups           UserGroups       `json:"groups,omitempty" structs:"groups,omitempty"`
	ApplicationRoles ApplicationRoles `json:"applicationRoles,omitempty" structs:"applicationRoles,omitempty"`
}

// AvatarUrls represents different dimensions of avatars / images
type AvatarUrls struct {
	Four8X48  string `json:"48x48,omitempty" structs:"48x48,omitempty"`
	Two4X24   string `json:"24x24,omitempty" structs:"24x24,omitempty"`
	One6X16   string `json:"16x16,omitempty" structs:"16x16,omitempty"`
	Three2X32 string `json:"32x32,omitempty" structs:"32x32,omitempty"`
}

// UserGroup represents the group list
type UserGroup struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

// UserGroups is a wrapper for UserGroup
type UserGroups struct {
	Size  int         `json:"size,omitempty" structs:"size,omitempty"`
	Items []UserGroup `json:"items,omitempty" structs:"items,omitempty"`
}

// ApplicationRoles is a wrapper for ApplicationRole
type ApplicationRoles struct {
	Size  int               `json:"size,omitempty" structs:"size,omitempty"`
	Items []ApplicationRole `json:"items,omitempty" structs:"items,omitempty"`
}

// ApplicationRole represents a role assigned to a user
type ApplicationRole struct {
	Key                  string   `json:"key"`
	Groups               []string `json:"groups"`
	Name                 string   `json:"name"`
	DefaultGroups        []string `json:"defaultGroups"`
	SelectedByDefault    bool     `json:"selectedByDefault"`
	Defined              bool     `json:"defined"`
	NumberOfSeats        int      `json:"numberOfSeats"`
	RemainingSeats       int      `json:"remainingSeats"`
	UserCount            int      `json:"userCount"`
	UserCountDescription string   `json:"userCountDescription"`
	HasUnlimitedSeats    bool     `json:"hasUnlimitedSeats"`
	Platform             bool     `json:"platform"`

	// Key `groupDetails` missing - https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-application-roles/#api-rest-api-3-applicationrole-key-get
	// Key `defaultGroupsDetails` missing - https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-application-roles/#api-rest-api-3-applicationrole-key-get
}

// GetCurrentUser returns details for the current user.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-myself/#api-rest-api-2-myself-get
func (s *UsersService) GetCurrentUser(ctx context.Context) (*User, error) {
	const apiEndpoint = "/rest/api/2/myself"
	var user User
	err := s.client.Invoke(context.Background(), http.MethodGet, apiEndpoint, nil, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers returns all user.
// 包括活动用户、非活动用户和以前删除的拥有Atlassian 帐户的用户。
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-users/#api-rest-api-2-users-get
func (s *UsersService) GetAllUsers(ctx context.Context, search *SearchOptions) ([]*User, error) {
	const apiEndpoint = "/rest/api/2/users"
	var user []*User
	if err := s.client.Invoke(ctx, http.MethodGet, apiEndpoint, search, &user); err != nil {
		return nil, err
	}
	return user, nil
}

type FindUsersOptions struct {
	*SearchOptions

	// Query field will search users displayName and emailAddress
	Query     *string `json:"query,omitempty" query:"query"`
	Username  *string `json:"username,omitempty" query:"username"`
	AccountId *string `json:"accountId,omitempty" query:"accountId"`
	Property  *string `json:"property,omitempty" query:"property"`
}

// FindUsers searches for user info from Jira:
// It can find users by email or display name using the query parameter
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-user-search/#api-rest-api-2-user-search-get
func (s *UsersService) FindUsers(ctx context.Context, req *FindUsersOptions) ([]*User, error) {
	const apiEndpoint = "/rest/api/2/user/search"
	var user []*User
	if err := s.client.Invoke(ctx, http.MethodGet, apiEndpoint, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

type FindUsersByQueryOptions struct {
	*SearchOptions

	// JQL https://www.atlassian.com/zh/software/jira/guides/jql/overview
	// `is assignee of PROJ` Returns the users that are assignees of at least one issue in project PROJ.
	// `is assignee of (PROJ-1, PROJ-2)` Returns users that are assignees on the issues PROJ-1 or PROJ-2.
	Query *string `json:"query,omitempty" query:"query"`
}

// FindUsersByQuery
// Finds users with a structured query and returns a paginated list of user details.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-user-search/#api-rest-api-2-user-search-query-get
func (s *UsersService) FindUsersByQuery(ctx context.Context, req *FindUsersByQueryOptions) ([]*User, error) {
	const apiEndpoint = "/rest/api/2/user/search/query"
	var user []*User
	if err := s.client.Invoke(ctx, http.MethodGet, apiEndpoint, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

type CreateUserOptions struct {
	EmailAddress *string  `json:"emailAddress,omitempty" query:"emailAddress"`
	Products     []string `json:"products,omitempty" query:"products"`
}

// Create creates an user in Jira.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/user-createUser
func (s *UsersService) Create(ctx context.Context, opts *CreateUserOptions) (*User, error) {
	const apiEndpoint = "/rest/api/2/user"
	var user User
	if err := s.client.Invoke(ctx, http.MethodPost, apiEndpoint, opts, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
