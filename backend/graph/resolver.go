package graph
//go:generate go run github.com/99designs/gqlgen generate
import "backend/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	users []*model.User
	projects []*model.Project
	codeSnippets []*model.CodeSnippet
	reviews []*model.Review
	comments []*model.Comment
	collabs []*model.ProjectCollaborator
}
