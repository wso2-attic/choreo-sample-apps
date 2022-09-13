package models

import (
	"context"
)

type Pet struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	CategoryId string `json:"category_id"`
	Available  bool   `json:"available"`
}

type PetRepository interface {
	Add(ctx context.Context, pet *Pet) error
	Update(ctx context.Context, update *Pet) error
	List(ctx context.Context) ([]Pet, error)
	GetById(ctx context.Context, id string) (*Pet, error)
}
