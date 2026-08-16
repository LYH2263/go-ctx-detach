package rpcx

import (
	"context"
	"testing"
	"time"
)

func TestParentCancel(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		done <- Do(parent, func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(3 * time.Second):
				return nil
			}
		})
	}()
	time.Sleep(20 * time.Millisecond)
	cancel()
	select {
	case err := <-done:
		if err == nil {
			t.Fatal("want cancel error")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("did not observe parent cancel")
	}
}
