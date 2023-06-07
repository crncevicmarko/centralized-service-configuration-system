package main

import (
	cs "github.com/cvule25/airs-projekat/configstore"
	tracer "github.com/cvule25/airs-projekat/tracer"
	opentracing "github.com/opentracing/opentracing-go"
	"io"
)

const (
	name = "config_service"
)

type Service struct {
	store  *cs.ConfigStore
	tracer opentracing.Tracer
	closer io.Closer
}

func NewPostServer() (*Service, error) {
	store, err := cs.New()
	if err != nil {
		return nil, err
	}

	tracer, closer := tracer.Init(name)
	opentracing.SetGlobalTracer(tracer)
	return &Service{
		store:  store,
		tracer: tracer,
		closer: closer,
	}, nil
}

func (ts *Service) GetTracer() opentracing.Tracer {
	return ts.tracer
}

func (ts *Service) GetCloser() io.Closer {
	return ts.closer
}

func (ts *Service) CloseTracer() error {
	return ts.closer.Close()
}
