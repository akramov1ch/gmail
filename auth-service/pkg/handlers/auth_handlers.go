package handlers

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base32"
	"time"

	"auth-service/internal/db"
	"auth-service/pkg/models"
	"auth-service/proto"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	proto.UnimplementedAuthServiceServer
}

func (s *AuthService) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = db.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", req.Name, req.Email, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return &proto.SignUpResponse{Message: "User registered successfully"}, nil
}

func (s *AuthService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user := models.User{}
	err := db.DB.QueryRow("SELECT id, password FROM users WHERE email = $1", req.Email).Scan(&user.ID, &user.Password)
	if err == sql.ErrNoRows {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	token, refreshToken, err := generateTokens(user.ID)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{Token: token, RefreshToken: refreshToken}, nil
}

func generateTokens(userID int64) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}

	refreshToken := base32.StdEncoding.EncodeToString(generateRandomBytes(32))
	return tokenString, refreshToken, nil
}

func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}
