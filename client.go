package configit

import (
	"context"
	"fmt"
)

var (
	ErrSyncTimeout   = fmt.Errorf("sync timeout")
	ErrInvalidConfig = fmt.Errorf("invalid config")
)

type Client interface {
	Load(ctx context.Context) (Config, error)
}

type Config interface {
	Field(ctx context.Context, path string, target interface{}) error
}
