package main

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
  // githuboauth "golang.org/x/oauth2/github"
	"context"
	"fmt"
	"os"
)

func main() {
  ctx := context.Background()
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "a66d6d45299cf1158ae8589c7d6795df780f3790"},
  )
  tc := oauth2.NewClient(ctx, ts)
  // get go-github client
  client := github.NewClient(tc)

  repo, _, err := client.Repositories.List(ctx, "EbookFoundation", nil)

  if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}
  fmt.Println(repo, err)
}
