package prom

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Ip_addr = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ip_addr",
		Help: "ip addr of host on communications",
		//ConstLabels: prometheus.Labels{"ip": ip[line]},
	}, []string{
		"ip", "direction",
	})
)

func Init() {
	prometheus.MustRegister(Ip_addr)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func Display_metrics(dir map[int]string, ip map[int]string, b10s map[int]float64, ip_addr *prometheus.GaugeVec) {
	for line, value := range dir {
		//sline := strconv.Itoa(line)
		if value == "src" {
			ip_addr.WithLabelValues(ip[line], "source").Set(b10s[line])
			println(ip[line], b10s[line])
		} else {
			ip_addr.WithLabelValues(ip[line], "dest").Set(b10s[line])
			println(ip[line], b10s[line])
		}

	}
}
