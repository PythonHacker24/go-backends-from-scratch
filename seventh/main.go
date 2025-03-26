package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type User struct {
    Name    string  `json:"name"`
    Email   string  `json:"email"`
    Age     int     `json:"age"`
}

func (user User) ValidateUser() error {
    /* if user.Name == "" {
        return fmt.Errorf("missing field: name")
    }
    */

    if user.Email == "" {
        return fmt.Errorf("missing field: email")
    }

    if user.Age <= 0 {
        return fmt.Errorf("invalid age")
    }

    return nil
}

func (user *User) Normalize() {
    if user.Name == "" {
        user.Name = "Unknown"
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    var user User
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&user)
    if err != nil {
        http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
        return
    }

    user.Normalize()
    err = user.ValidateUser() 
    if err != nil {
        response := map[string]string{"error": err.Error()}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, "Failed to encode json response", http.StatusInternalServerError)
        } 
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User data is valid"))
}

/* 
    {
        name: <string>,
        email: <string>,
        age: <int>
    }

    {
        "error": "invalid age"
    }
*/
