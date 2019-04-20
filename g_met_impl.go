// GMet is golang client of XMet API.
// For more, see also: https://github.com/chaocai2001/g_met
// Created on 2018.5
package g_met

import (
	"net"
	"os"
)

const (
	HOST_ADDR     = "host"
	HOST_NAME     = "hostname"
	MISSING_VALUE = "N/A"
	SYSTYPE       = "systype"
)

type GMetInstance struct {
	metAggregator MetAggregator // metrics aggregator
	metFormatter  MetFormatter  // metrics formatter
	metWriter     MetWriter     // metrics data writer
}

var HostAddr MetricItem
var HostName MetricItem
var SysType MetricItem

func init() {
	var err error
	HostAddr, err = IpAddress()
	if err != nil {
		HostAddr = Metric(HOST_ADDR, err.Error())
	}
	hostname, err := os.Hostname()
	if err != nil {
		HostName = Metric(HOST_NAME, err.Error())
	} else {
		HostName = Metric(HOST_NAME, hostname)
	}
	SysType = Metric(SYSTYPE, MISSING_VALUE)
}

func CreateGMetInstance(metAggregator MetAggregator, metFormatter MetFormatter, metWriter MetWriter) GMet {
	ins := GMetInstance{metAggregator, metFormatter, metWriter}
	return &ins
}

// Create GMet Instance with default settings.
// with seelog writer and json format
func CreateGMetInstanceByDefault(metricsFile string, sysType string) GMet {
	// create a metric dummy aggregator default
	aggregator, err := CreateDummyAggregator()
	// create a metric writer
	writer, err := CreateMetWriterBySeeLog(metricsFile)
	if err != nil {
		panic(err)
	}
	// create GMet instance by given the writer and the formatter
	gmet := CreateGMetInstance(aggregator, &JSON_Formatter{}, writer)
	// set the systype
	SysType.Value = sysType
	return gmet
}

func (gmet *GMetInstance) Send(metrics ...MetricItem) error {
	// aggregate
	// TODO: add error check
	gmet.metAggregator.Aggregate(metrics)
	// format
	if formatted, err := gmet.metFormatter.Format(metrics); err != nil {
		return err
	} else {
		gmet.metWriter.Write(formatted)
	}
	return nil
}

func (gmet *GMetInstance) Flush() {
	// TODO: repeated logic
	metrics := gmet.metAggregator.GetMetrics()
	if formatted, err := gmet.metFormatter.Format(metrics); err != nil {
		return
	} else {
		gmet.metWriter.Write(formatted)
	}
	gmet.metWriter.Flush()
}

func (gmet *GMetInstance) Close() error {
	return gmet.metWriter.Close()
}

func (gmet *GMetInstance) WithAggregator(aggregator MetAggregator) GMet {
	gmet.metAggregator = aggregator
	return gmet
}

// Get the local IP address
func IpAddress() (MetricItem, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return MetricItem{HOST_ADDR, MISSING_VALUE}, err
	}
	for _, address := range addrs {
		// Check if it is ip circle
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return MetricItem{HOST_ADDR, ipnet.IP.String()}, nil
			}

		}
	}
	return MetricItem{HOST_ADDR, MISSING_VALUE}, err
}
