package configit

import (
	"context"
	"github.com/go-git/go-git/v5"
	"time"
)

type Client interface {
}

type GithubClientConfig struct {
}

func NewGithubClient(ctx context.Context, conf *GithubClientConfig) (Client, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := git.NewClient(tc)

	content := make(chan []byte)
	duration := time.Minute * 5 // 5 minutes

	go syncRepoToMemory(client, owner, repo, content, duration)

	for {
		select {
		case latestContent := <-content:
			// Use the latest content in your application
		}
	}
}


