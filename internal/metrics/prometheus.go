package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

var (
    RequestCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "api_requests_total",
            Help: "Total number of requests to the API",
        },
        []string{"method", "endpoint"},
    )
)

func Init() {
    prometheus.MustRegister(RequestCounter)
}

func MetricsHandler() http.Handler {
    return promhttp.Handler()
}