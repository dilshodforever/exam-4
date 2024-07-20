package postgres

import (
	"database/sql"
	"log/slog"

	pb "github.com/dilshodforever/4-oyimtixon-auth-service/genprotos/auth"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type AuthStorage struct {
	db     *sql.DB
	client *redis.Client
}

func NewAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{db: db}
}

func (p *AuthStorage) Register(req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userId := uuid.NewString()
	query := `
		INSERT INTO users (id, username, email, password_hash, full_name)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, username, email, full_name, created_at
	`
	var user pb.RegisterResponse
	err := p.db.QueryRow(query, userId, req.Username, req.Email, req.Password, req.FullName).Scan(
		&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt,
	)
	if err != nil {
		slog.Info(err.Error())
		return nil, err
	}

	tokenQuery := `
		INSERT INTO refresh_tokens (username, token)
		VALUES ($1, $2)
	`
	_, err = p.db.Exec(tokenQuery, req.Username, req.Token)
	if err != nil {
		slog.Info(err.Error())
		return nil, err
	}

	return &user, nil
}

func (p *AuthStorage) Login(req *pb.LoginRequest) (*pb.LoginResponse, error) {
	query := `
		SELECT username
		FROM users
		WHERE username = $1 AND password_hash = $2
	`
	var username string
	err := p.db.QueryRow(query, req.Username, req.Password).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.LoginResponse{Message: "Invalid username or password", Success: false}, nil
		}
		return nil, err
	}
	var token string
	getTokenQuery := `
		SELECT token
		FROM refresh_tokens
		WHERE username = $1
	`
	err = p.db.QueryRow(getTokenQuery, username).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.LoginResponse{Message: "Token not found", Success: false}, nil
		}
		return nil, err
	}
	return &pb.LoginResponse{Token: token, Message: "Login successful", Success: true}, nil
}

func (p *AuthStorage) Logout(req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	query := `
		DELETE FROM refresh_tokens
		WHERE token = $1
	`
	_, err := p.db.Exec(query, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.LogoutResponse{Message: "Logged out successfully"}, nil
}

func (p *AuthStorage) ResetPassword(req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	query := `
		UPDATE users
		SET password_hash = $1
		WHERE email = $2 and username=$3
	`
	_, err := p.db.Exec(query, req.NewPassword, req.Email, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.ResetPasswordResponse{Message: "Password reset successfully"}, nil
}