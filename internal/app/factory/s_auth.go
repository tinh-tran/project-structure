package factory

import (
	"school_project/internal/app/api/v1/auth"
)

func GetAuthService() auth.IAuthService {
	repo := GetAuthRepository()
	return auth.NewAuthService(repo)
}