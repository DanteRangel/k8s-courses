package main

import (
    "log"
    "fmt"
    "net/http"
    "os"
    "time"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
    res := fmt.Sprintf(`{"time": "%v", "host": "%v"}`, time.Now(), os.Getenv("HOSTNAME"))
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(res))
}

func main() {
    http.HandleFunc("/", ServeHTTP)
    log.Fatal(http.ListenAndServe(":8080", nil))
}