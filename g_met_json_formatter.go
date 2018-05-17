package g_met

import (
	"bytes"
	"fmt"
	"time"
)

type JSON_Formatter struct{}

func valueToJSON(v interface{}) string {
	switch v.(type) {
	case string:
		return "\"" + v.(string) + "\""
	case time.Time:
		return "\"" + (v.(time.Time)).Format(time.RFC3339Nano) + "\""
	default:
		return fmt.Sprintf("%v", v)
	}
}

func keyToJSON(k string) string {
	return "\"" + k + "\""
}

func toJSON_SEC(k string, v interface{}) string {
	return keyToJSON(k) + ":" + valueToJSON(v)
}

func (formatter *JSON_Formatter) Format(metrics []MetricItem) (string, error) {
	buf := bytes.NewBufferString("")
	buf.WriteString("{")
	buf.WriteString(toJSON_SEC(TIMESTAMP_KEY, time.Now()))
	buf.WriteString(",")
	buf.WriteString(toJSON_SEC(HOST_ADDR, HostAddr.Value))

	for _, metric := range metrics {
		buf.WriteString(",")
		buf.WriteString(toJSON_SEC(metric.Key, metric.Value))
	}
	buf.WriteString("}")
	return buf.String(), nil
}
