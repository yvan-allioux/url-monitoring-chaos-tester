# url monitoring chaos tester
 tester for url monitoring
 
say "down" for 30 minutes ad "up" for 30 minutes

0 -> 30 minutes

    json with info up inside

0 -> 45 minutes

    json with info down inside

45 -> 60 

    error 500


url-monitoring-chaos-tester

go image 331M

php image 441M


```go
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
```

