package pkg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yourusername/code-review-platform/graph/model"
	"github.com/yourusername/code-review-platform/internal/supabase"
)

func CreateProject(ctx context.Context, input *model.NewProject) (*model.Project, error) {
	var createdProject model.Project
	newProject := model.NewProject{
		Name:      input.Name,
		OwnerID:   input.OwnerID,
		OwnerName: input.OwnerName,
	}

	// Insert the new project into the database
	responseBytes, count, err := supabase.Client.From("projects").Insert(newProject, false, "", "", "exact").Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("project was not created")
	}

	err = json.Unmarshal(responseBytes, &createdProject)
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return &createdProject, nil
}

func GetProjectsByUser(ctx context.Context, name string) (*model.Project, error) {

	var project model.Project
	responseBytes, count, err := supabase.Client.From("projects").
		Select("*", "exact", false).
		Eq("ownerName", name).
		Single().
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get project by user: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("project not found")
	}

	// Unmarshal the response data into the project struct
	err = json.Unmarshal(responseBytes, &project)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal project data: %w", err)
	}

	// Now that project is populated, marshal it to JSON for printing
	jsonData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal project data to JSON: %w", err)
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
	return &project, nil
}

func DeleteProject(ctx context.Context, input *model.UpdatedProject) (*model.Project, error) {
	// This will hold the actual deleted project data
	var deletedProject model.Project

	// Execute the deletion query in Supabase
	responseBytes, count, err := supabase.Client.From("projects").
		Delete("", "exact").
		Eq("name", *input.Name). // Filtering by project name
		Execute()

	if err != nil {
		return nil, fmt.Errorf("failed to delete project: %w", err)
	}

	// Check if no project was deleted
	if count == 0 {
		return nil, fmt.Errorf("project not found or not deleted")
	}

	// Unmarshal the deleted project data from responseBytes
	err = json.Unmarshal(responseBytes, &deletedProject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal project data: %w", err)
	}

	// Return the deleted project object
	return &deletedProject, nil
}


func UpdateProject(ctx context.Context, input *model.UpdatedProject) (*model.Project, error) {
	// Ensure at least one field is provided for update
	if *input.Name == "" && *input.Description == "" && *input.OwnerName == "" {
		return nil, fmt.Errorf("at least one field (name, description, or owner_name) must be provided for update")
	}

	// Initialize the updated project with new values from input
	updatedFields := map[string]interface{}{}
	if *input.Name != "" {
		updatedFields["name"] = *input.Name
	}
	if *input.Description != "" {
		updatedFields["description"] = *input.Description
	}
	if *input.OwnerName != "" {
		updatedFields["owner_name"] = *input.OwnerName
	}

	// Temporary hardcoded project ID (replace this with your actual logic)
	var projId = "45fa9495-a35d-451b-93f5-9f17c2fc3b46"

	// This will hold the updated project data
	var updatedProject model.Project

	// Execute the update query in Supabase
	responseBytes, _, err := supabase.Client.From("projects").
		Update(updatedFields, "", "exact"). // Updating only provided fields
		Eq("id", projId).      // Filter by project ID
		Execute()

	if err != nil {
		return nil, fmt.Errorf("failed to update project: %w", err)
	}

	// Unmarshal the updated project data from responseBytes
	err = json.Unmarshal(responseBytes, &updatedProject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal project data: %w", err)
	}

	// Ensure the project was actually updated
	if updatedProject.ID == "" {
		return nil, fmt.Errorf("project not found or not updated")
	}

	// Return the updated project object
	return &updatedProject, nil
}

