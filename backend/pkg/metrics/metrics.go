package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	KafkaMessagesConsumed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "kafka_messages_consumed_total",
		Help: "Total number of Kafka messages consumed",
	})
	OrderCacheHits = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "order_cache_hits_total",
		Help: "Total number of order cache hits",
	})
	OrderCacheMisses = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "order_cache_misses_total",
		Help: "Total number of order cache misses",
	})
	CacheLookupDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "order_cache_lookup_duration_seconds",
		Help:    "Duration of order cache lookup in seconds",
		Buckets: prometheus.DefBuckets,
	})
	DBLookupDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "order_db_lookup_duration_seconds",
		Help:    "Duration of order DB lookup in seconds",
		Buckets: prometheus.DefBuckets,
	})
	OrderDBLookupDuration = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "last_order_db_lookup_duration_seconds",
		Help: "Last order DB lookup duration in seconds",
	})
	OrderCacheLookupDuration = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "last_order_cache_lookup_duration_seconds",
		Help: "Last order cache lookup duration in seconds",
	})
)

func Register() {
	prometheus.MustRegister(
		KafkaMessagesConsumed,
		OrderCacheHits,
		OrderCacheMisses,
		CacheLookupDuration,
		DBLookupDuration,
		OrderDBLookupDuration,
		OrderCacheLookupDuration,
	)
	OrderCacheHits.Add(0)
	OrderCacheMisses.Add(0)
	KafkaMessagesConsumed.Add(0)
}
