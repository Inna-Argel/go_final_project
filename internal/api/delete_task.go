package api

import (
    "encoding/json"
    "net/http"

    "scheduler/internal/database"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.FormValue("id")
    if idStr == "" {
        http.Error(w, `{"error":"Не указан идентификатор"}`, http.StatusBadRequest)
        return
    }

    err := database.DeleteTask(idStr)
    if err != nil {
        http.Error(w, `{"error":"Ошибка удаления задачи"}`, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{})
}
