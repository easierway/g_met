package g_met

type DummyAggregator struct {
}

func (aggregator *DummyAggregator) Aggregate(metrics []MetricItem) error{
	return nil
}

func (aggregator *DummyAggregator) GetMetrics() ([]MetricItem) {
	return nil
}

func CreateDummyAggregator() (MetAggregator, error) {
	aggregator := new(DummyAggregator)
	return aggregator, nil
}
