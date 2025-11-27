package nestedroutines

import (
	"context"
	"fmt"
)

// AuthParam is a custom type for context keys to avoid collisions
type AuthParam string

// AuthParamKey is the key used to store auth parameters in context
var AuthParamKey AuthParam = "authParam"

func nestedPrimaryWorker(ctx context.Context, ch2 chan<- bool) {
	fmt.Println("nested worker started")
	val := ctx.Value(AuthParamKey)
	if val == nil {
		fmt.Println("requestId: <nil>")
	} else {
		fmt.Println("requestId", val.(string))
	}
	ch2 <- true
	fmt.Println("nested worker ended")
}

func primaryWorker(ctx context.Context, ch chan<- bool) {
	fmt.Println("primary worker started")
	ch2 := make(chan bool)
	// trying to overridde the value
	ctx = context.WithValue(ctx, AuthParamKey, "hrishiX_hecker")
	go nestedPrimaryWorker(ctx, ch2)

	<-ch2
	ch <- true
	fmt.Println("primary worker ended")
}

func Start(ctx context.Context) {
	fmt.Println("Starting the  nested routines application...")
	ch := make(chan bool)
	go primaryWorker(ctx, ch)
	<-ch
	fmt.Println("primary worker ended")
}
