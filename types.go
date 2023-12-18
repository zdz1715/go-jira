package jira

// AccountType
// docs: https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/#webhooks
type AccountType string

const (
	AtlassianAccountType = "atlassian"
	AppAccountType       = "app"
	CustomerAccountType  = "customer"
)

type IssueTypeLevel int32

const (
	SubtaskIssueTypLevel = iota - 1
	BaseIssueTypLevel
	EpicIssueTypLevel
)
