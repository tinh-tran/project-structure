package factory

import (
	"school_project/internal/app/api/v1/auth"
	"school_project/internal/pkg/database"
)

func GetAuthRepository () auth.IAuthRepository {
	r := database.GetConnection()
	return auth.NewAuthRepository(r.Db)
}