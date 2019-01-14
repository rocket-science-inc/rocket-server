package service

import (

)

// ApiService describes the service.
type ApiService interface {

}

type basicApiService struct{}

// NewBasicApiService returns a naive, stateless implementation of EventsService.
func NewBasicApiService() ApiService {
	return &basicApiService{}
}

// New returns a EventsService with all of the expected middleware wired in.
func New(middleware []Middleware) ApiService {
	var svc ApiService = NewBasicApiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
