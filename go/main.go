package main

import "github.com/cockroachdb/cockroach/pkg/util/log"

import "context"

func main() {
	// ctx := context.WithValue(context.Background(), "name", "yaojingguo")
	ac := log.AmbientContext{}
	ac.AddLogTag("name", "yaojingguo")
	ctx := ac.AnnotateCtx(context.Background())
	log.Dev.Info(ctx, "prepare to repel boarders")
	// log.Ops.Info(ctx, "prepare to repel boarders")
	// log.Dev.Fatal(ctx, "initialization failed")
	// log.Dev.Infof(ctx, "client error: %s")
}
