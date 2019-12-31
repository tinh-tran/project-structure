package auth

import
(
	"database/sql"
)

type IAuthRepository interface {
	GetUserByUserNameAndPassWord(userName, passWord string)(UserInfo, error)
	GetUserByID(userId int) (UserInfo, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{db: db}
}

func (a *authRepository) GetUserByUserNameAndPassWord(userName, passWord string)(UserInfo, error)  {
	return UserInfo{}, nil
}

func(a *authRepository) GetUserByID(userId int) (UserInfo, error){
	return UserInfo{}, nil
}
