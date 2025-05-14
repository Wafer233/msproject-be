package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
	"time"
)

// GORMMetrics collects metrics about GORM database operations
type GORMMetrics struct {
	registry             *prometheus.Registry
	dbOperationsDuration *prometheus.HistogramVec
	dbOperationsTotal    *prometheus.CounterVec
	dbErrorsTotal        *prometheus.CounterVec
}

// NewGORMMetrics creates a new GORM metrics collector
func NewGORMMetrics(namespace string) *GORMMetrics {
	registry := prometheus.NewRegistry()

	metrics := &GORMMetrics{
		registry: registry,
		dbOperationsDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:      "gorm_operation_duration_seconds",
				Help:      "Duration of GORM operations in seconds",
				Buckets:   prometheus.ExponentialBuckets(0.001, 2, 10), // From 1ms to ~1s
			},
			[]string{"operation", "entity", "status"},
		),
		dbOperationsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "gorm_operations_total",
				Help:      "Total number of GORM operations",
			},
			[]string{"operation", "entity"},
		),
		dbErrorsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "gorm_errors_total",
				Help:      "Total number of GORM errors",
			},
			[]string{"operation", "entity", "error_type"},
		),
	}

	registry.MustRegister(
		metrics.dbOperationsDuration,
		metrics.dbOperationsTotal,
		metrics.dbErrorsTotal,
	)

	return metrics
}

// GetRegistry returns the metrics registry
func (gm *GORMMetrics) GetRegistry() *prometheus.Registry {
	return gm.registry
}

// NewCallback creates a GORM callback for collecting metrics
func (gm *GORMMetrics) NewCallback() func(db *gorm.DB) {
	return func(db *gorm.DB) {
		startTime := time.Now()

		// Add callback for GORM's operations
		db.Callback().Create().Before("gorm:create").Register("metrics:before_create", func(db *gorm.DB) {
			gm.dbOperationsTotal.WithLabelValues("create", db.Statement.Table).Inc()
		})

		db.Callback().Create().After("gorm:create").Register("metrics:after_create", func(db *gorm.DB) {
			status := "success"
			if db.Error != nil {
				status = "error"
				gm.dbErrorsTotal.WithLabelValues("create", db.Statement.Table, db.Error.Error()).Inc()
			}
			gm.dbOperationsDuration.WithLabelValues("create", db.Statement.Table, status).
				Observe(time.Since(startTime).Seconds())
		})

		// Query callbacks
		db.Callback().Query().Before("gorm:query").Register("metrics:before_query", func(db *gorm.DB) {
			gm.dbOperationsTotal.WithLabelValues("query", db.Statement.Table).Inc()
		})

		db.Callback().Query().After("gorm:query").Register("metrics:after_query", func(db *gorm.DB) {
			status := "success"
			if db.Error != nil {
				status = "error"
				gm.dbErrorsTotal.WithLabelValues("query", db.Statement.Table, db.Error.Error()).Inc()
			}
			gm.dbOperationsDuration.WithLabelValues("query", db.Statement.Table, status).
				Observe(time.Since(startTime).Seconds())
		})

		// Update callbacks
		db.Callback().Update().Before("gorm:update").Register("metrics:before_update", func(db *gorm.DB) {
			gm.dbOperationsTotal.WithLabelValues("update", db.Statement.Table).Inc()
		})

		db.Callback().Update().After("gorm:update").Register("metrics:after_update", func(db *gorm.DB) {
			status := "success"
			if db.Error != nil {
				status = "error"
				gm.dbErrorsTotal.WithLabelValues("update", db.Statement.Table, db.Error.Error()).Inc()
			}
			gm.dbOperationsDuration.WithLabelValues("update", db.Statement.Table, status).
				Observe(time.Since(startTime).Seconds())
		})

		// Delete callbacks
		db.Callback().Delete().Before("gorm:delete").Register("metrics:before_delete", func(db *gorm.DB) {
			gm.dbOperationsTotal.WithLabelValues("delete", db.Statement.Table).Inc()
		})

		db.Callback().Delete().After("gorm:delete").Register("metrics:after_delete", func(db *gorm.DB) {
			status := "success"
			if db.Error != nil {
				status = "error"
				gm.dbErrorsTotal.WithLabelValues("delete", db.Statement.Table, db.Error.Error()).Inc()
			}
			gm.dbOperationsDuration.WithLabelValues("delete", db.Statement.Table, status).
				Observe(time.Since(startTime).Seconds())
		})
	}
}
