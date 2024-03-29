package middleware

import (
	"net/http"

	db "github.com/RarePepeCode/email-sequence/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type SeqResponse struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	OpenTrackingEnabled  bool   `json:"open_tracking"`
	ClickTrackingEnabled bool   `json:"click_tracking"`
}

type SeqStepResponse struct {
	Id       int    `json:"id"`
	SeqId    int    `json:"seq_id"`
	Index    int    `json:"index"`
	Subject  string `json:"subjectd"`
	Content  string `json:"content"`
	WaitDays int    `json:"wait_days"`
}

type createSequenceRequest struct {
	Name            string `json:"name"`
	Open_tracking   bool   `json:"open_tracking"`
	Click_trancking bool   `json:"click_tracking"`
}

type updateSequenceTrackingRequest struct {
	Open_tracking   bool `json:"open_tracking"`
	Click_trancking bool `json:"click_tracking"`
}

type updateSequenceStepDetailsRequest struct {
	Subject string `json:"subject" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type createSequenceStepRequest struct {
	SeqId    int    `json:"seq_id" binding:"required,min=1"`
	Index    int    `json:"index" binding:"required,min=1"`
	Subject  string `json:"subject"`
	Content  string `json:"content"`
	WaitDays int    `json:"wait_days"`
}

type reqeustPathId struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server Server) CreateSequence(ctx *gin.Context) {
	var body createSequenceRequest
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := db.CreateSequenceParams{
		Name:           pgtype.Text{String: body.Name, Valid: true},
		OpenTracking:   pgtype.Bool{Bool: body.Open_tracking, Valid: true},
		ClickTrancking: pgtype.Bool{Bool: body.Click_trancking, Valid: true},
	}
	seq, err := server.queries.CreateSequence(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, createSequenceResponse(&seq))
}

func (server Server) UpdateSequenceTracking(ctx *gin.Context) {
	var body updateSequenceTrackingRequest
	var pathId reqeustPathId

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err = ctx.ShouldBindUri(&pathId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, (errorResponse(err)))
		return
	}

	_, err = server.queries.GetSequence(ctx, pathId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No sequence exist for passed path id"})
		return
	}

	params := db.UpdateSequenceTrackingParams{
		ID:             pathId.ID,
		OpenTracking:   pgtype.Bool{Bool: body.Open_tracking, Valid: true},
		ClickTrancking: pgtype.Bool{Bool: body.Click_trancking, Valid: true},
	}
	updatedSeq, err := server.queries.UpdateSequenceTracking(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, createSequenceResponse(&updatedSeq))
}

func (server Server) UpdateSequenceStepDetails(ctx *gin.Context) {
	var body updateSequenceStepDetailsRequest
	var pathId reqeustPathId

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err = ctx.ShouldBindUri(&pathId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.queries.GetSequenceStep(ctx, pathId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No sequence step exist for passed path id"})
		return
	}

	params := db.UpdateSequenceStepDetailsParams{
		ID:      pathId.ID,
		Subject: pgtype.Text{String: body.Subject, Valid: true},
		Content: pgtype.Text{String: body.Content, Valid: true},
	}
	updatedSeqStep, err := server.queries.UpdateSequenceStepDetails(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, createSequenceStepResponse(&updatedSeqStep))
}

func (server Server) CreateSequenceStep(ctx *gin.Context) {
	var body createSequenceStepRequest
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.queries.GetSequence(ctx, int64(body.SeqId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No sequence exist for passed seq_id"})
		return
	}

	params := db.CreateSequenceStepParams{
		StepIndex:  int32(body.Index),
		SequenceID: int64(body.SeqId),
		Content:    pgtype.Text{String: body.Content, Valid: true},
		Subject:    pgtype.Text{String: body.Subject, Valid: true},
		WaitDays:   pgtype.Int4{Int32: int32(body.WaitDays), Valid: true},
	}
	seqStep, err := server.queries.CreateSequenceStep(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, createSequenceStepResponse(&seqStep))
}

func (server Server) DeleteSequenceStep(ctx *gin.Context) {
	var pathId reqeustPathId
	err := ctx.ShouldBindUri(&pathId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.queries.DeleteSequenceStep(ctx, pathId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.Status(http.StatusOK)
}

func createSequenceResponse(seq *db.Sequence) *SeqResponse {
	return &SeqResponse{
		Id:                   int(seq.ID),
		Name:                 seq.Name.String,
		ClickTrackingEnabled: seq.ClickTrancking.Bool,
		OpenTrackingEnabled:  seq.OpenTracking.Bool,
	}
}

func createSequenceStepResponse(seqStep *db.SequenceStep) *SeqStepResponse {
	return &SeqStepResponse{
		Id:       int(seqStep.ID),
		SeqId:    int(seqStep.SequenceID),
		Index:    int(seqStep.StepIndex),
		Subject:  seqStep.Subject.String,
		Content:  seqStep.Content.String,
		WaitDays: int(seqStep.WaitDays.Int32),
	}
}
