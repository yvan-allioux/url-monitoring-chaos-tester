package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"
)

type Response struct {
    Type    string `json:"Type"`
    Content string `json:"Content"`
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        currentTime := time.Now()
        minute := currentTime.Minute()

        var response Response
        if minute >= 45 {
            // Renvoyer une erreur 500
            http.Error(w, "500", http.StatusInternalServerError)
            return
        }
        if minute < 30 {
            response = Response{
                Type:    "success",
                Content: "success we are in the 30 first minutes of the hour.",
            }
        } else {
            response = Response{
                Type:    "error",
                Content: "error we are NOT in the 30 first minutes of the hour.",
            }
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    })

    log.Println("Server starting on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
