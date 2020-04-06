package context_sample

import (
	"context"
	"fmt"
	"time"
)

func ContextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

}

func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "Monitor1")
	go watch(ctx, "Monitor2")
	go watch(ctx, "Monitor3")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)

}

func ContextWithValue() {
	i := 0
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, "key", i)
	go handlevalue(valueCtx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)

}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func handlevalue(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出")
			return
		default:
			fmt.Println(ctx.Value("key"))
			time.Sleep(1 * time.Second)
		}
	}

}
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了")
			return
		default:
			fmt.Println(name, "监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func DoBissnis() {
	for {
		fmt.Println(time.Now())
		time.Sleep(2 * time.Second)
	}

}
