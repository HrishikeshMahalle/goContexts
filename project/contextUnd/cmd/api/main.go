package main

import (
	"context"
	"fmt"
	"net/http"

	core "github.com/hrishikeshmahalle/contextUnd/core/cancelRoutine"
	handler "github.com/hrishikeshmahalle/contextUnd/handler"
)

func main() {
	// 1
	// // using context.WithCancel to cancel the context after 3 seconds
	// ctx, cancel := context.WithCancel(context.Background())

	// time.AfterFunc(3*time.Second, cancel) // using time.AfterFunc to cancel the context after 3 seconds

	// 2
	// this extents the behaviour of context.WithCancel with timeout functionality
	// we can override the timeout by calling the cancel function
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // using context.WithTimeout to cancel the context after 3 seconds

	// // calling the cancel func
	// time.AfterFunc(2*time.Second, cancel)

	// 3
	// context.WithDeadline extends the timeout behaviour with a specific deadline
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(0*time.Second))
	// // // calling the cancel func
	// time.AfterFunc(3*time.Second, cancel)

	// core.Start(ctx)

	// 4 understanding context values
	// understandingContextValues()

	// 5 understading context with http router
	// understandingWithHttpRouter()

	ctx, cancel := context.WithCancel(context.Background())
	core.StartTwo(ctx, cancel)

}

// this was done to avoid colloins with the standard library inside WithValue function
// // AuthParam is a custom type for context keys to avoid collisions
// type AuthParam string

// // AuthParamKey is the key used to store auth parameters in context
// var AuthParamKey AuthParam = "authParam"

// func understandingContextValues() {
// 	authParamKeyVal := "hrishiAB1001"
// 	ctx := context.WithValue(context.Background(), nestedroutines.AuthParamKey, authParamKeyVal)
// 	nestedroutines.Start(ctx)
// }

// 5
// understading with http router
// 1) run the program
// 2) Hit the api endpoint http://localhost:8080/work using "curl localhost:8080/work" in terminal
// 3) If you wait for 5 seconds youll see the output hello world!!!
// 4) If you cancel the request by pressing ctrl+c in terminal youll see the output "Context cancelled"
func understandingWithHttpRouter() {
	http.HandleFunc("/work", handler.Handler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
