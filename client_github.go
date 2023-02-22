package configit

import (
	"context"
	"fmt"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

type GithubClientOpt struct {
	Token     string `yaml:"token" json:"token"`
	RepoOwner string `yaml:"repoOwner" json:"repoOwner"`
	Repo      string `yaml:"repo" json:"repo"`
	Branch    string `yaml:"branch" json:"branch"`
	RootPath  string `yaml:"rootPath"`
	Filter    func()
}

func (o *GithubClientOpt) Validate() error {

}

type githubClient struct {
	opt *GithubClientOpt
	cli *github.Client
}

func NewGithubClient(ctx context.Context, opt *GithubClientOpt, httpTransport *http.Transport) (Client, error) {
	if err := opt.Validate(); err != nil {
		return nil, fmt.Errorf("%w. %v", err, ErrInvalidConfig)
	}
	transport := http.DefaultTransport
	if httpTransport != nil {
		transport = httpTransport
	}
	cli := github.NewClient(
		&http.Client{
			Transport: &oauth2.Transport{
				Base:   transport,
				Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: opt.Token}),
			},
		})
	return &githubClient{
		opt: opt,
		cli: cli,
	}, nil
}

func (g *githubClient) Load(ctx context.Context) (Config, error) {
	_, dirContents, resp, err := g.cli.Repositories.GetContents(
		ctx, g.opt.RepoOwner, g.opt.Repo, "", nil)
	if err != nil {
		return nil, err
	}
	// TODO
	for _, item := range dirContents {
		if *item.Type == "file" {
			// Get the file contents.
			fileContent, _, _, err := g.cli.Repositories.GetContents(context.Background(), username, repoName, *item.Path, nil)
			if err != nil {
				return nil, err
			}
			err = ioutil.WriteFile(*item.Name, []byte(*fileContent.Content), 0644)
			if err != nil {
				return nil, err
			}
		}
	}

}
