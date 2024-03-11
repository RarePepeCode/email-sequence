package middleware

import (
	"context"
	"log"
	"net/http"

	db "github.com/RarePepeCode/email-sequence/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	queries db.Queries
}

func StartServer() *http.Server {
	server := &Server{
		queries: *dbConn(),
	}
	return server.initRouter()
}

func dbConn() *db.Queries {
	pool, err := pgxpool.New(context.Background(), "postgresql://user:pass@localhost:5432/sequence_email")
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	return db.New(pool)

}

func (server Server) initRouter() *http.Server {
	router := gin.Default()

	router.POST("/sequence", server.CreateSequence)
	router.PATCH("sequence/:id", server.UpdateSequenceTracking)
	router.POST("/sequence-step", server.CreateSequenceStep)

	router.PATCH("sequence-step/:id", server.UpdateSequenceStepDetails)
	router.DELETE("sequence-step/:id", server.DeleteSequenceStep)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	return srv
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
