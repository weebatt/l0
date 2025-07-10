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
)

func Register() {
	prometheus.MustRegister(KafkaMessagesConsumed, OrderCacheHits, OrderCacheMisses)
}
