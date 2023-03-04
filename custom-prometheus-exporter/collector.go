package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type fooCollector struct {
	fooMetric *prometheus.Desc
	barMetric *prometheus.Desc
}

func newFooCollector() *fooCollector {
	return &fooCollector{
		// prometheus.GaugeOpts{
		// 	Name: "foo_metric", Help: ") Shows whether a foo has occurred in our cluster"})
		fooMetric: prometheus.NewDesc(
			"foo_metric", "HELP) Show whether a foo has occurred in our cluster",
			nil, nil,
		),
		barMetric: prometheus.NewDesc("bar_metric", "HELP) bar metric help description",
			nil, nil,
		),
	}
}

func (collector *fooCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.fooMetric
	ch <- collector.barMetric
}

func (collector *fooCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	// var metricValue float64
	// if 1 == 1 {
	// 	metricValue = 1
	// }

	now := time.Now().Unix()

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(collector.fooMetric, prometheus.CounterValue, float64(now))
	ch <- prometheus.MustNewConstMetric(collector.barMetric, prometheus.CounterValue, float64(now-100))
	// ch <- prometheus.MustNewMetricWithExemplars(co)
}
