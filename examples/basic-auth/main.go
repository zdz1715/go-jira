package main

import (
	"context"
	"fmt"
	"github.com/zdz1715/go-jira"
	"os"
)

func main() {
	credential := &jira.BasicAuth{
		Endpoint: "Endpoint",
		Username: "YourAppKey",
		Password: "YourAppSecret",
	}

	client, err := jira.NewClient(credential, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 获取当前用户信息
	currentUser, err := client.User.GetCurrentUser(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("current user: %+v\n", currentUser)
}
