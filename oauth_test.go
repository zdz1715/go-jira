package jira

import "os"

var testBasicAuthCredential = &BasicAuth{
	Endpoint: os.Getenv("TEST_JIRA_SERVER_URL"),
	Username: os.Getenv("TEST_JIRA_USERNAME"),
	Password: os.Getenv("TEST_JIRA_PASSWORD"),
}
