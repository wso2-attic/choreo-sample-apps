package models

import (
	"context"
)

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CategoryRepository interface {
	Add(ctx context.Context, category *Category) error
	FindOneByName(ctx context.Context, name string) (*Category, error)
	FindById(ctx context.Context, id string) (*Category, error)
	GetAll(ctx context.Context) ([]Category, error)
}
