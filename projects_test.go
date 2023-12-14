package jira

import (
	"context"
	"testing"

	"github.com/zdz1715/go-utils/goutils"

	"github.com/zdz1715/ghttp"
)

func TestProjectsService_ListProjects(t *testing.T) {
	client, err := NewClient(testBasicAuthCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Project.ListProjects(context.Background(), &ListProjectOptions{
		SearchOptions: &SearchOptions{
			StartAt:    0,
			MaxResults: 10,
			Expand:     "issueTypes",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", reply)
}

func TestProjectsService_Get(t *testing.T) {
	client, err := NewClient(testBasicAuthCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Project.Get(context.Background(), "10000", &GetProjectOptions{
		Expand: goutils.Ptr("issueTypes"),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", reply)
}
