package handler

import (
	"context"
	"github.com/idcpj/micro/domain/model"
	"github.com/idcpj/micro/domain/service"
	"github.com/idcpj/micro/proto/user"
)

type User struct {
	userDataService service.IUserDataService
}

func NewUser(userDataService service.IUserDataService) *User {
	return &User{userDataService: userDataService}
}

func (u *User) Register(ctx context.Context, request *user.UserRegisterRequest, response *user.UserRegisterResponse) error {
	userInfo := &model.User{
		UserName:    request.GetUserName(),
		FirstName:   request.GetFirstName(),
		HasPassword: request.GetPwd(),
	}
	_, err := u.userDataService.AddUser(userInfo)
	if err != nil {
		return err
	}
	response.Message = "添加成功"
	return nil
}

func (u *User) Login(ctx context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	isOk, err := u.userDataService.CheckPwd(request.UserName, request.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, request *user.UserInfoRequest, response *user.UserInfoResponse) error {
	userInfo, err := u.userDataService.FindUserByName(request.UserName)
	if err != nil {
		return err
	}

	UserForResponse(userInfo,response)

	return nil
}

func UserForResponse(userInfo *model.User,response *user.UserInfoResponse)  {
	response.UserName = userInfo.UserName
	response.FirstName = userInfo.FirstName
	response.UserId = userInfo.ID
}
