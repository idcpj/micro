package repository

import (
	"github.com/idcpj/micro/user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	InitTable() error
	HasTable()bool
	FindUserByName(string) (*model.User, error)
	FindUserById(int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	DeleteUserById(int64) error
	UpdateUser(*model.User) error
	findAll() ([]model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) HasTable() bool {
	return u.db.HasTable(model.User{})
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) InitTable() error {
	return u.db.CreateTable(&model.User{}).Error
}

func (u *UserRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	return user, u.db.Where("user_name=?", name).Find(user).Error
}

func (u *UserRepository) FindUserById(userId int64) (*model.User, error) {
	user := &model.User{}
	return user, u.db.First(user, userId).Error
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.db.Create(user).Error
}

func (u *UserRepository) DeleteUserById(userid int64) error {
	return u.db.Where("id=?", userid).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.db.Model(user).Update(user).Error
}

func (u *UserRepository) findAll() (userAll []model.User, err error) {
	return userAll, u.db.Find(&userAll).Error
}
