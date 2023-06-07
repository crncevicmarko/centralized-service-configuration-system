package main

import (
	"context"
	"encoding/json"
	cs "github.com/cvule25/airs-projekat/configstore"
	tracer "github.com/cvule25/airs-projekat/tracer"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func decodeConfigBody(ctx context.Context, r io.Reader) (*cs.Config, error) {
	span := tracer.StartSpanFromContext(ctx, "decodeBody")
	defer span.Finish()

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt cs.Config
	if err := dec.Decode(&rt); err != nil {
		tracer.LogError(span, err)
		return nil, err
	}
	return &rt, nil
}

func renderJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	span := tracer.StartSpanFromContext(ctx, "renderJSON")
	defer span.Finish()

	js, err := json.Marshal(v)
	if err != nil {
		tracer.LogError(span, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createId() string {
	return uuid.New().String()
}
