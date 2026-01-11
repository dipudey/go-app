package user

import (
	"github.com/dipudey/go-app/internal/auth"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service interface {
	GetAllUsers() ([]Response, error)
	CreateUser(req CreateUserRequest) (Response, error)
	Login(req LoginRequest) (string, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllUsers() ([]Response, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var response []Response
	for _, u := range users {
		response = append(response, Response{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return response, nil
}

func (s *service) CreateUser(req CreateUserRequest) (Response, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return Response{}, err
	}
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	createdUser, err := s.repo.Create(user)
	if err != nil {
		return Response{}, err
	}

	return Response{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, nil
}

func (s *service) Login(req LoginRequest) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", err
	}
	
	//-- Generate JWT token
	token, err := auth.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		return "", err
	}
	return token, nil
}
