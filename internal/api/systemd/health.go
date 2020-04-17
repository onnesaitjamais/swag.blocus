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

package systemd

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes AFAIRE
func Routes(r *mux.Router) {
	r.HandleFunc(
		"/health",
		func(w http.ResponseWriter, _ *http.Request) { // For systemd watchdog /////////////////////////////////////////
			w.WriteHeader(http.StatusNoContent)
		},
	).Methods("GET")
}

/*
######################################################################################################## @(°_°)@ #######
*/
