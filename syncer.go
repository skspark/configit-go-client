package configit

import "context"

type Syncer struct {
	opt *SyncerOpt
	cli Client
}

type SyncerOpt struct {
}

func NewSyncer(ctx context.Context, opt *SyncerOpt, cli Client) (*Syncer, error) {
	return &Syncer{
		opt: opt,
		cli: cli,
	}, nil
}

func (s *Syncer) Start(ctx context.Context,
	onSuccess func(ctx context.Context, newConfig *Config),
	onFailure func(ctx context.Context, err error)) error {
	// TODO
}

func (s *Syncer) Stop(ctx context.Context) error {
	// TODO
}
