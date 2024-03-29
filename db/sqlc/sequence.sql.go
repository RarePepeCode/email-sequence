// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: sequence.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSequence = `-- name: CreateSequence :one
INSERT INTO sequences (
  name,
  open_tracking,
  click_trancking
) VALUES (
  $1, $2, $3
) RETURNING id, name, open_tracking, click_trancking
`

type CreateSequenceParams struct {
	Name           pgtype.Text
	OpenTracking   pgtype.Bool
	ClickTrancking pgtype.Bool
}

func (q *Queries) CreateSequence(ctx context.Context, arg CreateSequenceParams) (Sequence, error) {
	row := q.db.QueryRow(ctx, createSequence, arg.Name, arg.OpenTracking, arg.ClickTrancking)
	var i Sequence
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OpenTracking,
		&i.ClickTrancking,
	)
	return i, err
}

const createSequenceStep = `-- name: CreateSequenceStep :one
INSERT INTO sequence_steps (
  sequence_id,
  subject,
  content,
  step_index,
  wait_days
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, sequence_id, subject, content, step_index, wait_days
`

type CreateSequenceStepParams struct {
	SequenceID int64
	Subject    pgtype.Text
	Content    pgtype.Text
	StepIndex  int32
	WaitDays   pgtype.Int4
}

func (q *Queries) CreateSequenceStep(ctx context.Context, arg CreateSequenceStepParams) (SequenceStep, error) {
	row := q.db.QueryRow(ctx, createSequenceStep,
		arg.SequenceID,
		arg.Subject,
		arg.Content,
		arg.StepIndex,
		arg.WaitDays,
	)
	var i SequenceStep
	err := row.Scan(
		&i.ID,
		&i.SequenceID,
		&i.Subject,
		&i.Content,
		&i.StepIndex,
		&i.WaitDays,
	)
	return i, err
}

const deleteSequence = `-- name: DeleteSequence :exec
DELETE FROM sequences
WHERE id = $1
`

func (q *Queries) DeleteSequence(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteSequence, id)
	return err
}

const deleteSequenceStep = `-- name: DeleteSequenceStep :exec
DELETE FROM sequence_steps
WHERE id = $1
`

func (q *Queries) DeleteSequenceStep(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteSequenceStep, id)
	return err
}

const getSequence = `-- name: GetSequence :one
SELECT id, name, open_tracking, click_trancking FROM sequences
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSequence(ctx context.Context, id int64) (Sequence, error) {
	row := q.db.QueryRow(ctx, getSequence, id)
	var i Sequence
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OpenTracking,
		&i.ClickTrancking,
	)
	return i, err
}

const getSequenceStep = `-- name: GetSequenceStep :one
SELECT id, sequence_id, subject, content, step_index, wait_days FROM sequence_steps
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSequenceStep(ctx context.Context, id int64) (SequenceStep, error) {
	row := q.db.QueryRow(ctx, getSequenceStep, id)
	var i SequenceStep
	err := row.Scan(
		&i.ID,
		&i.SequenceID,
		&i.Subject,
		&i.Content,
		&i.StepIndex,
		&i.WaitDays,
	)
	return i, err
}

const getSequenceSteps = `-- name: GetSequenceSteps :many
SELECT id, sequence_id, subject, content, step_index, wait_days FROM sequence_steps
WHERE sequence_id = $1
ORDER BY step_index ASC
`

func (q *Queries) GetSequenceSteps(ctx context.Context, sequenceID int64) ([]SequenceStep, error) {
	rows, err := q.db.Query(ctx, getSequenceSteps, sequenceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SequenceStep
	for rows.Next() {
		var i SequenceStep
		if err := rows.Scan(
			&i.ID,
			&i.SequenceID,
			&i.Subject,
			&i.Content,
			&i.StepIndex,
			&i.WaitDays,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSequenceStepDetails = `-- name: UpdateSequenceStepDetails :one
UPDATE sequence_steps
  set subject = $2,
  content = $3
WHERE id = $1
RETURNING id, sequence_id, subject, content, step_index, wait_days
`

type UpdateSequenceStepDetailsParams struct {
	ID      int64
	Subject pgtype.Text
	Content pgtype.Text
}

func (q *Queries) UpdateSequenceStepDetails(ctx context.Context, arg UpdateSequenceStepDetailsParams) (SequenceStep, error) {
	row := q.db.QueryRow(ctx, updateSequenceStepDetails, arg.ID, arg.Subject, arg.Content)
	var i SequenceStep
	err := row.Scan(
		&i.ID,
		&i.SequenceID,
		&i.Subject,
		&i.Content,
		&i.StepIndex,
		&i.WaitDays,
	)
	return i, err
}

const updateSequenceTracking = `-- name: UpdateSequenceTracking :one
UPDATE sequences
  set open_tracking = $2,
  click_trancking = $3
WHERE id = $1
RETURNING id, name, open_tracking, click_trancking
`

type UpdateSequenceTrackingParams struct {
	ID             int64
	OpenTracking   pgtype.Bool
	ClickTrancking pgtype.Bool
}

func (q *Queries) UpdateSequenceTracking(ctx context.Context, arg UpdateSequenceTrackingParams) (Sequence, error) {
	row := q.db.QueryRow(ctx, updateSequenceTracking, arg.ID, arg.OpenTracking, arg.ClickTrancking)
	var i Sequence
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OpenTracking,
		&i.ClickTrancking,
	)
	return i, err
}
