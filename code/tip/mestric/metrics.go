package metrics

type (
	Sink interface {
		Put(k string, v interface{}) error
	}
	Collector struct {
	}
)

// ReportMetrics writes the collected metrics to a KV store instance.
func (c *Collector) ReportMetrics(s Sink) error {
	// for each metric call s.Put(k, v)
	return nil
}

// Observe records a value for a particular metric.
func (c *Collector) Observe(metric string, value interface{}) {
	// ...
}
