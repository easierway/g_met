package g_met


type AdnetAggregator struct {
	metrics map[string]interface{}
}

func (aggregator *AdnetAggregator) Aggregate(metric []MetricItem) error{
	return nil
}

func (aggregator *AdnetAggregator) GetMetrics() ([]MetricItem) {
	var metrics []MetricItem
	return metrics
}

func CreateAdnetAggregator() (MetAggregator, error) {
	var err error
	aggregator := new(AdnetAggregator)
	if err != nil {
		return nil, err
	}

	return aggregator, nil
}
