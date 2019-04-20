package g_met

type AdnetAggregator struct {
	data map[string]interface{}
}

func (aggregator *AdnetAggregator) Aggregate(metrics []MetricItem) error{
	// TODO: add lock, and custom adn itself
	for _, metric := range metrics {
		if metric.Key == "requestid" {
			if value, exist := aggregator.data["request_num"]; exist {
				aggregator.data["request_num"] = value.(int) + 1
			} else {
				aggregator.data["request_num"] = 1
			}
		}

		if metric.Key == "input_bytes" {
			if value, exist := aggregator.data["input_bytes"]; exist {
				aggregator.data["input_bytes"] = value.(int) + metric.Value.(int)
			} else {
				aggregator.data["input_bytes"] = metric.Value.(int)
			}
		}
	}
	return nil
}

func (aggregator *AdnetAggregator) GetMetrics() ([]MetricItem) {
	metrics := make([]MetricItem, 0, len(aggregator.data))
	for key, value := range aggregator.data {
		metrics = append(metrics, Metric(key, value))
	}
	return metrics
}

func CreateAdnetAggregator() (MetAggregator) {
	aggregator := new(AdnetAggregator)
	aggregator.data = make(map[string]interface{})
	return aggregator
}
