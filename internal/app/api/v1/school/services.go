package school

type ISchoolServices interface {
	CreateSchool (formData School)(int, error)
}

type schoolServices struct {
	schoolRepo ISchoolRepository
}

func NewSchoolServices(schoolRepo ISchoolRepository) *schoolServices {
	return &schoolServices{schoolRepo: schoolRepo}
}

func (s schoolServices) CreateSchool(formData School) (int, error) {
	return s.schoolRepo.CreateSchool(formData)
}