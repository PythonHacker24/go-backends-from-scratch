package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"
)

type User struct {
    ID      int     `json:"id"`
    Name    string  `json:"name"`
    Email   string  `json:"email"`
}

var (
    users = make(map[int]User)
    idSeq = 1
    mutex = &sync.Mutex{}
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
        case "GET":
            mutex.Lock()
            defer mutex.Unlock()

            usersList := make([]User, 0, len(users))
            for _, user := range users {
                usersList = append(usersList, user)
            }

            json.NewEncoder(w).Encode(usersList)

        case "POST":
            var user User
            if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
                http.Error(w, "Invalid JSON", http.StatusBadRequest)
                return
            }

            mutex.Lock()
            user.ID = idSeq
            idSeq++
            users[user.ID] = user 
            mutex.Unlock()

            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(user)

        default:
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var id int 
    _, err := fmt.Sscanf(r.URL.Path, "/users/%d", &id)
    if err != nil {
        http.Error(w, "Invalid User ID", http.StatusBadRequest)
        return
    }

    mutex.Lock()
    defer mutex.Unlock()

    user, exists := users[id]
    if !exists {
        http.Error(w, "User Not Found", http.StatusNotFound)
        return
    }

    switch r.Method {
        case "GET":
            json.NewEncoder(w).Encode(user)

        case "PUT":
            var updatedUser User
            if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
                http.Error(w, "Invalid JSON", http.StatusBadRequest)
                return 
            }

            updatedUser.ID = id 
            users[id] = updatedUser
            json.NewEncoder(w).Encode(updatedUser)

        case "DELETE":
            delete(users, id)
            w.WriteHeader(http.StatusNoContent)

        default: 
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/users/", userHandler)

    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
