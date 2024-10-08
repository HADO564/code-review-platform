package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"
	"fmt"

	pkg "github.com/yourusername/code-review-platform/db"
	"github.com/yourusername/code-review-platform/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	createdUser, err := pkg.CreateUser(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return createdUser, nil
}

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	panic(fmt.Errorf("not implemented: CreateProject - createProject"))
}

// CreateCodeSnippet is the resolver for the createCodeSnippet field.
func (r *mutationResolver) CreateCodeSnippet(ctx context.Context, input model.NewCodeSnippet) (*model.CodeSnippet, error) {
	panic(fmt.Errorf("not implemented: CreateCodeSnippet - createCodeSnippet"))
}

// CreateReview is the resolver for the createReview field.
func (r *mutationResolver) CreateReview(ctx context.Context, input model.NewReview) (*model.Review, error) {
	panic(fmt.Errorf("not implemented: CreateReview - createReview"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// AddProjectCollaborator is the resolver for the addProjectCollaborator field.
func (r *mutationResolver) AddProjectCollaborator(ctx context.Context, input model.NewProjectCollaborator) (*model.ProjectCollaborator, error) {
	panic(fmt.Errorf("not implemented: AddProjectCollaborator - addProjectCollaborator"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdatedUser) (*model.User, error) {
	updated_user, err := pkg.UpdateUser(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return updated_user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string, input model.UpdatedUser) (*model.User, error) {
	deletedUser, err := pkg.DeleteUser(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %v", err)
	}
	return deletedUser, nil
}

// DeleteProject is the resolver for the deleteProject field.
func (r *mutationResolver) DeleteProject(ctx context.Context, id string, input model.UpdatedProject) (*model.Project, error) {
	panic(fmt.Errorf("not implemented: DeleteProject - deleteProject"))
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, id string, input model.UpdatedProject) (*model.Project, error) {
	panic(fmt.Errorf("not implemented: UpdateProject - updateProject"))
}

// GetUserByEmail is the resolver for the getUserByEmail field.
func (r *queryResolver) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return pkg.GetUserByEmail(ctx, email)
}

// GetUserByUsername is the resolver for the getUserByUsername field.
func (r *queryResolver) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUserByUsername - getUserByUsername"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Project is the resolver for the project field.
func (r *queryResolver) Project(ctx context.Context, id string) (*model.Project, error) {
	panic(fmt.Errorf("not implemented: Project - project"))
}

// ProjectsByOwnerName is the resolver for the projectsByOwnerName field.
func (r *queryResolver) ProjectsByOwnerName(ctx context.Context, ownerName string) ([]*model.Project, error) {
	panic(fmt.Errorf("not implemented: ProjectsByOwnerName - projectsByOwnerName"))
}

// CodeSnippet is the resolver for the codeSnippet field.
func (r *queryResolver) CodeSnippet(ctx context.Context, id string) (*model.CodeSnippet, error) {
	panic(fmt.Errorf("not implemented: CodeSnippet - codeSnippet"))
}

// Review is the resolver for the review field.
func (r *queryResolver) Review(ctx context.Context, id string) (*model.Review, error) {
	panic(fmt.Errorf("not implemented: Review - review"))
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, id string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Comment - comment"))
}

// ProjectCollaborator is the resolver for the projectCollaborator field.
func (r *queryResolver) ProjectCollaborator(ctx context.Context, projectID string, userID string) (*model.ProjectCollaborator, error) {
	panic(fmt.Errorf("not implemented: ProjectCollaborator - projectCollaborator"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
