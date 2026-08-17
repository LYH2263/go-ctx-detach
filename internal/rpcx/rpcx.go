package rpcx

import (
	"context"
	"time"
)

func Do(parent context.Context, work func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	return work(ctx)
}
