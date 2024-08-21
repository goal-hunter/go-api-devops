package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"

    "github.com/sirupsen/logrus"
)

type HistoryResponse struct {
    Domain    string   `json:"domain"`
    Result    []string `json:"result"`
    QueriedAt string   `json:"queried_at"`
}

func HistoryHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT domain, result, queried_at FROM queries ORDER BY queried_at DESC LIMIT 20")
        if err != nil {
            logrus.Errorf("Failed to fetch history: %v", err)
            http.Error(w, "Failed to fetch history", http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var history []HistoryResponse
        for rows.Next() {
            var result string
            var queriedAt string
            var domain string
            if err := rows.Scan(&domain, &result, &queriedAt); err != nil {
                logrus.Errorf("Failed to scan row: %v", err)
                continue
            }
            history = append(history, HistoryResponse{Domain: domain, Result: []string{result}, QueriedAt: queriedAt})
        }

        json.NewEncoder(w).Encode(history)
    }
}