package router

import (
	"net/http"
	"school_project/internal/app/api/v1/school"
	"school_project/internal/app/factory"
)

func SchoolRouter(Dispatch map[string]http.HandlerFunc) map[string]http.HandlerFunc{
	schoolService := factory.GetSchoolService()
	authController := school.NewSchoolController(schoolService)
	Dispatch["CreateSchool"] = authController.CreateSchool
	return Dispatch
}
