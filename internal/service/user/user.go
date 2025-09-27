package user

import (
	"context"
	adto "easycart/internal/controller/http/auth/dto"
	udto "easycart/internal/controller/http/user/dto"
	"easycart/internal/domain"
	"easycart/internal/repository"
	"easycart/internal/service"
	"easycart/pkg/jwt"
	"easycart/pkg/password"
	"errors"
	"github.com/google/uuid"
)

type UserService struct {
	repo      repository.UserStorager
	secretKey []byte
}

func NewUserService(repo repository.UserStorager) service.UserServicer {
	return &UserService{
		repo: repo,
	}
}

func (c *UserService) Register(ctx context.Context, req adto.RegisterRequest) (uuid.UUID, error) {
	hashed, err := password.Generate(req.Password)
	if err != nil {
		return uuid.UUID{}, err
	}
	user := domain.CreateUser{
		Name:        req.Name,
		Password:    hashed,
		PhoneNumber: req.PhoneNumber,
		Avatar:      req.Avatar,
	}

	return c.repo.Create(ctx, user)
}

func (c *UserService) Login(ctx context.Context, req adto.LoginRequest) (string, int64, error) {
	u, err := c.repo.Find(ctx, "", req.PhoneNumber)
	if err != nil {
		return "", 0, err
	}
	if res := password.Compare(req.Password, u[0].Password); res == false {
		return "", 0, errors.New("invalid password")
	}
	token, expiredAt, err := jwt.Encode(jwt.JWT{
		u[0].ID.String(),
		u[0].PhoneNumber,
	}, c.secretKey)
	if err != nil {
		return "", 0, err
	}
	return *token, *expiredAt, nil
}

func (c *UserService) GetById(ctx context.Context, userId string) (*udto.User, error) {
	user, err := c.repo.Find(ctx, userId, "")
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return &udto.User{}, err
		}
		return &udto.User{}, err
	}

	res := UserDomainToDto(user[0])

	return &res, nil
}

func (c *UserService) List(ctx context.Context) (*[]udto.User, error) {
	user, err := c.repo.Find(ctx, "", "")
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return &[]udto.User{}, err
		}
		return &[]udto.User{}, err
	}

	var usersRes []udto.User
	for _, u := range user {
		usersRes = append(usersRes, UserDomainToDto(u))
	}
	return &usersRes, nil
}
