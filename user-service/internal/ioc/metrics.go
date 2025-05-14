package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/metrics"
)

func ProvideGORMMetrics(cfg *config.Config) *metrics.GORMMetrics {
	return metrics.NewGORMMetrics(cfg.Metrics.Namespace)
}
