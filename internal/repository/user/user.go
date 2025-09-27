package user

import (
	"context"
	"easycart/internal/domain"
	"easycart/internal/repository"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"time"
)

type UserStorage struct {
	db *goqu.Database
}

func NewUserStorage(db *goqu.Database) repository.UserStorager {
	return &UserStorage{
		db: db,
	}
}

func (c *UserStorage) Create(ctx context.Context, user domain.CreateUser) (uuid.UUID, error) {
	data := CreateUserToModel(user)
	data.ID = uuid.New()
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	_, err := c.db.Insert(userTableName).Rows(data).Executor().Exec()
	if err != nil {
		return uuid.UUID{}, err
	}
	return data.ID, nil
}

func (c *UserStorage) Find(ctx context.Context, id string) (domain.User, error) {
	filter := goqu.Ex{"id": id}

}
