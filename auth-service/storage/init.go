package postgres

import (
	pbAuth "github.com/dilshodforever/4-oyimtixon-auth-service/genprotos/auth"
	pbUser "github.com/dilshodforever/4-oyimtixon-auth-service/genprotos/user"
)

type InitRoot interface {
	Auth() Auth
	User() User
}

type Auth interface {
	Register(req *pbAuth.RegisterRequest) (*pbAuth.RegisterResponse, error)
	Login(req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error)
	Logout(req *pbAuth.LogoutRequest) (*pbAuth.LogoutResponse, error)
	ResetPassword(req *pbAuth.ResetPasswordRequest) (*pbAuth.ResetPasswordResponse, error)
}

type User interface {
	GetProfile(req *pbUser.GetProfileRequest) (*pbUser.GetProfileResponse, error)
	UpdateProfile(req *pbUser.UpdateProfileRequest) (*pbUser.UpdateProfileResponse, error)
	ChangePassword(req *pbUser.ChangePasswordRequest) (*pbUser.ChangePasswordResponse, error)
	UpdateXps(req *pbUser.UpdateXpRequest) (*pbUser.UpdateXpResponse, error) 
}
