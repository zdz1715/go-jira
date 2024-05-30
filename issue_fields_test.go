package jira

import (
	"context"
	"testing"

	"github.com/zdz1715/ghttp"
)

func TestIssuesService_GetFields(t *testing.T) {
	client, err := NewClient(testBasicAuthCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Issue.GetFields(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", reply)
}
