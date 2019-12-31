package school

import "database/sql"

type ISchoolRepository interface {
	CreateSchool (formData School)(int, error)
}

type schoolRepository struct {
	db *sql.DB
}

func NewSchoolRepository(db *sql.DB) *schoolRepository {
	return &schoolRepository{db: db}
}

func (s schoolRepository) CreateSchool (formData School)(int, error) {
	return 1, nil
}
