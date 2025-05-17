package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/config"
	"github.com/Wafer233/msproject-be/api-gateway/internal/metrics"
)

func ProvideMetricsCollector(cfg *config.Config) *metrics.MetricsCollector {
	return metrics.NewMetricsCollector(&cfg.Metrics)
}
