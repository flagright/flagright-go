# flagright-go
The Flagright Go library provides access to the Flagright API from Golang. 

### Usage

#### Create a consumer user

```package main

import (
	"context"
	"fmt"
	"time"

	"github.com/flagright/flagright-go/api"
	"github.com/flagright/flagright-go/api/consumerusers"
	"github.com/flagright/flagright-go/api/option"
)

func main() {
	consumerClient := consumerusers.NewClient(
		option.WithBaseURL("https://api.flagright.dev"), // modify this for sandbox or production
		option.WithApiKey("YOUR_API_KEY"),
	)
	var user api.User = api.User{
		UserId:           "user123-go",
		CreatedTimestamp: float64(time.Now().Unix()),
	}

	req := &api.ConsumerUsersCreateRequest{
		Body: &user,
	}

	_, err := consumerClient.Create(context.Background(), req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
```
