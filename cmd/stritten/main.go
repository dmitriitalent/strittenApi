package main

import (
	"context"
	"log"

	stritten "github.com/dmitriitalent/strittenApi"
	"github.com/dmitriitalent/strittenApi/internal/handler"
	"github.com/dmitriitalent/strittenApi/internal/repository"
	"github.com/dmitriitalent/strittenApi/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error occured while initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
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

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := stritten.Server{}
	if err := srv.Run(viper.GetString("PORT"), handlers.InitRouters()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())

		srv.Shutdown(context.Background())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
