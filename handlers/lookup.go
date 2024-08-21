package handlers

import (
    "database/sql"
    "encoding/json"
    "net"
    "net/http"
    "time"

    "github.com/sirupsen/logrus"
)

type LookupRequest struct {
    Domain string `json:"domain"`
}

type LookupResponse struct {
    IPs []string `json:"ips"`
}

func LookupHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req LookupRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }

        ips, err := net.LookupIP(req.Domain)
        if err != nil {
            http.Error(w, "Failed to resolve domain", http.StatusInternalServerError)
            return
        }

        var ipv4s []string
        for _, ip := range ips {
            if ip.To4() != nil {
                ipv4s = append(ipv4s, ip.String())
            }
        }

        // Save query to database
        if _, err := db.Exec("INSERT INTO queries (domain, result, queried_at) VALUES ($1, $2, $3)", req.Domain, ipv4s[0], time.Now()); err != nil {
            logrus.Errorf("Failed to log query: %v", err)
        }

        response := LookupResponse{IPs: ipv4s}
        
        if err := json.NewEncoder(w).Encode(response); err != nil {
            logrus.Errorf("Failed to encode response: %v", err)
            http.Error(w, "Failed to encode response", http.StatusInternalServerError)
            return
        }
    }
}