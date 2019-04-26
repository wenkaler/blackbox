package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type handlerConfig struct {
	logger log.Logger
	router *mux.Router
	svc    service
}

func newHandler(cfg *handlerConfig) {

	projectRouter := cfg.router.PathPrefix("/project").Subrouter()

	createProject := makeCreateProjectEndpoint(cfg.svc)
	projectRouter.Path("/create").Methods(http.MethodPost).Handler(kithttp.NewServer(
		createProject,
		decodeByteRequest,
		encodeResponse,
	))

	updateProject := makeUpdateProjectEndpoint(cfg.svc)
	projectRouter.Path("/update").Methods(http.MethodPost).Handler(kithttp.NewServer(
		updateProject,
		decodeByteRequest,
		encodeResponse,
	))

	removeProject := makeRemoveProjectEndpoint(cfg.svc)
	projectRouter.Path("/remove").Methods(http.MethodGet).Handler(kithttp.NewServer(
		removeProject,
		decodeToken,
		encodeResponse,
	))

	listProject := makeListProjectEndpoint(cfg.svc)
	projectRouter.Path("/list").Methods(http.MethodGet).Handler(kithttp.NewServer(
		listProject,
		decodeToken,
		encodeResponse,
	))

	getProject := makeGetProjectEndpoint(cfg.svc)
	projectRouter.Path("/get").Methods(http.MethodGet).Handler(kithttp.NewServer(
		getProject,
		decodeToken,
		encodeResponse,
	))

	settingRouter := cfg.router.PathPrefix("/setting").Subrouter()

	updateSetting := makeUpdateSettingEndpoint(cfg.svc)
	settingRouter.Path("/update").Methods(http.MethodPost).Handler(kithttp.NewServer(
		updateSetting,
		decodeByteRequest,
		encodeResponse,
	))

	getSetting := makeGetSettingEndpoint(cfg.svc)
	settingRouter.Path("/get").Methods(http.MethodGet).Handler(kithttp.NewServer(
		getSetting,
		decodeToken,
		encodeResponse,
	))

	confirmSetting := makeConfirmSettingEndpoint(cfg.svc)
	settingRouter.Path("/confirm").Methods(http.MethodGet).Handler(kithttp.NewServer(
		confirmSetting,
		decodeToken,
		encodeResponse,
	))

	initSetting := makeInitialSettingEndpoint(cfg.svc)
	settingRouter.Path("/init").Methods(http.MethodGet).Handler(kithttp.NewServer(
		initSetting,
		decodeTokenVarious,
		encodeResponse,
	))

	cleanSettings := makeCleanUnusedSettingsEndpoint(cfg.svc)
	settingRouter.Path("/clean").Methods(http.MethodGet).Handler(kithttp.NewServer(
		cleanSettings,
		decodeTokenVarious,
		encodeResponse,
	))

}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if response != nil {
		w.Header().Set("Content-Type", "application/json; charset=uft-8")
		_, err := w.Write(response.([]byte))
		if err != nil {
			return fmt.Errorf("encodeResponse: %v", err)
		}
	}
	return nil
}

func decodeError(r *http.Response) error {
	defaultErr := &Error{Code: r.StatusCode, Level: "api", Message: http.StatusText(r.StatusCode)}

	e := &Error{}
	if err := json.NewDecoder(r.Body).Decode(e); err == nil {
		defaultErr = e
	}
	return defaultErr
}
