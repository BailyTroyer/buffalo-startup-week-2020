package handler

import (
	"fmt"
	"net/http"

	"bflobox-api/internal/util"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

// API Wrapper struct with router & config
type API struct {
	Config util.Config
	Router *mux.Router
}

// Initialize Mux, Middleware and routes
func (a *API) Initialize() {

	// init logger
	util.InitLogger()

	// Mux Router
	a.Router = mux.NewRouter()

	// Prometheus middleware
	a.Router.Use(util.PrometheusMiddleware)
	a.Router.Path("/metrics").Handler(promhttp.Handler())

	// standard logging middleware
	a.Router.Use(util.LoggingMiddleware)

	// API endpoints/routes
	a.get("/health", a.Health)
}

// Run HTTP server
func (a *API) Run() {
	zap.L().Info("listening")
	zap.L().Fatal("", zap.Error(http.ListenAndServe(fmt.Sprintf(":%v", a.Config.Meta.Port), cors.Default().Handler(a.Router))))
}

// Get wrap the router for GET method
func (a *API) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
