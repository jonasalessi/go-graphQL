package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/jonasalessi/go-graphQL/graph/generated"
	"github.com/jonasalessi/go-graphQL/graph/model"
)

// Courses is the resolver for the courses field.
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	return r.courseDb.GetAllByCategory(obj.ID), nil
}

// Category is the resolver for the category field.
func (r *courseResolver) Category(ctx context.Context, obj *model.Course) (*model.Category, error) {
	categoryId, err := r.courseDb.GetCategoryIdByCourse(obj.ID)
	if err != nil {
		return nil, err
	}
	category, err := r.categoryDb.GetById(categoryId)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	return r.categoryDb.Save(input.Name, *input.Description), nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	categoryExists := r.categoryDb.ExistsById(input.CategoryID)
	if categoryExists {
		return r.courseDb.Save(input.Name, *input.Description, input.CategoryID), nil
	}
	return nil, errors.New("Category ID " + input.CategoryID + " not exists!")
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.categoryDb.GetAll(), nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.courseDb.GetAll(), nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
