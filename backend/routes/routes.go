package routes

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"orderkuota/adapters/handlers"
	"orderkuota/adapters/repository"
	"orderkuota/adapters/service"
	"orderkuota/config"
	"orderkuota/middleware"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func NewRouter(router *mux.Router, db *sql.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	//Public routes (tanpa jwt)
	router.HandleFunc("/register", userHandler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", userHandler.Login).Methods("POST", "OPTIONS")

	// Protected routes (menggunakan jwt)
	authRoutes := router.PathPrefix("/api").Subrouter()
	authRoutes.Use(middleware.AuthMiddleware)
	authRoutes.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET", "OPTIONS")
	authRoutes.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET", "OPTIONS")
	authRoutes.HandleFunc("/users/{id}", userHandler.UpdateUserByID).Methods("PUT", "OPTIONS")
	authRoutes.HandleFunc("/users/{id}", userHandler.DeleteUserByID).Methods("DELETE", "OPTIONS")
}

func StartServer() {
	config.InitiateLog()
	defer config.CloseLog()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "9090"
	}

	db, _ := config.InitDB()
	defer db.Close()

	router := mux.NewRouter()
	middleware.EnableCORS(router)
	NewRouter(router, db)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("starting server on port " + port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("\nShutting down server....")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	wg.Wait()
	fmt.Println("Server gracefully stopped")
}
