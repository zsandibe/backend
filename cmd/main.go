package main

import (
	"backend/internal/db"
	"backend/route/api"
	"log"
	"net/http"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v\n", err)
	}

	mux := api.Router()
	log.Println("Запуск веб-сервера на http://localhost:8080/")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("ошибка при завершении работы сервера: %v\n", err)
	}
}
