package jira

import (
	"context"
	"testing"

	"github.com/zdz1715/ghttp"
	"github.com/zdz1715/go-utils/goutils"
)

func TestIssuesService_GetProjectIssueType(t *testing.T) {
	client, err := NewClient(testBasicAuthCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Issue.GetProjectIssueType(context.Background(), &GetProjectIssueTypeOptions{
		ProjectId: goutils.Ptr("10000"),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", reply)
}

func TestIssuesService_GetCreateMetadataForProject(t *testing.T) {
	client, err := NewClient(testBasicAuthCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Issue.GetCreateMetadataForProject(context.Background(), "10000", &SearchOptions{
		StartAt:    0,
		MaxResults: 100,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", reply)
}
