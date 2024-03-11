package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSeq(t *testing.T) {
	srv := StartServer()
	params := createSequenceRequest{
		Name:            "TEST NAME",
		Open_tracking:   true,
		Click_trancking: true,
	}

	w := httptest.NewRecorder()
	rb, _ := json.Marshal(params)
	req, _ := http.NewRequest(http.MethodPost, "/sequence", bytes.NewBuffer(rb))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	srv.Handler.ServeHTTP(w, req)
	body, err := io.ReadAll(w.Result().Body)
	require.NoError(t, err)
	var resp SeqResponse
	json.Unmarshal(body, &resp)

	require.Equal(t, http.StatusCreated, w.Code)
	require.NotNil(t, resp.Id)
	require.Equal(t, params.Name, resp.Name)
	require.Equal(t, params.Open_tracking, resp.OpenTrackingEnabled)
	require.Equal(t, params.Click_trancking, resp.ClickTrackingEnabled)
}
