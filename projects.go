package jira

import (
	"context"
	"fmt"
	"net/http"
)

type ProjectsService service

// Project represents a Jira Project.
type Project struct {
	Expand          string             `json:"expand,omitempty" structs:"expand,omitempty"`
	Self            string             `json:"self,omitempty" structs:"self,omitempty"`
	ID              string             `json:"id,omitempty" structs:"id,omitempty"`
	IsPrivate       bool               `json:"is_private"`
	Key             string             `json:"key,omitempty" structs:"key,omitempty"`
	Description     string             `json:"description,omitempty" structs:"description,omitempty"`
	Lead            User               `json:"lead,omitempty" structs:"lead,omitempty"`
	Components      []ProjectComponent `json:"components,omitempty" structs:"components,omitempty"`
	IssueTypes      []IssueType        `json:"issueTypes,omitempty" structs:"issueTypes,omitempty"`
	URL             string             `json:"url,omitempty" structs:"url,omitempty"`
	Email           string             `json:"email,omitempty" structs:"email,omitempty"`
	AssigneeType    string             `json:"assigneeType,omitempty" structs:"assigneeType,omitempty"`
	Versions        []Version          `json:"versions,omitempty" structs:"versions,omitempty"`
	Name            string             `json:"name,omitempty" structs:"name,omitempty"`
	Roles           map[string]string  `json:"roles,omitempty" structs:"roles,omitempty"`
	AvatarUrls      AvatarUrls         `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
	ProjectCategory ProjectCategory    `json:"projectCategory,omitempty" structs:"projectCategory,omitempty"`
	ProjectTypeKey  string             `json:"projectTypeKey"`
}

type ListProjectOptions struct {
	*SearchOptions
	OrderBy    *string  `json:"orderBy,omitempty" query:"orderBy"`
	Query      *string  `json:"query,omitempty" query:"query"`
	Action     *string  `json:"action,omitempty" query:"action"`
	TypeKey    *string  `json:"typeKey,omitempty" query:"typeKey"`
	CategoryId *int64   `json:"categoryId,omitempty" query:"categoryId"`
	Id         []int    `json:"id,omitempty" query:"id"`
	Keys       []string `json:"keys,omitempty" query:"keys"`
}

// ListProjects Returns a paginated list of projects visible to the user.
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-search-get
func (s *ProjectsService) ListProjects(ctx context.Context, opts *ListProjectOptions) (*Pagination[Project], error) {
	const apiEndpoint = "/rest/api/2/project/search"
	var projects Pagination[Project]
	if err := s.client.Invoke(ctx, http.MethodGet, apiEndpoint, opts, &projects); err != nil {
		return nil, err
	}
	return &projects, nil
}

type GetProjectOptions struct {
	Expand     *string  `json:"expand,omitempty" query:"expand"`
	Properties []string `json:"properties,omitempty" query:"properties"`
}

// Get returns a full representation of the project for the given issue key.
// Jira will attempt to identify the project by the projectIdOrKey path parameter.
// This can be an project id, or an project key.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-get
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ProjectsService) Get(ctx context.Context, projectIdOrKey string, opts ...*GetProjectOptions) (*Project, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/project/%s", projectIdOrKey)
	var project Project
	if len(opts) > 0 && opts[0] != nil {
		if err := s.client.Invoke(ctx, http.MethodGet, apiEndpoint, opts[0], &project); err != nil {
			return nil, err
		}
		return &project, nil
	}

	if err := s.client.Invoke(ctx, http.MethodGet, apiEndpoint, nil, &project); err != nil {
		return nil, err
	}
	return &project, nil
}
