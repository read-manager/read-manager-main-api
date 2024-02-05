package users

import (
	"github.com/google/uuid"

	"github.com/read-manager/read-manager-main-api/internal/shared/validator"
	valueobjects "github.com/read-manager/read-manager-main-api/internal/shared/value-objects"
)

type IUser interface {}

type users struct {
	Id         uuid.UUID
	Email      string
	Password   []byte
	Name       string
	Nickname   string
	ReadPoints int64
}

func NewUser(email string, password string, name string, nickname string) (IUser, map[string]string) {
	v := validator.New()
	emailVO := valueobjects.NewEmail(v, email)
    passwordVO := valueobjects.NewPassword(v, password)
    userNameVO := valueobjects.NewUserName(v, name)
    nicknameVO := valueobjects.NewNickname(v, nickname)
	if !v.Valid() {
		return nil, v.Errors
	}
    return &users{
        Id: uuid.New(),
        Email: emailVO.Get(),
        Password: passwordVO.GetHash(),
        Name: userNameVO.Get(),
        Nickname: nicknameVO.Get(),
    }, nil
}
