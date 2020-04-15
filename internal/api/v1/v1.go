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
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/arnumina/swag/component/registry"
	"github.com/arnumina/swag/service"
	"github.com/arnumina/swag/util/failure"
	"github.com/gorilla/mux"
)

func errorHandler(s *service.Service) func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		s.Logger().Error( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Reverse proxy error",
			"request", r.Header.Get("X-Request-ID"),
			"url", r.URL.String(),
			"reason", err.Error(),
		)

		http.Error(w, err.Error(), http.StatusBadGateway)
	}
}

func simpleFilter(s *service.Service) func(*registry.Service) bool {
	interval := s.Registry().Interval()
	now := time.Now()

	return func(s *registry.Service) bool {
		return s.Status == "running" && now.Sub(s.Heartbeat) <= time.Duration(interval)*time.Second
	}
}

func findEndpoint(s *service.Service, name string) (string, int, error) {
	services, err := s.Registry().Find(name)
	if err != nil {
		return "", 0, err
	}

	services = services.Filter(simpleFilter(s))

	if services.Len() == 0 {
		return "", 0,
			failure.New(nil).
				Set("name", name).
				Msg("there are no services running with this name") ////////////////////////////////////////////////////
	}

	services.Shuffle()

	service := services[0]
	host := service.FQDN
	port := service.Port

	if host == s.FQDN() {
		host = ""
	}

	return host, port, nil
}

func newReverseProxy(s *service.Service, r *http.Request) (*httputil.ReverseProxy, error) {
	vars := mux.Vars(r)
	uri := vars["uri"]

	host, port, err := findEndpoint(s, vars["service"])
	if err != nil {
		return nil, err
	}

	target, err := url.Parse(fmt.Sprintf("http://%s:%d/api/%s", host, port, uri))
	if err != nil {
		return nil, err
	}

	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		req.URL.RawQuery = target.RawQuery
		req.Header.Set("User-Agent", "swag.blocus/v"+s.Version())
	}

	reverseProxy := &httputil.ReverseProxy{
		Director:     director,
		ErrorHandler: errorHandler(s),
	}

	return reverseProxy, nil
}

func gateKeeper(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reverseProxy, err := newReverseProxy(s, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		reverseProxy.ServeHTTP(w, r)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
