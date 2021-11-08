// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package artsign

import (
	"fmt"
	"io"
	"strconv"
)

type CreateUserLikeInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	UserID           int     `json:"userID"`
	WorkID           int     `json:"workID"`
}

type CreateUserLikePayload struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type CreateUserTreasureInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	UserID           int     `json:"userID"`
	WorkID           int     `json:"workID"`
}

type CreateUserTreasurePayload struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type Status string

const (
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
)

var AllStatus = []Status{
	StatusInProgress,
	StatusCompleted,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusInProgress, StatusCompleted:
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
