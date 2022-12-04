package graph

import "github.com/jonasalessi/go-graphQL/internal/infra/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	categoryDb database.CategoryDB
	courseDb   database.CourseDB
}
