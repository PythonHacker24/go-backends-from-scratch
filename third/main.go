package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	// "time"
)

func headerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Implement logic 

        w.Header().Set("X-Custom-Header", "Pokemon")

        // End of middleware logic
        next.ServeHTTP(w, r)
    })
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
        next.ServeHTTP(w, r)
    })
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Home Page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is the about page!")
}

func main() {
    mux := http.NewServeMux()

    mux.Handle("/", loggingMiddleware(headerMiddleware(http.HandlerFunc(homeHandler))))
    mux.Handle("/about", loggingMiddleware(headerMiddleware(http.HandlerFunc(aboutHandler))))

    log.Println("Starting server on port 8080 ...")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        // the logic that should be executed in case the listen adn serve returns error
        log.Fatal("Server Failed!", err)     
    }
}
