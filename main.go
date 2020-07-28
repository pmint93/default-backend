package main

import (
    "fmt"
    "os"
    "net/http"
    "strconv"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {
    response_code_str := getEnv("RESPONSE_CODE", "404")
    response_code, err := strconv.Atoi(response_code_str)
    if err != nil {
        fmt.Println("Invalid response code")
    }
    response_text := getEnv("RESPONSE_TEXT", "default backend - " + response_code_str)
    w.WriteHeader(response_code)
    fmt.Fprintf(w, response_text)
}

func main() {
    http.HandleFunc("/", SimpleServer)
    http.ListenAndServe(":8080", nil)
}
