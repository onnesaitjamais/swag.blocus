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
	"net/http"
	"time"

	"github.com/arnumina/swag/service"
	"github.com/arnumina/swag/util"
	"github.com/gorilla/mux"
)

func loggingMiddleware(s *service.Service) mux.MiddlewareFunc {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			id := util.NewUUID()
			r.Header.Set("X-Request-ID", id)

			s.Logger().Trace( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"Request",
				"id", id,
				"host", r.RemoteAddr,
				"agent", r.Header.Get("User-Agent"),
				"method", r.Method,
				"uri", r.RequestURI,
			)

			inner.ServeHTTP(w, r)

			s.Logger().Trace( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"Request",
				"id", id,
				"elapsed", time.Since(startTime).String(),
			)
		})
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
