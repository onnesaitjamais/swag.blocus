/*
#######
##                                   __   __
##        ____    _____ ____ _      / /  / /__  ______ _____
##       (_-< |/|/ / _ `/ _ `/ _   / _ \/ / _ \/ __/ // (_-<
##      /___/__,__/\_,_/\_, / (_) /_.__/_/\___/\__/\_,_/___/
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package blocus

import (
	"github.com/arnumina/swag"
	"github.com/arnumina/swag/service"
	"github.com/arnumina/swag/util/options"
	"github.com/gorilla/mux"

	v0 "github.com/arnumina/swag.blocus/internal/api/v0"
	v1 "github.com/arnumina/swag.blocus/internal/api/v1"
)

const _port = 65534

func initialize(r *mux.Router, s *service.Service) error {
	v0.Routes(r.PathPrefix("/api/v0").Subrouter())
	v1.Routes(r.PathPrefix("/api/v1").Subrouter(), s)

	return nil
}

// Run AFAIRE
func Run(version, builtAt string) error {
	router := mux.NewRouter()

	service, err := swag.NewService(
		"blocus",
		version,
		builtAt,
		swag.Config(
			"default",
			options.Options{
				"port": _port,
			},
		),
		swag.Server(
			"http",
			options.Options{
				"handler":    router,
				"health_URI": "/api/v0/health",
				"local":      false,
			},
		),
	)
	if err != nil {
		return err
	}

	defer service.Close()

	if err := initialize(router, service); err != nil {
		service.Logger().Critical( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Error when initializing this service",
			"name", service.Name(),
			"version", version,
			"reason", err.Error(),
		)

		return err
	}

	if err := service.Run(); err != nil {
		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
