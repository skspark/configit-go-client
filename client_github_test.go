package configit

import (
	"testing"
	"time"
)

func Test_client(t *testing.T) {
	ctx := context.Background()
	c, err := NewGithubClient(ctx, &GithubClientOpt{
		endpoint: "",
		repo:     "",
		secret:   "",
	})
	if err != nil {
		panic(err)
	}
	// simply load on startup
	cfg, err := c.Load(ctx)
	if err != nil {
		panic(err)
	}
	type (
		testType struct {
			a string
			b string
		}
	)
	tt := testType{}
	if err := cfg.Field(ctx, "foo.bar", &tt); err != nil {
		panic(err)
	}

	// sync periodically from storage
	s, err := c.Syncer(ctx, &SyncerOpt{
		interval: time.Duration,
		timeout: time.Duration,
	})
	if err != nil {
		panic(err)
	}
	s.Start(ctx, func(ctx context.Context, newConfig *Config) {

	}, func())
	s.Stop(ctx)

}
