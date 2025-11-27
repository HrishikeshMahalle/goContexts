package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handler started")

	ctx := req.Context()
	checkString := make(chan string, 1)

	go func() {
		time.Sleep(5 * time.Second)

		checkString <- "hello world!!!"
	}()

	select {
	case res := <-checkString:
		fmt.Fprintf(w, res)
	case <-ctx.Done():
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
		fmt.Println("Context cancelled")
		return
	}
}
