package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//Base API server instance description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

//API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server/configure loggers, router, database connection and etc....
func (api *API) Start() error {
	//Trying to confugre logger
	if err := api.configreLoggerField(); err != nil {
		return err
	}
	//Подтверждение того, что логгер сконфигурирован
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	//Конфигурируем маршрутизатор
	api.configreRouterField()
	//На этапе валидного завршениея стратуем http-сервер
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
