package middleware

import (
	"context"
	cm "pnp-master/Framework/git/order/common"
	"pnp-master/Framework/git/order/services"
	"time"

	log "github.com/Sirupsen/logrus"
)

func BasicMiddleware() services.ServiceMiddleware {
	return func(next services.PaymentServices) services.PaymentServices {
		return BasicMiddlewareStruct{next}
	}
}

type BasicMiddlewareStruct struct {
	services.PaymentServices
}

func (mw BasicMiddlewareStruct) OrderHandler(ctx context.Context, request cm.Message) cm.Message {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("OrderHandler ends")
	}(time.Now())
	log.WithField("request", request).Info("OrderHandler begins")
	return mw.PaymentServices.OrderHandler(ctx, request)
}

// add middleware CustomerHandler and ProductHandler
func (mw BasicMiddlewareStruct) CustomerHandler(ctx context.Context, request cm.Customer) cm.Customer {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("CustomerHandler ends")
	}(time.Now())
	log.WithField("request", request).Info("CustomerHandler begins")
	return mw.PaymentServices.CustomerHandler(ctx, request)
}

func (mw BasicMiddlewareStruct) ProductHandler(ctx context.Context, request cm.Product) cm.Product {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("ProductHandler ends")
	}(time.Now())
	log.WithField("request", request).Info("ProductHandler begins")
	return mw.PaymentServices.ProductHandler(ctx, request)
}

func (mw BasicMiddlewareStruct) FastPayHandler(ctx context.Context, request cm.FastPayRequest) cm.FastPayResponse {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("FastPayHandler ends")
	}(time.Now())
	log.WithField("request", request).Info("FastPayHandler begins")
	return mw.PaymentServices.FastPayHandler(ctx, request)
}

func (mw BasicMiddlewareStruct) TripsHandler(ctx context.Context, request cm.MyTrips) cm.MytripsResponse {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("TripsHandler ends")
	}(time.Now())
	log.WithField("request", request).Info("TripsHandler begins")
	return mw.PaymentServices.TripsHandler(ctx, request)
}
