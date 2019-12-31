package auth

import (
	"encoding/json"
	"school_project/internal/app/models"
)

type UserRole struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Permission  string `json:"permission"`
	ClassList   string `json:"class"`
	models.BaseModel
}

type UserInfo struct {
	UserId      int             `json:"userID"`
	RoleId      string          `json:"roleId"`
	RoleTitle   string          `json:"roleTitle"`
	Department  string          `json:"department"`
	School      string          `json:"school"`
	Permissions json.RawMessage `json:"permission"`
	IsActive    bool            `json:"isActive"`
	ClassIds    []string        `json:"classIds"`
	models.BaseModel
}
