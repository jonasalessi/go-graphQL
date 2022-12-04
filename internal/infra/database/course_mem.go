package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jonasalessi/go-graphQL/graph/model"
)

type CourseDB struct {
	data                 []*model.Course
	oneCourseOneCategory map[string]string
}

func (c *CourseDB) Save(name, description, categoryId string) (course *model.Course) {
	if c.oneCourseOneCategory == nil {
		c.oneCourseOneCategory = make(map[string]string)
	}
	id := uuid.New().String()
	course = &model.Course{
		ID:          id,
		Name:        name,
		Description: &description,
	}
	c.oneCourseOneCategory[id] = categoryId
	c.data = append(c.data, course)
	return
}

func (c *CourseDB) GetCategoryIdByCourse(id string) (string, error) {
	categoryId := c.oneCourseOneCategory[id]
	if categoryId == "" {
		return "", errors.New("This course does not have category!")
	}
	return categoryId, nil
}

func (c *CourseDB) GetAll() []*model.Course {
	return c.data
}

func (c *CourseDB) GetById(id string) (*model.Course, error) {
	for _, cat := range c.data {
		if cat.ID == id {
			return cat, nil
		}
	}
	return nil, errors.New("Course not found!")
}

func (c *CourseDB) GetAllByCategory(id string) []*model.Course {
	courses := []*model.Course{}
	for courseId, categoryId := range c.oneCourseOneCategory {
		if categoryId == id {
			course, err := c.GetById(courseId)
			if err == nil {
				courses = append(courses, course)
			}
		}
	}
	return courses
}
