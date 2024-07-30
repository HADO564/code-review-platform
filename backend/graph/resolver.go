package graph

//go:generate go run github.com/99designs/gqlgen generate
import (
	"encoding/json"
	"log"

	"backend/database"
	"backend/graph/model"

	"github.com/supabase-community/supabase-go"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users        []*model.User
	projects     []*model.Project
	codeSnippets []*model.CodeSnippet
	reviews      []*model.Review
	comments     []*model.Comment
	collabs      []*model.ProjectCollaborator
}

func NewResolver() *Resolver {
	r := &Resolver{}
	r.loadData()
	return r
}

func (r *Resolver) loadData() {
	client := database.GetClient()

	// Load users
	users, err := fetchUsers(client)
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}
	r.users = users

	// Load projects
	projects, err := fetchProjects(client)
	if err != nil {
		log.Fatalf("Error fetching projects: %v", err)
	}
	r.projects = projects

	// Load code snippets
	codeSnippets, err := fetchCodeSnippets(client)
	if err != nil {
		log.Fatalf("Error fetching code snippets: %v", err)
	}
	r.codeSnippets = codeSnippets

	// Load reviews
	reviews, err := fetchReviews(client)
	if err != nil {
		log.Fatalf("Error fetching reviews: %v", err)
	}
	r.reviews = reviews

	// Load comments
	comments, err := fetchComments(client)
	if err != nil {
		log.Fatalf("Error fetching comments: %v", err)
	}
	r.comments = comments

	// Load collaborators
	collabs, err := fetchCollaborators(client)
	if err != nil {
		log.Fatalf("Error fetching collaborators: %v", err)
	}
	r.collabs = collabs
}

func fetchUsers(client *supabase.Client) ([]*model.User, error) {
	var users []*model.User
	data, _, err := client.From("users").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func fetchProjects(client *supabase.Client) ([]*model.Project, error) {
	var projects []*model.Project
	data, _, err := client.From("projects").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}


func fetchCodeSnippets(client *supabase.Client) ([]*model.CodeSnippet, error) {
	var codeSnippets []*model.CodeSnippet
	data, _, err := client.From("snippets").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &codeSnippets); err != nil {
		return nil, err
	}
	return codeSnippets, nil
}


func fetchReviews(client *supabase.Client) ([]*model.Review, error) {
	var reviews []*model.Review
	data, _, err := client.From("reviews").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}


func fetchComments(client *supabase.Client) ([]*model.Comment, error) {
	var comments []*model.Comment
	data, _, err := client.From("comments").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func fetchCollaborators(client *supabase.Client) ([]*model.ProjectCollaborator, error) {
	var collabs []*model.ProjectCollaborator
	data, _, err := client.From("collaborators").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &collabs); err != nil {
		return nil, err
	}
	return collabs, nil
}


