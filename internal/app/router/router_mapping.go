package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	global "school_project/internal/app/init_global"
	"school_project/internal/app/middleware"
	"school_project/internal/app/router/router_defined"
)

var Dispatch map[string]http.HandlerFunc

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	//configs.VerifyAllEnv()
	InitRouterMapping()
	var routersArr = global.RoutersArr
	routersArr = router_defined.ReadRouterFile()
	for _, val := range routersArr {
		s := r.PathPrefix(fmt.Sprintf("/api/%s/", val.GroupAPI)).Subrouter()
		for _, valR := range val.Data {
			s.HandleFunc(valR.Path, Dispatch[valR.Handler]).
				Methods(valR.Methods...).
				Name(valR.Handler)
		}
	}
	r.Use(middleware.AuthenticateMiddleware, middleware.RoleMiddleware)
	return r
}
