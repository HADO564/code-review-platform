// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CodeSnippet struct {
	ID        string    `json:"id"`
	Project   *Project  `json:"project,omitempty"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Language  string    `json:"language"`
	Version   int       `json:"version"`
	CreatedBy *User     `json:"createdBy,omitempty"`
	Reviews   []*Review `json:"reviews"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
}

type Comment struct {
	ID         string  `json:"id"`
	Review     *Review `json:"review,omitempty"`
	User       *User   `json:"user,omitempty"`
	Content    string  `json:"content"`
	LineNumber *int    `json:"lineNumber,omitempty"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

type Mutation struct {
}

type NewCodeSnippet struct {
	ProjectID   string `json:"projectId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Language    string `json:"language"`
	CreatedByID string `json:"createdById"`
}

type NewComment struct {
	ReviewID   string `json:"reviewId"`
	UserID     string `json:"userId"`
	Content    string `json:"content"`
	LineNumber *int   `json:"lineNumber,omitempty"`
}

type NewProject struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	OwnerID     *string `json:"ownerId,omitempty"`
	OwnerName   string  `json:"ownerName"`
}

type NewProjectCollaborator struct {
	ProjectID string           `json:"projectId"`
	UserID    string           `json:"userId"`
	Role      CollaboratorRole `json:"role"`
}

type NewReview struct {
	SnippetID  string       `json:"snippetId"`
	ReviewerID string       `json:"reviewerId"`
	Status     ReviewStatus `json:"status"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Project struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Description   *string                `json:"description,omitempty"`
	Owner         *User                  `json:"owner,omitempty"`
	OwnerName     string                 `json:"ownerName"`
	CodeSnippets  []*CodeSnippet         `json:"codeSnippets"`
	Collaborators []*ProjectCollaborator `json:"collaborators"`
	CreatedAt     string                 `json:"createdAt"`
	UpdatedAt     string                 `json:"updatedAt"`
}

type ProjectCollaborator struct {
	Project *Project         `json:"project"`
	User    *User            `json:"user"`
	Role    CollaboratorRole `json:"role"`
}

type Query struct {
}

type Review struct {
	ID        string       `json:"id"`
	Snippet   *CodeSnippet `json:"snippet,omitempty"`
	Reviewer  *User        `json:"reviewer,omitempty"`
	Status    ReviewStatus `json:"status"`
	Comments  []*Comment   `json:"comments"`
	CreatedAt string       `json:"createdAt"`
	UpdatedAt string       `json:"updatedAt"`
}

type UpdatedProject struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	OwnerName   *string `json:"owner_name,omitempty"`
}

type UpdatedUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
	Projects  []*Project `json:"projects"`
	Reviews   []*Review  `json:"reviews"`
}

type CollaboratorRole string

const (
	CollaboratorRoleOwner       CollaboratorRole = "OWNER"
	CollaboratorRoleContributor CollaboratorRole = "CONTRIBUTOR"
	CollaboratorRoleViewer      CollaboratorRole = "VIEWER"
)

var AllCollaboratorRole = []CollaboratorRole{
	CollaboratorRoleOwner,
	CollaboratorRoleContributor,
	CollaboratorRoleViewer,
}

func (e CollaboratorRole) IsValid() bool {
	switch e {
	case CollaboratorRoleOwner, CollaboratorRoleContributor, CollaboratorRoleViewer:
		return true
	}
	return false
}

func (e CollaboratorRole) String() string {
	return string(e)
}

func (e *CollaboratorRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CollaboratorRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CollaboratorRole", str)
	}
	return nil
}

func (e CollaboratorRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ReviewStatus string

const (
	ReviewStatusPending  ReviewStatus = "PENDING"
	ReviewStatusApproved ReviewStatus = "APPROVED"
	ReviewStatusRejected ReviewStatus = "REJECTED"
)

var AllReviewStatus = []ReviewStatus{
	ReviewStatusPending,
	ReviewStatusApproved,
	ReviewStatusRejected,
}

func (e ReviewStatus) IsValid() bool {
	switch e {
	case ReviewStatusPending, ReviewStatusApproved, ReviewStatusRejected:
		return true
	}
	return false
}

func (e ReviewStatus) String() string {
	return string(e)
}

func (e *ReviewStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReviewStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReviewStatus", str)
	}
	return nil
}

func (e ReviewStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
