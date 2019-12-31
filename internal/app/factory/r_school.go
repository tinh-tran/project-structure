package factory

import (
	"school_project/internal/app/api/v1/school"
	"school_project/internal/pkg/database"
)

func GetSchoolRepository() school.ISchoolRepository {
	r := database.GetConnection()
	return school.NewSchoolRepository(r.Db)
}


