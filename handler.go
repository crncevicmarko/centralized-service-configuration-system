package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mime"
	"net/http"

	cs "github.com/cvule25/airs-projekat/configstore"
	tracer "github.com/cvule25/airs-projekat/tracer"
	"github.com/gorilla/mux"
)

func (ts *Service) createConfigHandler(w http.ResponseWriter, req *http.Request) {
	httpHitsCreateConfig.Inc()
	span := tracer.StartSpanFromRequest("createConfigHandler", ts.tracer, req)
	defer span.Finish()

	span.LogFields(tracer.LogString("handler", fmt.Sprintf("handling config create at %s\n", req.URL.Path)))
	ctx := tracer.ContextWithSpan(context.Background(), span)

	idempotencyKey := req.Header.Get("Idempotency-Key")
	ok, err := ts.store.IdempotencyKeyExists(ctx, idempotencyKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if ok {
		err := errors.New("request already sent")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	} else {
		ts.store.AddIdempotencyKey(ctx, idempotencyKey)
	}

	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		tracer.LogError(span, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeConfigBody(ctx, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config, err := ts.store.Post(ctx, rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, config)
}

// swagger:route GET /config/{id}/{version}/ config getConfigById
// Get config by id
//
// responses:
//
//	404: ErrorResponse
//	200: ResponseConfig
func (ts *Service) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	httpHitsGetConfig.Inc()
	span := tracer.StartSpanFromRequest("getConfigHandler", ts.tracer, req)
	defer span.Finish()

	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get config at %s\n", req.URL.Path)),
	)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	task, err := ts.store.Get(ctx, id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, task)
}

// swagger:route GET /config/{id}/{version}/{label}/ config getConfigByLabel
// Get config by label
//
// responses:
//
//	404: ErrorResponse
//	200: ResponseConfig
func (ts *Service) getConfigByLabel(w http.ResponseWriter, req *http.Request) {
	httpHitsGetConfigByLabel.Inc()
	span := tracer.StartSpanFromRequest("getConfigByLabel", ts.tracer, req)
	defer span.Finish()

	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get config by label at %s\n", req.URL.Path)),
	)

	ctx := tracer.ContextWithSpan(context.Background(), span)

	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	label := mux.Vars(req)["label"]

	task, err := ts.store.GetPostsByLabels(ctx, id, version, label)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(ctx, w, task)
}

// swagger:route DELETE /config/{id}/ config deleteConfig
// Delete config
//
// responses:
//
//	404: ErrorResponse
//	204: NoContentResponse
func (ts *Service) deleteConfigHandler(w http.ResponseWriter, req *http.Request) {
	httpHitsDeleteConfig.Inc()
	span := tracer.StartSpanFromRequest("deleteConfigHandler", ts.tracer, req)
	defer span.Finish()

	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete config at %s\n", req.URL.Path)),
	)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	msg, err := ts.store.Delete(ctx, id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, msg)
}

// swagger:route POST /group/ group createGroup
// Create group
//
// responses:
//
//	415: ErrorResponse
//	400: ErrorResponse
//	201: ResponseGroup
func (ts *Service) createGroupHandler(w http.ResponseWriter, req *http.Request) {
	httpHitsCreateGroup.Inc()
	span := tracer.StartSpanFromRequest("createGroupHandler", ts.tracer, req)
	defer span.Finish()

	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling post group at %s\n", req.URL.Path)),
	)

	ctx := tracer.ContextWithSpan(context.Background(), span)

	idempotencyKey := req.Header.Get("Idempotency-Key")
	ok, err := ts.store.IdempotencyKeyExists(ctx, idempotencyKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if ok {
		err := errors.New("request already sent")
		http.Error(w, err.Error(), http.StatusConflict)
		return
	} else {
		ts.store.AddIdempotencyKey(ctx, idempotencyKey)
	}

	var configs []*cs.Config
	err = json.NewDecoder(req.Body).Decode(&configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := ts.store.AddConfigurationGroup(ctx, configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(task)
}

// swagger:route GET /group/{id}/ group getGroup
// Get group
//
// responses:
//
//	404: ErrorResponse
//	200: ResponseGroup
func (ts *Service) getGroupHandler(w http.ResponseWriter, req *http.Request) {
	httpHitsGetGroup.Inc()
	span := tracer.StartSpanFromRequest("getGroupHandler", ts.tracer, req)
	defer span.Finish()

	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get group at %s\n", req.URL.Path)),
	)

	ctx := tracer.ContextWithSpan(context.Background(), span)

	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]

	task, err := ts.store.GetGoupById(ctx, id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(ctx, w, task)
}

// swagger:route PUT /group/{id}/ group updateGroup -----------------------------------------------------
// Update group

// responses:

// 415: ErrorResponse
// 400: ErrorResponse
// 201: ResponseGroup
func (ts *Service) updateGroupHandler(w http.ResponseWriter, req *http.Request) {
	httpHitsUpdateGroup.Inc()
	span := tracer.StartSpanFromRequest("updateGroupHandler", ts.tracer, req)
	defer span.Finish()

	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling put group %s\n", req.URL.Path)),
	)

	ctx := tracer.ContextWithSpan(context.Background(), span)

	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeConfigBody(ctx, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := ts.store.Post(ctx, rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, post)

}

// // swagger:route DELETE /group/{id}/{version}/ group deleteGroup
// // Delete group by id
// //
// // responses:
// //
// //	404: ErrorResponse
// //	204: NoContentResponse
func (ts *Service) deleteGroupHandler(w http.ResponseWriter, req *http.Request) {
	httpHitsDeleteGroup.Inc()
	span := tracer.StartSpanFromRequest("deleteGroupHandler", ts.tracer, req)
	defer span.Finish()

	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete group %s\n", req.URL.Path)),
	)

	ctx := tracer.ContextWithSpan(context.Background(), span)

	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	task, err := ts.store.DeleteGoupById(ctx, id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, task)

}

func (ts *Service) getDataHandler(w http.ResponseWriter, req *http.Request) {
	httpHitsGetAll.Inc()
	span := tracer.StartSpanFromRequest("getDataHandler", ts.tracer, req)
	defer span.Finish()

	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling getAll groups %s\n", req.URL.Path)),
	)

	ctx := tracer.ContextWithSpan(context.Background(), span)

	allGroups, err := ts.store.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, allGroups)
}

func (ts *Service) swaggerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./swagger.yaml")
}
