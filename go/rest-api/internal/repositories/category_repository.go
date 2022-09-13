package repositories

import (
	"context"
	"sync"

	"github.com/google/uuid"

	"example.choreo.dev/internal/models"
)

type categoryRepository struct {
	store *[]models.Category
	lock  *sync.RWMutex
}

func NewCategoryRepository(initialData []models.Category) models.CategoryRepository {
	if initialData == nil {
		initialData = make([]models.Category, 0)
	}
	return &categoryRepository{
		store: &initialData,
		lock:  &sync.RWMutex{},
	}
}

func (r categoryRepository) FindOneByName(ctx context.Context, name string) (*models.Category, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, c := range *r.store {
		if c.Name == name {
			return &c, nil
		}
	}
	return nil, nil
}

func (r categoryRepository) FindById(ctx context.Context, id string) (*models.Category, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, c := range *r.store {
		if c.Id == id {
			return &c, nil
		}
	}
	return nil, nil
}

func (r categoryRepository) Add(ctx context.Context, c *models.Category) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if c.Id == "" {
		c.Id = uuid.NewString()
	}
	v := *c
	*r.store = append(*r.store, v)
	return nil
}

func (r categoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return *r.store, nil
}
