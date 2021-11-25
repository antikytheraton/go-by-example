package main

import (
	"context"
	"time"
)

func worker(ctx context.Context, name string, pollInterval time.Duration) {
	poller := time.NewTicker(pollInterval)
	defer poller.Stop()
}
