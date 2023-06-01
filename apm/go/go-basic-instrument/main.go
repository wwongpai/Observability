package main

import (
	"fmt"
	"log"
	"net/http"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	tracer.Start(
		tracer.WithService("go-newbie"),
		tracer.WithEnv("dev"),
	)
	defer tracer.Stop()

	// Create a traced mux router
	mux := httptrace.NewServeMux()

	// Continue using the router as you normally would.
	mux.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
