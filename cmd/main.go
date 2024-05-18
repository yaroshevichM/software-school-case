package main

import (
	"fmt"
	"os"
	_ "time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	common "github.com/yaroshevichM/software-school-case"
	"github.com/yaroshevichM/software-school-case/pkg/handler"
	"github.com/yaroshevichM/software-school-case/pkg/repository"
	"github.com/yaroshevichM/software-school-case/pkg/scheduler"
	"github.com/yaroshevichM/software-school-case/pkg/service"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error when loading env vars: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf("Error when init config: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	scheduler := scheduler.NewScheduler(services)

	err = scheduler.AddProcessEmail("0 0 0 * * *")

	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	scheduler.Start()

	server := new(common.Server)
	if err := server.Run("3000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Recived error when started server: %s", err.Error())
	}

	select {}
}
