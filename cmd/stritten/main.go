package main

import (
	"context"
	"log"

	stritten "github.com/dmitriitalent/strittenApi"
	router "github.com/dmitriitalent/strittenApi/internal"
	"github.com/dmitriitalent/strittenApi/internal/handlers"
	"github.com/dmitriitalent/strittenApi/internal/repositories"
	"github.com/dmitriitalent/strittenApi/internal/services"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error occured while initializing config: %s", err.Error())
	}

	db, err := repositories.NewPostgresDB(repositories.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Error occured while initializing postgres DB: %s", err.Error())
	}

	repos := repositories.NewRepositories(db)
	services := services.NewServices(repos)
	handlers := handlers.NewHandlers(services)

	srv := stritten.Server{}
	router := *router.NewRouter(*handlers);
	if err := srv.Run(viper.GetString("PORT"), router.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())

		srv.Shutdown(context.Background())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
