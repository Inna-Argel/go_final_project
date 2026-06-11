package api

import (
    "encoding/json"
    "net/http"

    "scheduler/internal/database"
)

const defaultTasksLimit = 50

func TasksHandler(w http.ResponseWriter, r *http.Request) {
    tasks, err := database.Tasks(defaultTasksLimit)
    if err != nil {
        http.Error(w, `{"error":"Ошибка получения задач"}`, http.StatusInternalServerError)
        return
    }

    if tasks == nil {
        tasks = []database.Task{}
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{"tasks": tasks})
}
