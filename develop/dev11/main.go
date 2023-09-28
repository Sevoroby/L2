package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var maxID int                       // максимальный ID события
var eventMap = make(map[int]*Event) // мапа для хранения событий
var usersMap = make(map[int]*User)  // мапа для хранения пользователей

// Event - Событие
type Event struct {
	ID     int
	UserID int
	Date   time.Time
	Title  string
}

// User - Пользователь
type User struct {
	ID   int
	Name string
}

// Config - Конфиг
type Config struct {
	Port string
}

// Создание конфига
func readConfig() *Config {
	config := &Config{}
	config.Port = "8080"
	return config
}
func main() {
	// Чтение конфига
	port := readConfig().Port
	// Генерация 100 пользователей
	for i := 1; i <= 100; i++ {
		usersMap[i] = &User{ID: 0, Name: "User" + strconv.Itoa(i)}
	}
	// Подключение обработчиков
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	logHandler := logMiddleware(http.DefaultServeMux)

	log.Println("Сервер запущен на порту", port)
	err := http.ListenAndServe(port, logHandler)
	log.Fatal(err)
}

// Middleware для логгирования запросов
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// Обработчик для создания события
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	mapForJSON := map[string]interface{}{}

	userID, err := parseAndValidateUserID(r)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusBadRequest, mapForJSON)
		return
	}

	date, err := parseAndValidateDate(r)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusBadRequest, mapForJSON)
		return
	}
	title, err := parseAndValidateTitle(r)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusBadRequest, mapForJSON)
		return
	}

	event, err := createEvent(date, title, userID)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusServiceUnavailable, mapForJSON)
		return
	}
	mapForJSON = map[string]interface{}{"result": "Событие успешно создано", "data": event}
	sendJSONResponse(w, http.StatusOK, mapForJSON)
}

// Обработчик для изменения события
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	mapForJSON := map[string]interface{}{}

	eventID, err := parseAndValidateEventID(r)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusBadRequest, mapForJSON)
		return
	}
	userID, err := parseAndValidateUserID(r)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusBadRequest, mapForJSON)
		return
	}

	date, err := parseAndValidateDate(r)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusBadRequest, mapForJSON)
		return
	}
	title, err := parseAndValidateTitle(r)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusBadRequest, mapForJSON)
		return
	}
	event, err := updateEvent(eventID, date, title, userID)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusServiceUnavailable, mapForJSON)
		return
	}
	mapForJSON = map[string]interface{}{"result": "Событие успешно изменено", "data": event}
	sendJSONResponse(w, http.StatusOK, mapForJSON)
}

// Обработчик для удаления события
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	mapForJSON := map[string]interface{}{}

	eventID, err := parseAndValidateEventID(r)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusBadRequest, mapForJSON)
		return
	}
	err = deleteEvent(eventID)
	if err != nil {
		mapForJSON = map[string]interface{}{"error": err.Error()}
		sendJSONResponse(w, http.StatusServiceUnavailable, mapForJSON)
		return
	}
	mapForJSON = map[string]interface{}{"result": "Событие успешно удалено"}
	sendJSONResponse(w, http.StatusOK, mapForJSON)
}

// Обработчик для получения событий за день
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	mapForJSON := map[string]interface{}{}
	res := getEventsForDay()
	mapForJSON = map[string]interface{}{"result": res}
	sendJSONResponse(w, http.StatusOK, mapForJSON)
}

// Обработчик для получения событий за неделю
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	res := getEventsForWeek()
	mapForJSON := map[string]interface{}{"result": res}
	sendJSONResponse(w, http.StatusOK, mapForJSON)
}

// Обработчик для получения событий за месяц
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	res := getEventsForMonth()
	mapForJSON := map[string]interface{}{"result": res}
	sendJSONResponse(w, http.StatusOK, mapForJSON)
}

// Преобразование ответа в JSON и отправка
func sendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)

}

// Парсинг и валидация ID события
func parseAndValidateEventID(r *http.Request) (int, error) {
	eventIDStr := r.FormValue("id")
	if len(eventIDStr) < 1 {
		return 0, fmt.Errorf("Параметр 'id' обязателен")
	}
	eventIDInt, err := strconv.Atoi(eventIDStr)
	if err != nil {
		return 0, fmt.Errorf("Параметр 'id' должен быть числом")
	}
	if _, inMap := eventMap[eventIDInt]; !inMap {
		return 0, fmt.Errorf("Не существует события с таким id")
	}
	return eventIDInt, nil
}

// Парсинг и валидация параметра даты
func parseAndValidateDate(r *http.Request) (time.Time, error) {
	dateStr := r.FormValue("date")
	if len(dateStr) < 1 {
		return time.Time{}, fmt.Errorf("Параметр 'date' обязателен")
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("Параметр 'date' должен быть в формате 'YYYY-MM-DD'")
	}
	return date, nil
}

// Парсинг и валидация параметра названия
func parseAndValidateTitle(r *http.Request) (string, error) {
	titleStr := r.FormValue("title")
	if len(titleStr) < 1 {
		return "", fmt.Errorf("Параметр 'title' обязателен")
	}
	return titleStr, nil
}

// Парсинг и валидация параметра ID пользователя
func parseAndValidateUserID(r *http.Request) (int, error) {
	userIDStr := r.FormValue("user_id")
	if len(userIDStr) < 1 {
		return 0, fmt.Errorf("Параметр 'user_id' обязателен")
	}
	userIDInt, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, fmt.Errorf("Параметр 'user_id' должен быть числом")
	}
	return userIDInt, nil
}

// Генерация ID события
func generateEventID() int {
	maxID++
	return maxID
}

// Создание события
func createEvent(date time.Time, title string, userID int) (*Event, error) {
	if _, inMap := usersMap[userID]; !inMap {
		return nil, fmt.Errorf("Не существует пользователя с таким id")
	}
	eventID := generateEventID()
	event := &Event{
		ID:     eventID,
		UserID: userID,
		Date:   date,
		Title:  title,
	}

	eventMap[eventID] = event

	return event, nil
}

// Изменение события
func updateEvent(eventID int, date time.Time, title string, userID int) (*Event, error) {
	if _, inMap := eventMap[userID]; !inMap {
		return nil, fmt.Errorf("Не существует события с таким id")
	}
	if _, inMap := usersMap[userID]; !inMap {
		return nil, fmt.Errorf("Не существует пользователя с таким id")
	}
	event := eventMap[eventID]
	event.Date = date
	event.Title = title
	event.UserID = userID
	return event, nil
}

// Получение событий за день
func getEventsForDay() []*Event {
	res := []*Event{}
	for _, v := range eventMap {
		if dateEqual(v.Date, time.Now()) {
			res = append(res, v)
		}
	}
	return res
}

// Получение событий за неделю
func getEventsForWeek() []*Event {
	res := []*Event{}
	weekStart := time.Now()
	weekEnd := time.Now().Add(time.Hour * 24 * 7)
	for _, v := range eventMap {
		if (v.Date.After(weekStart) || dateEqual(v.Date, weekStart)) && (v.Date.Before(weekEnd) || dateEqual(v.Date, weekEnd)) {
			res = append(res, v)
		}
	}
	return res
}

// Получение событий за месяц
func getEventsForMonth() []*Event {
	res := []*Event{}
	monthStart := time.Now()
	monthEnd := time.Now().Add(time.Hour * 24 * 30)
	for _, v := range eventMap {
		if (v.Date.After(monthStart) || dateEqual(v.Date, monthStart)) && (v.Date.Before(monthEnd) || dateEqual(v.Date, monthEnd)) {
			res = append(res, v)
		}
	}
	return res
}

// Удаление события
func deleteEvent(eventID int) error {
	if _, inMap := eventMap[eventID]; !inMap {
		return fmt.Errorf("Не существует события с таким id")
	}
	delete(eventMap, eventID)
	return nil
}

// Сравнения дат без времени
func dateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
