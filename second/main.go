package main

import (
    "fmt"
    "net/http"
)

type HomeHandler struct{} 

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to Go Server!")
}

func main() {
    mux := http.NewServeMux()

    mux.Handle("/", HomeHandler{})

    mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World! You are accessing %s and using User Agent %s\n", r.URL.Path, r.Header.Get("User-Agent"))
    })

    mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, `{"status": "ok"}`)
    })

    fmt.Println("Server starting on port 8080 ...")
    http.ListenAndServe(":8080", mux)
}

// /pokemon -> endpoint which I am desiring 

// POST /pokemon -> create a pokemon -> create_pokemon() -> handling everything related to creating a pokemon
// GET /pokemon -> gets all the pokemon -> get_pokemon() 
// DELETE /pokemon -> deletes the pokemon -> delete_pokemon() 

// /pokemon -> how will the backend know where to route
// ServeMux -> Multiplexer 

// GO ahead and study about how a request and response looks like 


// user -> localhost:8080/hello -> mux -> /hello -> line 25 -> fmt -> "Hello World" -> reponded to the user back
