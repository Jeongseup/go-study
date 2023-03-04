package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// //Define the metrics we wish to expose
// var fooMetric = prometheus.NewGauge(prometheus.GaugeOpts{
// 	Name: "foo_metric", Help: ") Shows whether a foo has occurred in our cluster"})

// var barMetric = prometheus.NewGauge(prometheus.GaugeOpts{
// 	Name: "bar_metric", Help: "Shows whether a bar has occurred in our cluster"})

// func init() {
// 	//Register metrics with prometheus
// 	prometheus.MustRegister(fooMetric)
// 	prometheus.MustRegister(barMetric)

// 	//Set fooMetric to 1
// 	fooMetric.Set(0)

// 	//Set barMetric to 0
// 	barMetric.Set(1)
// 	log.Info("there")

// }

func main() {
	fc := newFooCollector()
	reg := prometheus.NewRegistry()

	reg.MustRegister(fc)
	// prometheus.MustRegister(fc)

	log.Info("here")
	// log.WithFields(log.Fields{
	// 	"animal": "walrus",
	// }).Info("A walrus appears")

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Info("Beginnig to sereve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
