package handlers

import (
    "net/http"
    "encoding/json"
    "time"
)

// HealthResponse represents the response structure for the /health endpoint.
type HealthResponse struct {
    Status    string    `json:"status"`
    Timestamp time.Time `json:"timestamp"`
}

// HealthHandler handles requests to the /health endpoint.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
    // Create a response object with a "healthy" status and current timestamp
    response := HealthResponse{
        Status:    "healthy",
        Timestamp: time.Now(),
    }

    // Set the Content-Type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Encode the response as JSON and write it to the response writer
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}