package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mauricionofre/person-api/internal/config"
	"github.com/mauricionofre/person-api/internal/handler"
	"github.com/mauricionofre/person-api/internal/repository"
	"github.com/mauricionofre/person-api/internal/service"
	"github.com/mauricionofre/person-api/pkg/rabbitmq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	rabbitMQ, err := rabbitmq.NewRabbitMQ(cfg.RabbitMQURL, cfg.RabbitMQExchange)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	personRepo := repository.NewPersonRepository(db)
	personService := service.NewPersonService(personRepo, rabbitMQ)
	personHandler := handler.NewPersonHandler(personService)

	r := gin.Default()

	r.POST("/persons", personHandler.Create)
	r.GET("/persons/:id", personHandler.GetByID)
	r.PUT("/persons/:id", personHandler.Update)
	r.DELETE("/persons/:id", personHandler.Delete)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
