package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	// try to comment line 12-14, see the result
	ctx = AppendCtxLog(ctx, map[string]interface{}{
		"init": 123,
	})

	fmt.Println("Main ->", ctx.Value("logs"))

	NextFlow(ctx)
	fmt.Println("Back to Main Flow 1->", ctx.Value("logs"))

	NexFlow2(ctx)
	fmt.Println("Back to Main Flow 2->", ctx.Value("logs"))

	fmt.Println("GrandChildren", ctx.Value("grand-child"))
}

func NextFlow(ctx context.Context) {
	ctx = AppendCtxLog(ctx, map[string]interface{}{
		"b1": 789,
	})

	NextDeepFlow(ctx)
}

func NextDeepFlow(ctx context.Context) {
	ctx = AppendCtxLog(ctx, map[string]interface{}{
		"b2": 111,
	})
}

func NexFlow2(ctx context.Context) {
	ctx = AppendCtxLog(ctx, map[string]interface{}{
		"c1": 1,
	})

	ctx = context.WithValue(ctx, "grand-child", "flow2")

	NextDeepFlow2(ctx)
}

func NextDeepFlow2(ctx context.Context) {
	ctx = AppendCtxLog(ctx, map[string]interface{}{
		"c2": 2,
	})
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
