package main

import (
	"log"
	"net/http"
	"path/filepath"

	"scheduler/internal/api"
	"scheduler/internal/database"
)

func main() {
	port := "7540"

	// Инициализация БД
	err := database.Init("../scheduler.db")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("База данных инициализирована")

	// Регистрация обработчиков API
	http.HandleFunc("/api/nextdate", api.NextDateHandler)
	http.HandleFunc("/api/task/done", api.DoneTaskHandler)
	http.HandleFunc("/api/task", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        api.GetTaskHandler(w, r)
    case http.MethodPost:
        api.AddTaskHandler(w, r)
    case http.MethodPut:
        api.UpdateTaskHandler(w, r)
    case http.MethodDelete:
        api.DeleteTaskHandler(w, r)
    default:
        http.Error(w, `{"error":"Метод не поддерживается"}`, http.StatusMethodNotAllowed)
    }
})
	http.HandleFunc("/api/tasks", api.TasksHandler)

	// Файловый сервер для фронтенда
	webDir := filepath.Join("..", "web")
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	log.Printf("Сервер запущен на порту %s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
