package main

import (
	"log"

	"appmusic/pkg/handler"
	"appmusic/pkg/server"
)

func main() {

	// Создаем сервер
	handlers := new(handler.Handler)
	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running server: %s", err.Error())
	}

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
