package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	cm "pnp-master/Framework/git/order/common"
	"pnp-master/Framework/git/order/middleware"
	"pnp-master/Framework/git/order/services"
	"pnp-master/Framework/git/order/transport"

	log "github.com/Sirupsen/logrus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func initHandlers() {
	var svc services.PaymentServices
	svc = services.PaymentService{}
	svc = middleware.BasicMiddleware()(svc)

	root := cm.Config.RootURL
	fmt.Println(cm.Config.RootURL)
	http.Handle(fmt.Sprintf("%s/orders", root), httptransport.NewServer(
		transport.OrderEndpoint(svc), transport.DecodeRequest, transport.EncodeResponse,
	))

	// menambahkan request untuk customer dan product
	rootCustomer := cm.Config.RootURLCustomer
	http.Handle(fmt.Sprintf("%s/customers", rootCustomer), httptransport.NewServer(
		transport.CustomerEndpoint(svc), transport.DecodeCustomerRequest, transport.EncodeResponse,
	))

	rootProduct := cm.Config.RootURLProduct
	http.Handle(fmt.Sprintf("%s/products", rootProduct), httptransport.NewServer(
		transport.ProductEndpoint(svc), transport.DecodeProductRequest, transport.EncodeResponse,
	))

	rootData := cm.Config.RootURLData
	http.Handle(fmt.Sprintf("%s/fastpay", rootData), httptransport.NewServer(
		transport.FastpayEndpoint(svc), transport.DecodeFastpayRequest, transport.EncodeResponse,
	))

	http.Handle(fmt.Sprintf("%s/trips", root), httptransport.NewServer(
		transport.TripsEndpoint(svc), transport.DecodeTripRequest, transport.EncodeResponse,
	))
}

var logger *log.Entry

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})
	//log.SetReportCaller(true)
}

func main() {
	configFile := flag.String("conf", "conf-dev.yml", "main configuration file")
	flag.Parse()
	initLogger()
	log.WithField("file", *configFile).Info("Loading configuration file")
	cm.LoadConfigFromFile(configFile)
	initHandlers()

	var err error
	if cm.Config.RootURL != "" || cm.Config.ListenPort != "" {
		err = http.ListenAndServe(cm.Config.ListenPort, nil)
	}

	if err != nil {
		log.WithField("error", err).Error("Unable to start the server")
		os.Exit(1)
	}
}
