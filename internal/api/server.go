package api

import (
	"fmt"

	db "github.com/ShadrackAdwera/go-multitenancy/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  db.TxStore
}

func NewServer(store db.TxStore) *Server {
	srv := Server{
		store: store,
	}

	router := gin.Default()

	router.POST("/api/tenant", srv.createTenant)

	srv.router = router

	return &srv
}

func (srv *Server) Start(addr string) error {
	return srv.router.Run(addr)
}

func errJSON(err error) gin.H {
	return gin.H{"error": fmt.Sprintf("err: %s", err.Error())}
}
