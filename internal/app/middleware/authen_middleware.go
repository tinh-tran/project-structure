package middleware

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
	global "school_project/internal/app/init_global"
	"school_project/internal/pkg/constants"
	"school_project/internal/pkg/jwt"
	"school_project/internal/pkg/util"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With,Authorization,access-control-allow-origin,tokenapi,x-access-token,x-app-id")
		if r.Method == "OPTIONS" {
			util.RespondJSONError(w, http.StatusNoContent, "Access denied")
			return
		}
		handlerName := "UNKNOWN"
		if route := mux.CurrentRoute(r); route != nil {
			routeName := route.GetName()
			if routeName != "" {
				handlerName = routeName
			}
		}
		var isPublicApi = IsPublicAPI(handlerName)
		if isPublicApi {
			next.ServeHTTP(w, r)
			return
		}
		var authTokenValue = r.Header.Get("Authorization")
		var rep = map[string]interface{}{}
		if authTokenValue == "" {
			rep["error"] = "invalid payload"
			util.RespondJSONError(w, http.StatusNoContent, "Access denied")
			return
		}
		// parseToken
		claims, err := jwt.ParseJWTToken(authTokenValue, []byte(constants.AuthTokenSalt))
		if err != nil {
			util.RespondJSONError(w, http.StatusNoContent, "Access denied")
			return
		}
		if claims.UserInfo == "" {
			util.RespondJSONError(w, http.StatusForbidden, "UnAuthenticate")
			return
		} else {
			context.Set(r, "userInfo", claims.UserInfo)
			next.ServeHTTP(w, r)
		}
	})
}

func IsPublicAPI(handleName string) bool {
	for _, val := range global.RoutersArr {
		for _, valR := range val.Data {
			if handleName == valR.Handler {
				if valR.PublicApi == true {
					return true
				}
			}
		}
	}
	return false
}
