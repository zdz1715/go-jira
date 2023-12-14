package jira

import (
	"context"
	"testing"

	"github.com/zdz1715/ghttp"
	"github.com/zdz1715/go-utils/goutils"
)

func TestUsersService_GetAllUsers(t *testing.T) {
	client, err := NewClient(testBasicAuthCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.User.GetAllUsers(context.Background(), &SearchOptions{
		StartAt:    0, //  StartAt + MaxResults
		MaxResults: 100,
		//Expand:     "emailAddress",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", reply)
}

func TestUsersService_FindUsers(t *testing.T) {
	client, err := NewClient(testBasicAuthCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.User.FindUsers(context.Background(), &FindUsersOptions{
		Query: goutils.Ptr("."),
		SearchOptions: &SearchOptions{
			StartAt:    0,
			MaxResults: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("len: %d, users: %+v", len(reply), reply)
}

func TestUsersService_FindUsersByQuery(t *testing.T) {
	client, err := NewClient(testBasicAuthCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.User.FindUsersByQuery(context.Background(), &FindUsersByQueryOptions{
		Query: goutils.Ptr("."),
		SearchOptions: &SearchOptions{
			StartAt:    0,
			MaxResults: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("len: %d, users: %+v", len(reply), reply)
}
