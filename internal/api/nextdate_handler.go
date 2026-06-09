package api

import (

	"net/http"
	"time"
)

// NextDateHandler обрабатывает GET /api/nextdate
func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из запроса
	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	// Если параметр now не задан, используем текущую дату
	var now time.Time
	var err error
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse("20060102", nowStr)
		if err != nil {
			http.Error(w, "Неверный формат now", http.StatusBadRequest)
			return
		}
	}

	// Вычисляем следующую дату
	nextDate, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Возвращаем результат
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Write([]byte(nextDate))
}
