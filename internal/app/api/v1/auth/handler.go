package auth

import (
	"net/http"

	"school_project/internal/pkg/util"
)

type authController struct {
	authService IAuthService
	authToken Authentication
}

func NewAuthController(authService IAuthService) *authController {
	authToken := NewAuthentication()
	return &authController{authService: authService, authToken:authToken}
}

func (c *authController) CheckLogin(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("username")
	passWord := r.Header.Get("password")
	userInfo , err := c.authService.GetUserByUserNameAndPassWord(userName, passWord)
	if err != nil {
		util.RespondJSONError(w, http.StatusUnauthorized, err.Error())
		return
	}
	tokenString, err := c.authToken.GenerateToken(userInfo)
	if err != nil {
		util.RespondJSONError(w, http.StatusUnauthorized, err.Error())
		return
	}
	var tokenResp = map[string]string{}
	tokenResp["token"]= tokenString
	util.RespondJSON(w, http.StatusOK, tokenString)
	return
}
