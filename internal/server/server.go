package server

import (
	"context"
	"fidelizou-go/internal/db"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

type Server struct {
	port int

	db *db.Queries
}

func NewServer() *http.Server {
	ctx := context.Background()
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	NewServer := &Server{
		port: port,

		db: db.NewInstance(ctx),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	authConfig()

	return server
}

func authConfig() {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/auth/google/callback"),
	)
}
