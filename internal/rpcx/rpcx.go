package rpcx

import (
	"context"
	"time"
)

func Do(parent context.Context, work func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // BUG
	defer cancel()
	_ = parent
	return work(ctx)
}
