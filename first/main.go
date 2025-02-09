package main

import (
    "fmt"
    "net/http"
)

// *http.Request -> location -> User requests and parameters are present -> user provided data
// http.ResponseWriter -> Backend writes it's response
func apiHandler(w http.ResponseWriter, r *http.Request) {
    // "Hello World -> w"
    fmt.Fprintln(w, "Hello World!")
}

func main() {
    // localhost:8080/api -> called -> hanlder -> function
    http.HandleFunc("/api", apiHandler)
    http.HandleFunc("/api/user", apiHandler)

    fmt.Println("Starting server at port 8080 ...")
    
    // host:port -> localhost:8080 | :8080 -> Listen and Server at all interfaces on port 8080  
    http.ListenAndServe(":8080", nil)
}
