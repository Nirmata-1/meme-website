package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"
)

type MemeResponse struct {
    PostLink  string `json:"postLink"`
    Title     string `json:"title"`
    URL       string `json:"url"`
    Subreddit string `json:"subreddit"`
}

var baseURL = "https://meme-api.com/gimme"

func initBaseURL() {
    wholesome := os.Getenv("WHOLESOME")
    switch wholesome {
    case "true":
        baseURL += "/wholesome"
    case "", "false":
    // Default behavior
    default:
        log.Printf("Warning: Invalid value for WHOLESOME: %s. Using default behavior.", wholesome)
    }
}

func getMeme(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get(baseURL)
    if err != nil {
        http.Error(w, "Failed to get meme", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        http.Error(w, "Failed to get meme", http.StatusInternalServerError)
        return
    }

    var meme MemeResponse
    if err := json.NewDecoder(resp.Body).Decode(&meme); err != nil {
        http.Error(w, "Failed to decode meme response", http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Failed to load template", http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, meme); err != nil {
        http.Error(w, "Failed to render template", http.StatusInternalServerError)
    }
}

func main() {
    initBaseURL()
    http.HandleFunc("/", getMeme)

    port := os.Getenv("PORT")
    if port == "" {
        port = "80"
    }

    fmt.Printf("Server starting on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}