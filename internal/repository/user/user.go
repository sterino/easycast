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

func (c *UserStorage) Create(_ context.Context, user domain.CreateUser) (uuid.UUID, error) {
	data := CreateUserToModel(user)
	data.ID = uuid.New()
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	_, err := c.db.Insert(UserTableName).Rows(data).Executor().Exec()
	if err != nil {
		return uuid.UUID{}, err
	}
	return data.ID, nil
}

func (c *UserStorage) Find(_ context.Context, id, phoneNumber string) ([]domain.User, error) {
	filter := goqu.Ex{"id": id}
	if id != "" {
		filter["id"] = id
	}
	if phoneNumber != "" {
		filter["phone_number"] = phoneNumber
	}
	query := c.db.From(UserTableName)
	if !filter.IsEmpty() {
		query = query.Where(filter)
	}

	var result []User
	err := query.Order(goqu.L("created_at").Asc()).Limit(100).ScanStructs(&result)
	if err != nil {
		return nil, err
	}
	var res []domain.User
	for _, v := range result {
		res = append(res, UserModelToDomain(v))
	}
	return res, nil
}
