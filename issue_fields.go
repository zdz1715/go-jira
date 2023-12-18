package jira

import (
	"context"
	"net/http"
	"time"
)

type Field struct {
	ID          string      `json:"id,omitempty" structs:"id,omitempty"`
	Key         string      `json:"key,omitempty" structs:"key,omitempty"`
	Name        string      `json:"name,omitempty" structs:"name,omitempty"`
	Custom      bool        `json:"custom,omitempty" structs:"custom,omitempty"`
	Navigable   bool        `json:"navigable,omitempty" structs:"navigable,omitempty"`
	Orderable   bool        `json:"orderable,omitempty"`
	Searchable  bool        `json:"searchable,omitempty" structs:"searchable,omitempty"`
	ClauseNames []string    `json:"clauseNames,omitempty" structs:"clauseNames,omitempty"`
	Schema      FieldSchema `json:"schema,omitempty" structs:"schema,omitempty"`
}

type FieldSchema struct {
	Type     string `json:"type,omitempty" structs:"type,omitempty"`
	Items    string `json:"items,omitempty" structs:"items,omitempty"`
	Custom   string `json:"custom,omitempty" structs:"custom,omitempty"`
	System   string `json:"system,omitempty" structs:"system,omitempty"`
	CustomID int64  `json:"customId,omitempty" structs:"customId,omitempty"`
}

func (s *IssuesService) GetFields(ctx context.Context) ([]*Field, error) {
	const apiEndpoint = "/rest/api/2/field"
	var result []*Field
	if err := s.client.Invoke(ctx, http.MethodGet, apiEndpoint, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Watches represents a type of how many and which user are "observing" a Jira issue to track the status / updates.
type Watches struct {
	Self       string     `json:"self,omitempty" structs:"self,omitempty"`
	WatchCount int        `json:"watchCount,omitempty" structs:"watchCount,omitempty"`
	IsWatching bool       `json:"isWatching,omitempty" structs:"isWatching,omitempty"`
	Watchers   []*Watcher `json:"watchers,omitempty" structs:"watchers,omitempty"`
}

// Watcher represents a simplified user that "observes" the issue
type Watcher struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	AccountID   string `json:"accountId,omitempty" structs:"accountId,omitempty"`
	DisplayName string `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active      bool   `json:"active,omitempty" structs:"active,omitempty"`
}

// Component represents a "component" of a Jira issue.
// Components can be user defined in every Jira instance.
type Component struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
}

// Progress represents the progress of a Jira issue.
type Progress struct {
	Progress int `json:"progress" structs:"progress"`
	Total    int `json:"total" structs:"total"`
	Percent  int `json:"percent" structs:"percent"`
}

// TimeTracking represents the timetracking fields of a Jira issue.
type TimeTracking struct {
	OriginalEstimate         string `json:"originalEstimate,omitempty" structs:"originalEstimate,omitempty"`
	RemainingEstimate        string `json:"remainingEstimate,omitempty" structs:"remainingEstimate,omitempty"`
	TimeSpent                string `json:"timeSpent,omitempty" structs:"timeSpent,omitempty"`
	OriginalEstimateSeconds  int    `json:"originalEstimateSeconds,omitempty" structs:"originalEstimateSeconds,omitempty"`
	RemainingEstimateSeconds int    `json:"remainingEstimateSeconds,omitempty" structs:"remainingEstimateSeconds,omitempty"`
	TimeSpentSeconds         int    `json:"timeSpentSeconds,omitempty" structs:"timeSpentSeconds,omitempty"`
}

// Subtasks represents all issues of a parent issue.
type Subtasks struct {
	ID     string      `json:"id" structs:"id"`
	Key    string      `json:"key" structs:"key"`
	Self   string      `json:"self" structs:"self"`
	Fields IssueFields `json:"fields" structs:"fields"`
}

// FixVersion represents a software release in which an issue is fixed.
type FixVersion struct {
	Self            string `json:"self,omitempty" structs:"self,omitempty"`
	ID              string `json:"id,omitempty" structs:"id,omitempty"`
	Name            string `json:"name,omitempty" structs:"name,omitempty"`
	Description     string `json:"description,omitempty" structs:"description,omitempty"`
	Archived        *bool  `json:"archived,omitempty" structs:"archived,omitempty"`
	Released        *bool  `json:"released,omitempty" structs:"released,omitempty"`
	ReleaseDate     string `json:"releaseDate,omitempty" structs:"releaseDate,omitempty"`
	UserReleaseDate string `json:"userReleaseDate,omitempty" structs:"userReleaseDate,omitempty"`
	ProjectID       int    `json:"projectId,omitempty" structs:"projectId,omitempty"` // Unlike other IDs, this is returned as a number
	StartDate       string `json:"startDate,omitempty" structs:"startDate,omitempty"`
}

// Epic represents the epic to which an issue is associated
// Not that this struct does not process the returned "color" value
type Epic struct {
	ID      int    `json:"id" structs:"id"`
	Key     string `json:"key" structs:"key"`
	Self    string `json:"self" structs:"self"`
	Name    string `json:"name" structs:"name"`
	Summary string `json:"summary" structs:"summary"`
	Done    bool   `json:"done" structs:"done"`
}

// Sprint represents a sprint on Jira agile board
type Sprint struct {
	ID            int        `json:"id" structs:"id"`
	Name          string     `json:"name" structs:"name"`
	CompleteDate  *time.Time `json:"completeDate" structs:"completeDate"`
	EndDate       *time.Time `json:"endDate" structs:"endDate"`
	StartDate     *time.Time `json:"startDate" structs:"startDate"`
	OriginBoardID int        `json:"originBoardId" structs:"originBoardId"`
	Self          string     `json:"self" structs:"self"`
	State         string     `json:"state" structs:"state"`
	Goal          string     `json:"goal,omitempty" structs:"goal"`
}

// Parent represents the parent of a Jira issue, to be used with subtask issue types.
type Parent struct {
	ID  string `json:"id,omitempty" structs:"id,omitempty"`
	Key string `json:"key,omitempty" structs:"key,omitempty"`
}

type IssueFields struct {
	Expand                        string        `json:"expand,omitempty" structs:"expand,omitempty"`
	Type                          IssueType     `json:"issuetype,omitempty" structs:"issuetype,omitempty"`
	Project                       Project       `json:"project,omitempty" structs:"project,omitempty"`
	Environment                   string        `json:"environment,omitempty" structs:"environment,omitempty"`
	Resolution                    *Resolution   `json:"resolution,omitempty" structs:"resolution,omitempty"`
	Priority                      *Priority     `json:"priority,omitempty" structs:"priority,omitempty"`
	Resolutiondate                *time.Time    `json:"resolutiondate,omitempty" structs:"resolutiondate,omitempty"`
	Created                       *time.Time    `json:"created,omitempty" structs:"created,omitempty"`
	Duedate                       *time.Time    `json:"duedate,omitempty" structs:"duedate,omitempty"`
	Watches                       *Watches      `json:"watches,omitempty" structs:"watches,omitempty"`
	Assignee                      *User         `json:"assignee,omitempty" structs:"assignee,omitempty"`
	Updated                       *time.Time    `json:"updated,omitempty" structs:"updated,omitempty"`
	Description                   string        `json:"description,omitempty" structs:"description,omitempty"`
	Summary                       string        `json:"summary,omitempty" structs:"summary,omitempty"`
	Creator                       *User         `json:"Creator,omitempty" structs:"Creator,omitempty"`
	Reporter                      *User         `json:"reporter,omitempty" structs:"reporter,omitempty"`
	Components                    []*Component  `json:"components,omitempty" structs:"components,omitempty"`
	Status                        *Status       `json:"status,omitempty" structs:"status,omitempty"`
	Progress                      *Progress     `json:"progress,omitempty" structs:"progress,omitempty"`
	AggregateProgress             *Progress     `json:"aggregateprogress,omitempty" structs:"aggregateprogress,omitempty"`
	TimeTracking                  *TimeTracking `json:"timetracking,omitempty" structs:"timetracking,omitempty"`
	TimeSpent                     int           `json:"timespent,omitempty" structs:"timespent,omitempty"`
	TimeEstimate                  int           `json:"timeestimate,omitempty" structs:"timeestimate,omitempty"`
	TimeOriginalEstimate          int           `json:"timeoriginalestimate,omitempty" structs:"timeoriginalestimate,omitempty"`
	Worklog                       *Worklog      `json:"worklog,omitempty" structs:"worklog,omitempty"`
	IssueLinks                    []*IssueLink  `json:"issuelinks,omitempty" structs:"issuelinks,omitempty"`
	Comments                      *Comments     `json:"comment,omitempty" structs:"comment,omitempty"`
	FixVersions                   []*FixVersion `json:"fixVersions,omitempty" structs:"fixVersions,omitempty"`
	AffectsVersions               []*Version    `json:"versions,omitempty" structs:"versions,omitempty"`
	Labels                        []string      `json:"labels,omitempty" structs:"labels,omitempty"`
	Subtasks                      []*Subtasks   `json:"subtasks,omitempty" structs:"subtasks,omitempty"`
	Attachments                   []*Attachment `json:"attachment,omitempty" structs:"attachment,omitempty"`
	Epic                          *Epic         `json:"epic,omitempty" structs:"epic,omitempty"`
	Sprint                        *Sprint       `json:"sprint,omitempty" structs:"sprint,omitempty"`
	Parent                        *Parent       `json:"parent,omitempty" structs:"parent,omitempty"`
	AggregateTimeOriginalEstimate int           `json:"aggregatetimeoriginalestimate,omitempty" structs:"aggregatetimeoriginalestimate,omitempty"`
	AggregateTimeSpent            int           `json:"aggregatetimespent,omitempty" structs:"aggregatetimespent,omitempty"`
	AggregateTimeEstimate         int           `json:"aggregatetimeestimate,omitempty" structs:"aggregatetimeestimate,omitempty"`
}
