package main

import (
    "fmt"
    "net/http"
    "strings"
)

// https://api.example.com/api/v1/greet?name=Aditya
// https://api.example.com/api/v1/greet

// Understanding Query Parameters
func greetHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    name := query.Get("name")
    if name == "" {
        name = "Guest" 
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}

// Extracting Path Variables

// https://api.example.com/user/123
// 1 -> User 
// 2 -> UserID
func userHandler(w http.ResponseWriter, r *http.Request) {
    pathSegments := strings.Split(r.URL.Path, "/")
    if len(pathSegments) >= 3 && pathSegments[1] == "user" {
        userID := pathSegments[2]
        fmt.Fprintf(w, "User ID: %s", userID)
    } else {
        http.NotFound(w, r)
    }
}

// Combining Both => Handling Query Parameters + Extracting Path Variables
// https://api.example.com/username/123?includeDetails=true
func userDetailsHandler(w http.ResponseWriter, r *http.Request) {
    pathSegments := strings.Split(r.URL.Path, "/")
    query := r.URL.Query()
    includeDetails := query.Get("includeDetails")

    if len(pathSegments) >= 3 && pathSegments[1] == "username" {
        userID := pathSegments[2]
        response := fmt.Sprintf("User ID: %s", userID)
        if includeDetails == "true" {
            response += " (Details included)"
        }
        fmt.Fprintln(w, response)
    } else {
        http.NotFound(w, r)
    }
}

func main() {
    http.HandleFunc("/greet", greetHandler)
    http.HandleFunc("/user/", userHandler)
    http.HandleFunc("/username/", userDetailsHandler)

    fmt.Println("Listening at port 8080 ...")
    // We are using default mux for demonstration purposes
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to listen at port 8080", err)
    }
}

// https://api.example.com/api/v1/notion?sessionToken="823492384"&referralToken="user123" -> User123 was the one who refered -> referralToken -> +50 Points
// https://api.example.com/api/v1/notion?sessionToken="823492384" -> User came here and signed up by his own.

// https://api.example.com/api/v1/status/order/1234
