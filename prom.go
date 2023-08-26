package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method"},
	)
	httpRequestsDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_requests_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestsDuration)
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	// Start timing the request
	start := time.Now()

	// Handle the request
	httpRequestsTotal.WithLabelValues(r.Method).Inc()
	// ...

	// Calculate the duration and observe it in the histogram
	duration := time.Since(start).Seconds()
	httpRequestsDuration.WithLabelValues(r.Method).Observe(duration)
}
func main1() {
	http.HandleFunc("/myendpoint", myHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
