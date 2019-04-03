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
)

type GMetInstance struct {
	metWriter    MetWriter    // metrics data writer
	metFormatter MetFormatter // metrics formatter
}

var HostAddr MetricItem
var HostName MetricItem

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

}

func CreateGMetInstance(metWriter MetWriter, metFormatter MetFormatter) GMet {
	ins := GMetInstance{metWriter, metFormatter}
	return &ins
}

// Create GMet Instance with default settings.
// (with seelog writer and ltr format
func CreateGMetInstanceByDefault(metricsFile string) GMet {
	// create a metric writer
	writer, err := CreateMetWriterBySeeLog(metricsFile)
	if err != nil {
		panic(err)
	}
	// create GMet instance by given the writer and the formatter
	gmet := CreateGMetInstance(writer, &JSON_Formatter{})
	return gmet
}

func (gmet *GMetInstance) Send(metrics ...MetricItem) error {
	if formatted, err := gmet.metFormatter.Format(metrics); err != nil {
		return err
	} else {
		gmet.metWriter.Write(formatted)
	}
	return nil
}

func (gmet *GMetInstance) Flush() {
	gmet.metWriter.Flush()
}

func (gmet *GMetInstance) Close() error {
	return gmet.metWriter.Close()
}

// Get the local IP adress
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
