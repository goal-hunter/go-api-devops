package main

import (
    "encoding/json"
    "net/http"
    "time"
    "os"

    "github.com/M-Killer-dev/devops-assignment/internal/db"
    "github.com/M-Killer-dev/devops-assignment/handlers"
    "github.com/M-Killer-dev/devops-assignment/internal/metrics"
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
)

type RootResponse struct {
    Version    string `json:"version"`
    Date       int64  `json:"date"`
    Kubernetes bool   `json:"kubernetes"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    response := RootResponse{
        Version:    "0.1.0",
        Date:       time.Now().Unix(),
        Kubernetes: isKubernetes(),
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        logrus.Errorf("Failed to encode response: %v", err)
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}

func isKubernetes() bool {
    // Check if running in Kubernetes by looking for the presence of an environment variable
    return os.Getenv("KUBERNETES_SERVICE_HOST") != ""
}

func main() {
    // Set log output to a file
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        logrus.Fatalf("error opening log file: %v", err)
    }
    logrus.SetOutput(file)

    // Set log format to JSON (optional)
    logrus.SetFormatter(&logrus.JSONFormatter{})

    
    // Connect to the database
    dbConn, err := db.Connect()
    if err != nil {
        logrus.Errorf("could not connect to the database: %v", err)
    }

    // Run database migrations
    if err := db.Migrate(dbConn); err != nil {
        logrus.Errorf("could not run database migrations: %v", err)
    }

    r := mux.NewRouter()

    // Register routes
    r.HandleFunc("/", rootHandler).Methods("GET")
    r.HandleFunc("/v1/tools/lookup", handlers.LookupHandler(dbConn)).Methods("POST")
    r.HandleFunc("/v1/tools/validate", handlers.ValidateHandler()).Methods("POST")
    r.HandleFunc("/v1/history", handlers.HistoryHandler(dbConn)).Methods("GET")
    r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
    r.Handle("/metrics", metrics.MetricsHandler())

    logrus.Info("Starting server on :3000")
    if err := http.ListenAndServe(":3000", r); err != nil {
        logrus.Errorf("could not start server: %v", err)
    }
}