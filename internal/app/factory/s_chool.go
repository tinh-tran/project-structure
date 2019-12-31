package factory

import (
	"school_project/internal/app/api/v1/school"
)

func GetSchoolService() school.ISchoolServices {
	repo := GetSchoolRepository()
	return school.NewSchoolServices(repo)
}
