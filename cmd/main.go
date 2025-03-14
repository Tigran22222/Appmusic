package main

import (
	"appmusic/pkg/handler"
	"appmusic/pkg/repository"
	"appmusic/pkg/server"
	"appmusic/pkg/service"
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	// Инициализация конфигурации
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())

	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error loading .env variables: %s", err.Error())
	}

	// Создание подключения к базе данных с использованием sqlx
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	// Создание репозитория с подключением к базе данных
	repos := repository.NewRepository(db) // Передаем db как *sqlx.DB
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Создание и запуск сервера
	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running server: %s", err.Error())
		}
	}()
	logrus.Print("Appmusic Started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Appmusic Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error occurred while shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("error occurred while closing db: %s", err.Error())
	}
}

// Инициализация конфигурации
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

/*import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=your_password dbname=tables sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("База данных недоступна:", err)
	}

	fmt.Println("Успешное подключение к базе данных!")
}

*/
