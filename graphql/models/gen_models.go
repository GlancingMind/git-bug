// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"

	"github.com/MichaelMure/git-bug/bug"
	"github.com/MichaelMure/git-bug/identity"
	"github.com/MichaelMure/git-bug/util/git"
)

// An object that has an author.
type Authored interface {
	IsAuthored()
}

type AddCommentInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// "The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
	// The first message of the new bug.
	Message string `json:"message"`
	// The collection of file's hash required for the first message.
	Files []git.Hash `json:"files"`
}

type AddCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug *bug.Snapshot `json:"bug"`
	// The resulting operation.
	Operation *bug.AddCommentOperation `json:"operation"`
}

// The connection type for Bug.
type BugConnection struct {
	// A list of edges.
	Edges []*BugEdge      `json:"edges"`
	Nodes []*bug.Snapshot `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type BugEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *bug.Snapshot `json:"node"`
}

type ChangeLabelInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// "The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
	// The list of label to add.
	Added []string `json:"added"`
	// The list of label to remove.
	Removed []string `json:"Removed"`
}

type ChangeLabelPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug *bug.Snapshot `json:"bug"`
	// The resulting operation.
	Operation *bug.LabelChangeOperation `json:"operation"`
	// The effect each source label had.
	Results []*bug.LabelChangeResult `json:"results"`
}

type CloseBugInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// "The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
}

type CloseBugPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug *bug.Snapshot `json:"bug"`
	// The resulting operation.
	Operation *bug.SetStatusOperation `json:"operation"`
}

type CommentConnection struct {
	Edges      []*CommentEdge `json:"edges"`
	Nodes      []*bug.Comment `json:"nodes"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

type CommentEdge struct {
	Cursor string       `json:"cursor"`
	Node   *bug.Comment `json:"node"`
}

type CommitAsNeededInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// "The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
}

type CommitAsNeededPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug *bug.Snapshot `json:"bug"`
}

type CommitInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// "The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
}

type CommitPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug *bug.Snapshot `json:"bug"`
}

type IdentityConnection struct {
	Edges      []*IdentityEdge      `json:"edges"`
	Nodes      []identity.Interface `json:"nodes"`
	PageInfo   *PageInfo            `json:"pageInfo"`
	TotalCount int                  `json:"totalCount"`
}

type IdentityEdge struct {
	Cursor string             `json:"cursor"`
	Node   identity.Interface `json:"node"`
}

type LabelConnection struct {
	Edges      []*LabelEdge `json:"edges"`
	Nodes      []bug.Label  `json:"nodes"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

type LabelEdge struct {
	Cursor string    `json:"cursor"`
	Node   bug.Label `json:"node"`
}

type NewBugInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// "The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The title of the new bug.
	Title string `json:"title"`
	// The first message of the new bug.
	Message string `json:"message"`
	// The collection of file's hash required for the first message.
	Files []git.Hash `json:"files"`
}

type NewBugPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The created bug.
	Bug *bug.Snapshot `json:"bug"`
	// The resulting operation.
	Operation *bug.CreateOperation `json:"operation"`
}

type OpenBugInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// "The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
}

type OpenBugPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug *bug.Snapshot `json:"bug"`
	// The resulting operation.
	Operation *bug.SetStatusOperation `json:"operation"`
}

// The connection type for an Operation
type OperationConnection struct {
	Edges      []*OperationEdge `json:"edges"`
	Nodes      []bug.Operation  `json:"nodes"`
	PageInfo   *PageInfo        `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

// Represent an Operation
type OperationEdge struct {
	Cursor string        `json:"cursor"`
	Node   bug.Operation `json:"node"`
}

// Information about pagination in a connection.
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor string `json:"startCursor"`
	// When paginating forwards, the cursor to continue.
	EndCursor string `json:"endCursor"`
}

type SetTitleInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// "The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
	// The new title.
	Title string `json:"title"`
}

type SetTitlePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug *bug.Snapshot `json:"bug"`
	// The resulting operation
	Operation *bug.SetTitleOperation `json:"operation"`
}

// The connection type for TimelineItem
type TimelineItemConnection struct {
	Edges      []*TimelineItemEdge `json:"edges"`
	Nodes      []bug.TimelineItem  `json:"nodes"`
	PageInfo   *PageInfo           `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

// Represent a TimelineItem
type TimelineItemEdge struct {
	Cursor string           `json:"cursor"`
	Node   bug.TimelineItem `json:"node"`
}

type LabelChangeStatus string

const (
	LabelChangeStatusAdded         LabelChangeStatus = "ADDED"
	LabelChangeStatusRemoved       LabelChangeStatus = "REMOVED"
	LabelChangeStatusDuplicateInOp LabelChangeStatus = "DUPLICATE_IN_OP"
	LabelChangeStatusAlreadyExist  LabelChangeStatus = "ALREADY_EXIST"
	LabelChangeStatusDoesntExist   LabelChangeStatus = "DOESNT_EXIST"
)

var AllLabelChangeStatus = []LabelChangeStatus{
	LabelChangeStatusAdded,
	LabelChangeStatusRemoved,
	LabelChangeStatusDuplicateInOp,
	LabelChangeStatusAlreadyExist,
	LabelChangeStatusDoesntExist,
}

func (e LabelChangeStatus) IsValid() bool {
	switch e {
	case LabelChangeStatusAdded, LabelChangeStatusRemoved, LabelChangeStatusDuplicateInOp, LabelChangeStatusAlreadyExist, LabelChangeStatusDoesntExist:
		return true
	}
	return false
}

func (e LabelChangeStatus) String() string {
	return string(e)
}

func (e *LabelChangeStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LabelChangeStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LabelChangeStatus", str)
	}
	return nil
}

func (e LabelChangeStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Status string

const (
	StatusOpen   Status = "OPEN"
	StatusClosed Status = "CLOSED"
)

var AllStatus = []Status{
	StatusOpen,
	StatusClosed,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusOpen, StatusClosed:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
