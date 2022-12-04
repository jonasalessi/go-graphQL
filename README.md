## Go + GraphQL using gqlgen

## Generating the resolvers
```
go run github.com/99designs/gqlgen generate
```

## GraphQL script
```
mutation CreateCategory {
  createCategory(input: {name: "Category A"    description: "Something else"}){
    id
    name
  }
}

query ListAllCategories {
  categories{
    id
    name
    description
  }
}

mutation CreateCourse {
  createCourse(input: {
    name: "Course G", 
    description: "Some course", 
    categoryId: ADD HERE THE CATEGORY ID
  }) {
    id
    name
  }
}

query ListAllCourses{
  courses {
    id
    name
    category {
      name
    }
  }
}
```