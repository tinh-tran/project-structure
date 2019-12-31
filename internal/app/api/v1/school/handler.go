package school

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"school_project/internal/pkg/util"
)

type schoolController struct {
	schoolService ISchoolServices
}

func NewSchoolController(schoolService ISchoolServices) *schoolController {
	return &schoolController{schoolService: schoolService}
}

func (c *schoolController) CreateSchool(w http.ResponseWriter, r *http.Request) {
	var formData School
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &formData)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	int, err := c.schoolService.CreateSchool(formData)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.RespondJSONSuccess(w, fmt.Sprintf("Create School Success %d", int))
	return
}
