// Configuration
//
//	Title: Configuration
//
//	Schemes: http
//	Version: 0.0.1
//	BasePath: /
//
//	Produces:
//	  - application/json
//
// swagger:meta
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	server, err := NewPostServer()
	if err != nil {
		log.Fatal(err)
		return
	}

	router.HandleFunc("/config/", count(server.createConfigHandler)).Methods("POST")
	router.HandleFunc("/config/{id}/{version}", count(server.getConfigHandler)).Methods("GET")
	router.HandleFunc("/config/{id}/{version}/{label}", count(server.getConfigByLabel)).Methods("GET")
	router.HandleFunc("/config/{id}/{version}", count(server.deleteConfigHandler)).Methods("DELETE")
	router.HandleFunc("/group/{id}/{version}", count(server.getGroupHandler)).Methods("GET")
	router.HandleFunc("/group/", count(server.createGroupHandler)).Methods("POST")
	router.HandleFunc("/group/", count(server.updateGroupHandler)).Methods("PUT")
	router.HandleFunc("/group/{id}/{version}", count(server.deleteGroupHandler)).Methods("DELETE")
	router.HandleFunc("/groupAll/", count(server.getDataHandler)).Methods("GET")
	router.HandleFunc("/swagger.yaml", count(server.swaggerHandler)).Methods("GET")
	router.Path("/metrics").Handler(metricsHandler())

	//show Traces UI on http://localhost:16686
	//show Metrics UI on show UI on localhost:9090

	// SwaggerUI
	optionsDevelopers := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	developerDocumentationHandler := middleware.SwaggerUI(optionsDevelopers, nil)
	router.Handle("/docs", developerDocumentationHandler)
	// start server
	srv := &http.Server{Addr: "0.0.0.0:8000", Handler: router}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}
