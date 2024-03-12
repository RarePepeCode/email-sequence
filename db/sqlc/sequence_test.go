package db

import (
	"context"
	"log"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	seq := createTestSequence(t)
	queries.DeleteSequence(context.Background(), seq.ID)
}

func TestGetSequence(t *testing.T) {
	seq := createTestSequence(t)
	defer queries.DeleteSequence(context.Background(), seq.ID)
	fetchSeq, err := queries.GetSequence(context.Background(), seq.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchSeq)
	require.Equal(t, seq.ID, fetchSeq.ID)
}

func TestUpdateSequenceTracking(t *testing.T) {
	seq := createTestSequence(t)
	defer queries.DeleteSequence(context.Background(), seq.ID)
	require.False(t, seq.OpenTracking.Bool)

	params := UpdateSequenceTrackingParams{
		ID:             seq.ID,
		OpenTracking:   pgtype.Bool{Bool: true, Valid: true},
		ClickTrancking: pgtype.Bool{Bool: false, Valid: true},
	}
	updatedSeq, err := queries.UpdateSequenceTracking(context.Background(), params)
	require.NoError(t, err)
	require.False(t, seq.OpenTracking.Bool == updatedSeq.OpenTracking.Bool)
	require.False(t, seq.ClickTrancking.Bool == updatedSeq.ClickTrancking.Bool)
}

func TestDeleteSequence(t *testing.T) {
	seq := createTestSequence(t)
	queries.DeleteSequence(context.Background(), seq.ID)

	fetchSeq, err := queries.GetSequence(context.Background(), seq.ID)
	require.Error(t, err)
	require.Empty(t, fetchSeq)
}

func TestCreateSequenceStep(t *testing.T) {
	seq := createTestSequence(t)
	seqStep := createTestSequenceStep(t, int(seq.ID), 1)
	defer queries.DeleteSequence(context.Background(), seq.ID)
	defer queries.DeleteSequenceStep(context.Background(), seqStep.ID)

	params := CreateSequenceStepParams{
		SequenceID: int64(seq.ID),
		Subject:    pgtype.Text{String: "Test Subject", Valid: true},
		Content:    pgtype.Text{String: "Test Content", Valid: true},
		StepIndex:  1,
	}
	failedConstraintseqStep, err := queries.CreateSequenceStep(context.Background(), params)
	require.Error(t, err)
	require.Equal(t, int64(0), failedConstraintseqStep.ID)
}

func TestGetSequenceSteps(t *testing.T) {
	seq := createTestSequence(t)
	seqStep1 := createTestSequenceStep(t, int(seq.ID), 2)
	seqStep2 := createTestSequenceStep(t, int(seq.ID), 1)
	defer queries.DeleteSequence(context.Background(), seq.ID)
	defer queries.DeleteSequenceStep(context.Background(), seqStep1.ID)
	defer queries.DeleteSequenceStep(context.Background(), seqStep2.ID)

	seqSteps, err := queries.GetSequenceSteps(context.Background(), seq.ID)
	log.Println(seqStep1)
	log.Println(seqStep2)

	require.NoError(t, err)
	require.Equal(t, 2, len(seqSteps))
	require.Equal(t, 1, int(seqSteps[0].StepIndex))
}

func TestUpdateSequenceStepDetails(t *testing.T) {
	seq := createTestSequence(t)
	seqStep := createTestSequenceStep(t, int(seq.ID), 1)
	defer queries.DeleteSequence(context.Background(), seq.ID)
	defer queries.DeleteSequenceStep(context.Background(), seqStep.ID)
	params := UpdateSequenceStepDetailsParams{
		ID:      seqStep.ID,
		Subject: pgtype.Text{String: "New Subject", Valid: true},
		Content: pgtype.Text{String: "New Content", Valid: true},
	}

	updatedStep, err := queries.UpdateSequenceStepDetails(context.Background(), params)
	require.NoError(t, err)
	require.Equal(t, params.Subject, updatedStep.Subject)
	require.Equal(t, params.Content, updatedStep.Content)
}

func TestGetSequenceStep(t *testing.T) {
	seq := createTestSequence(t)
	seqStep := createTestSequenceStep(t, int(seq.ID), 1)
	defer queries.DeleteSequence(context.Background(), seq.ID)
	defer queries.DeleteSequenceStep(context.Background(), seqStep.ID)

	fetchSeqStep, err := queries.GetSequenceStep(context.Background(), seqStep.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchSeqStep)
	require.Equal(t, seqStep.ID, fetchSeqStep.ID)
}

func TestDeleteSequenceStep(t *testing.T) {
	seq := createTestSequence(t)
	seqStep := createTestSequenceStep(t, int(seq.ID), 1)
	defer queries.DeleteSequence(context.Background(), seq.ID)

	queries.DeleteSequenceStep(context.Background(), seqStep.ID)

	fetchSeqStep, err := queries.GetSequenceStep(context.Background(), seq.ID)
	require.Error(t, err)
	require.Empty(t, fetchSeqStep)
}

func createTestSequence(t *testing.T) Sequence {
	params := CreateSequenceParams{
		Name:           pgtype.Text{String: "Test Sequence", Valid: true},
		OpenTracking:   pgtype.Bool{Bool: false, Valid: true},
		ClickTrancking: pgtype.Bool{Bool: true, Valid: true},
	}
	seq, err := queries.CreateSequence(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, seq)
	require.Equal(t, params.Name, seq.Name)
	return seq
}

func createTestSequenceStep(t *testing.T, seqId int, index int) SequenceStep {
	params := CreateSequenceStepParams{
		SequenceID: int64(seqId),
		Subject:    pgtype.Text{String: "Test Subject", Valid: true},
		Content:    pgtype.Text{String: "Test Content", Valid: true},
		StepIndex:  int32(index),
	}
	seqStep, err := queries.CreateSequenceStep(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, seqStep)
	require.Equal(t, params.SequenceID, seqStep.SequenceID)
	return seqStep
}
