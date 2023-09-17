package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int
var name string

func main() {
	flag.IntVar(&port, "port", 8081, "http port to listen")
	flag.StringVar(&name, "name", "test", "name of the server")
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("starting server \"%s\" at %d\n", name, port)

	http.Handle("/", logRequest(helloHandler()))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(
			fmt.Sprintf("Hello from BackendServer %s running at %d", name, port),
		)); err != nil {
			w.WriteHeader(500)
		}
	})
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"Received request from %s\n%s %s\nHost: %s\nUser-Agent: %s\n",
			r.RemoteAddr, r.Method, r.URL, r.Host, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}
