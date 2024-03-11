-- name: CreateSequence :one
INSERT INTO sequences (
  name,
  open_tracking,
  click_trancking
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetSequence :one
SELECT * FROM sequences
WHERE id = $1 LIMIT 1;

-- name: UpdateSequenceTracking :one
UPDATE sequences
  set open_tracking = $2,
  click_trancking = $3
WHERE id = $1
RETURNING *;

-- name: DeleteSequence :exec
DELETE FROM sequences
WHERE id = $1;

-- name: CreateSequenceStep :one
INSERT INTO sequence_steps (
  sequence_id,
  subject,
  content,
  step_index
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetSequenceSteps :many
SELECT * FROM sequence_steps
WHERE sequence_id = $1;

-- name: GetSequenceStep :one
SELECT * FROM sequence_steps
WHERE id = $1 LIMIT 1;

-- name: UpdateSequenceStepDetails :one
UPDATE sequence_steps
  set subject = $2,
  content = $3
WHERE id = $1
RETURNING *;

-- name: DeleteSequenceStep :exec
DELETE FROM sequence_steps
WHERE id = $1;