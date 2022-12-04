package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jonasalessi/go-graphQL/graph/model"
)

type CategoryDB struct {
	data []*model.Category
}

func (c *CategoryDB) Save(name, description string) (category *model.Category) {
	id := uuid.New().String()
	category = &model.Category{
		ID:          id,
		Name:        name,
		Description: &description,
	}
	c.data = append(c.data, category)
	return
}

func (c *CategoryDB) GetAll() []*model.Category {
	return c.data
}

func (c *CategoryDB) GetById(id string) (*model.Category, error) {
	for _, cat := range c.data {
		if cat.ID == id {
			return cat, nil
		}
	}
	return nil, errors.New("Category not found!")
}

func (c *CategoryDB) ExistsById(id string) bool {
	_, err := c.GetById(id)
	if err != nil {
		return false
	}
	return true
}
