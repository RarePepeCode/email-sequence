package middleware

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func testMainX(*testing.M) {
	// pool, err := pgxpool.New(context.Background(), "postgresql://user:pass@localhost:5432/sequence_email")
	// if err != nil {
	// 	log.Fatal("Cannot connect to db:", err)
	// }

	// queries = New(pool)

	// os.Exit(m.Run())
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
