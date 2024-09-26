package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"products/configs"
	"products/db"
	"products/logger"
	"products/pkg/controllers"
	"products/server"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

// @title Products API
// @version 1.0
// @description API Server for Products Application

// @host localhost:8686
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %s", err)
	}

	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	if err := logger.Init(); err != nil {
		log.Fatalf("Ошибка инициализации логгера: %s", err)
	}

	var err error
	err = db.ConnectToDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %s", err)
	}

	if err = db.Migrate(); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %s", err)
	}
	err = db.InsertSeeds()
	if err != nil {
		logger.Error.Println("Insert data not succesfuly")
	}

	mainServer := new(server.Server)
	go func() {
		if err = mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.RunRoutes()); err != nil {
			log.Fatalf("Ошибка при запуске HTTP сервера: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\nНачало завершения программ\n")

	if sqlDB, err := db.GetDBConn().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Ошибка при закрытии соединения с БД: %s", err)
		}
	} else {
		log.Fatalf("Ошибка при получении *sql.DB из GORM: %s", err)
	}
	fmt.Println("Соединение с БД успешно закрыто")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = mainServer.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}

	fmt.Println("HTTP-сервис успешно выключен")
	fmt.Println("Конец завершения программы")
}
