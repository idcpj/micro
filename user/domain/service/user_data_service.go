package service

import (
	"errors"
	"github.com/idcpj/micro/domain/model"
	"github.com/idcpj/micro/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ERROR_VALIDATE_PASSWORD = errors.New("密码比对错误")
	ERROR_USER_INVALID      = errors.New("账号无效")
)

type IUserDataService interface {
	AddUser(user *model.User) (id int64, error error)
	DeleteUser(id int64) error
	UpdateUser(user *model.User, isChangepwd bool) error
	FindUserByName(name string) (*model.User, error)
	CheckPwd(userName string, pwd string) (isok bool, err error)
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{
		UserRepository: userRepository,
	}
}

func (u *UserDataService) AddUser(user *model.User) (id int64, error error) {
	pwdByte, err := GeneratePassword(user.HasPassword)
	if err != nil {
		return user.ID, err
	}
	user.HasPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u *UserDataService) DeleteUser(id int64) error {
	return u.UserRepository.DeleteUserById(id)
}

func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) error {
	if isChangePwd {
		pwd, err := GeneratePassword(user.HasPassword)
		if err != nil {
			return ERROR_VALIDATE_PASSWORD
		}
		user.HasPassword = string(pwd)
	}
	return u.UserRepository.UpdateUser(user)
}

func (u *UserDataService) FindUserByName(name string) (*model.User, error) {
	return u.UserRepository.FindUserByName(name)
}

func (u *UserDataService) CheckPwd(userName string, pwd string) (isok bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, ERROR_USER_INVALID
	}
	return ValiDatePwd(pwd, user.HasPassword)
}

func GeneratePassword(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

func ValiDatePwd(pwd string, hased string) (isok bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(hased)); err != nil {
		return false, ERROR_VALIDATE_PASSWORD
	}
	return true, nil
}
