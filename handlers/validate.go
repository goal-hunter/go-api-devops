package handlers

import (
    "encoding/json"
    "net"
    "net/http"
    "github.com/sirupsen/logrus"
)

type ValidateRequest struct {
    IP string `json:"ip"`
}

type ValidateResponse struct {
    Valid bool `json:"valid"`
}

func ValidateHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req ValidateRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }

        ip := net.ParseIP(req.IP)
        isValid := ip != nil && ip.To4() != nil

        response := ValidateResponse{Valid: isValid}

        w.Header().Set("Content-Type", "application/json")

        if err := json.NewEncoder(w).Encode(response); err != nil {
            logrus.Errorf("Failed to encode response: %v", err)
            http.Error(w, "Failed to process response", http.StatusInternalServerError)
        }
    }
}