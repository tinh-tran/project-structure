package auth

import "errors"

type IAuthService interface {
	GetUserByUserNameAndPassWord(userName, passWord string)(UserInfo, error)
	GetUserByID(userId int) (UserInfo, error)
}

type authService struct {
	authRepo IAuthRepository
}

func NewAuthService(repository IAuthRepository) *authService {
	return &authService{authRepo: repository}
}

func (s *authService) GetUserByUserNameAndPassWord(userName, passWord string)(UserInfo, error)  {
	if userName == "" || passWord == "" {
		return UserInfo{}, errors.New("Username and password required")
	}
	return s.authRepo.GetUserByUserNameAndPassWord(userName, passWord)
}

func (s *authService) GetUserByID(userId int) (UserInfo, error) {
	return s.authRepo.GetUserByID(userId)
}
