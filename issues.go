package jira

import (
	"context"
	"net/http"
)

type IssuesService service

// Issue represents a Jira issue.
type Issue struct {
	Expand      string            `json:"expand,omitempty" structs:"expand,omitempty"`
	ID          string            `json:"id,omitempty" structs:"id,omitempty"`
	Self        string            `json:"self,omitempty" structs:"self,omitempty"`
	Key         string            `json:"key,omitempty" structs:"key,omitempty"`
	Changelog   *Changelog        `json:"changelog,omitempty" structs:"changelog,omitempty"`
	Transitions []Transition      `json:"transitions,omitempty" structs:"transitions,omitempty"`
	Names       map[string]string `json:"names,omitempty" structs:"names,omitempty"`
}

// IssueFields represents single fields of a Jira issue.
// Every Jira issue has several fields attached.

// TransitionField represents the value of one Transition
type TransitionField struct {
	Required bool `json:"required" structs:"required"`
}

// Transition represents an issue transition in Jira
type Transition struct {
	ID     string                     `json:"id" structs:"id"`
	Name   string                     `json:"name" structs:"name"`
	To     Status                     `json:"to" structs:"status"`
	Fields map[string]TransitionField `json:"fields" structs:"fields"`
}

// Wrapper struct for search result
type transitionResult struct {
	Transitions []Transition `json:"transitions" structs:"transitions"`
}

// ChangelogItems reflects one single changelog item of a history item
type ChangelogItems struct {
	Field      string `json:"field" structs:"field"`
	FieldId    string `json:"fieldId"`
	FieldType  string `json:"fieldtype" structs:"fieldtype"`
	From       string `json:"from" structs:"from"`
	FromString string `json:"fromString" structs:"fromString"`
	To         string `json:"to" structs:"to"`
	ToString   string `json:"toString" structs:"toString"`
}

// ChangelogHistory reflects one single changelog history entry
type ChangelogHistory struct {
	Id      string           `json:"id" structs:"id"`
	Author  User             `json:"author" structs:"author"`
	Created string           `json:"created" structs:"created"`
	Items   []ChangelogItems `json:"items" structs:"items"`
}

// Changelog reflects the change log of an issue
type Changelog struct {
	Histories []ChangelogHistory `json:"histories,omitempty"`
}

type CreateIssueOptions struct {
	//Fields     *IssueFields     `json:"fields,omitempty" query:"fields"`
	//Properties []EntityProperty `json:"properties,omitempty" query:"properties"`
}

// Create creates an issue in Jira.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-post
func (s *IssuesService) Create(ctx context.Context, opts *CreateIssueOptions) (*Issue, error) {
	const apiEndpoint = "/rest/api/2/issue"
	var issue Issue
	if err := s.client.Invoke(ctx, http.MethodPost, apiEndpoint, opts, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

type GetProjectIssueTypeOptions struct {
	ProjectId *string         `json:"projectId,omitempty" query:"projectId"`
	Level     *IssueTypeLevel `json:"level,omitempty" query:"level"`
}

// GetProjectIssueType
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-types/#api-rest-api-2-issuetype-project-get
func (s *IssuesService) GetProjectIssueType(ctx context.Context, opts *GetProjectIssueTypeOptions) ([]*IssueType, error) {
	const apiEndpoint = "/rest/api/2/issuetype/project"
	var issueType []*IssueType
	if err := s.client.Invoke(ctx, http.MethodGet, apiEndpoint, opts, &issueType); err != nil {
		return nil, err
	}
	return issueType, nil
}
