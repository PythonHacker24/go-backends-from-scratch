package main

import (
    "log"
    "net/http"
)

func main() {
    
    fs := http.FileServer(http.Dir("./static"))
    
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html")
    })

    log.Println("Serving on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
