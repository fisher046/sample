package main

import (
	"os"
)

const (
	envAPIPort     = "API_PORT"
	envMetricsPort = "METRICS_PORT"

	defaultAPIPort     = "8000"
	defaultMetricsPort = "8001"
)

func main() {
	apiPort := os.Getenv(envAPIPort)
	if apiPort == "" {
		apiPort = defaultAPIPort
	}

	metricsPort := os.Getenv(envMetricsPort)
	if metricsPort == "" {
		metricsPort = defaultMetricsPort
	}
}
