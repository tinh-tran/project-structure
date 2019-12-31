package router

import (
	"net/http"
	"school_project/internal/app/api/v1/auth"
	"school_project/internal/app/factory"
)

func AuthRouter(Dispatch map[string]http.HandlerFunc) map[string]http.HandlerFunc{
	authService := factory.GetAuthService()
	authController := auth.NewAuthController(authService)
	Dispatch["CheckLogin"] = authController.CheckLogin
	return Dispatch
}