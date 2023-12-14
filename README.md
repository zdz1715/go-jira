# go-jira
Jira Go SDK

## Contents
- [Installation](#Installation)
- [Quick start](#quick-start)
- [ToDo](#todo)

## Installation
```shell
go get -u github.com/zdz1715/go-jira@latest
```
## Quick start
```go
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
```

## ToDo
> [!NOTE]
> 现在提供的方法不多，会逐渐完善，也欢迎来贡献代码，只需要编写参数结构体、响应结构体就可以很快添加一个方法，参考下方代码。
```go
// GetCurrentUser returns details for the current user.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-myself/#api-rest-api-2-myself-get
func (s *UsersService) GetCurrentUser(ctx context.Context) (*User, error) {
    const apiEndpoint = "/rest/api/2/myself"
    var user User
    err := s.client.Invoke(context.Background(), http.MethodGet, apiEndpoint, nil, &user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
```