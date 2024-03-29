// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Sequence struct {
	ID             int64
	Name           pgtype.Text
	OpenTracking   pgtype.Bool
	ClickTrancking pgtype.Bool
}

type SequenceStep struct {
	ID         int64
	SequenceID int64
	Subject    pgtype.Text
	Content    pgtype.Text
	StepIndex  int32
	WaitDays   pgtype.Int4
}
