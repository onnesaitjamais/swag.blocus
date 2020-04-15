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

package v1

import (
	"github.com/arnumina/swag/service"
	"github.com/gorilla/mux"
)

// Routes AFAIRE
func Routes(r *mux.Router, s *service.Service) {
	r.Use(
		loggingMiddleware(s),
	)

	r.HandleFunc("/{service}/{uri:.+}", gateKeeper(s))
}

/*
######################################################################################################## @(°_°)@ #######
*/
