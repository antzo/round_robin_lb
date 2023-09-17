package main

import (
	"io"
	"log"
	"net/http"
)

var serverList = []string{":8081", ":8082"}
var idx = 0

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err := http.ListenAndServe(":8080", logRequest(proxyRequest(http.DefaultServeMux))); err != nil {
		log.Fatal()
	}
}

func proxyRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create new request
		proxyURL := "http://" + serverList[idx] + r.RequestURI
		proxyReq, err := http.NewRequestWithContext(r.Context(), r.Method, proxyURL, r.Body)
		if err != nil {
			http.Error(w, "error creating proxy request", http.StatusInternalServerError)
			return
		}
		for k, values := range r.Header {
			for _, v := range values {
				proxyReq.Header.Add(k, v)
			}
		}

		resp, err := http.DefaultTransport.RoundTrip(proxyReq)
		if err != nil {
			http.Error(w, "error sending proxy request", http.StatusInternalServerError)
			return
		}
		defer func() {
			if resp.Body != nil {
				if e := resp.Body.Close(); e != nil {
					http.Error(w, "error closing proxy response body", http.StatusInternalServerError)
				}
			}
		}()

		// copy proxy response to the original response
		for k, values := range resp.Header {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)
		if _, err := io.Copy(w, resp.Body); err != nil {
			http.Error(w, "error copying proxy response body to original response", http.StatusInternalServerError)
		}

		idx++
		if idx >= len(serverList) {
			idx = 0
		}

		handler.ServeHTTP(w, r)
	})
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s\n%s %s\nHost: %s\nUser-Agent: %s\n", r.RemoteAddr, r.Method, r.URL, r.Host, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}
