package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = AppendCtxLog(ctx, map[string]interface{}{
		"a": 123,
	})

	ctx = AppendCtxLog(ctx, map[string]interface{}{
		"b": 456,
		"x": 999,
	})

	fmt.Println("Main ->", ctx.Value("logs"))

	NextFlow(ctx)
}

func NextFlow(ctx context.Context) {
	ctx = AppendCtxLog(ctx, map[string]interface{}{
		"c": 789,
	})

	fmt.Println("Next flow ->", ctx.Value("logs"))
}

func AppendCtxLog(ctx context.Context, fields map[string]interface{}) context.Context {

	var logs map[string]interface{}
	var ok bool

	if logs, ok = ctx.Value("logs").(map[string]interface{}); !ok {
		logs = make(map[string]interface{})
	}

	for k, v := range fields {
		logs[k] = v
	}

	return context.WithValue(ctx, "logs", logs)
}
