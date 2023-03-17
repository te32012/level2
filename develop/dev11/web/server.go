package main

import (
	"encoding/json"
	"net/http"
	"dev11/logic"
)

type MyServer struct {
	Server http.Server
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", CreateEventHandler)
	mux.HandleFunc("/update_event", UpdateEventHandler)
	mux.HandleFunc("/delete_event", DeleteEventHandler)
	mux.HandleFunc("/events_for_day", EventsForDayHandler)
	mux.HandleFunc("/events_for_week", EventsForWeekHandler)
	mux.HandleFunc("/events_for_month", EventsForMonthHandler)
	return mux
}

// дата и содержание заметки
func CreateEventHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" && request.URL.Path == "/create_event" {
		// возвращаем ошибку
	}
	var envent Envent 
	json.Unmarshal()
}

// дата и содержание
func UpdateEventHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" && request.URL.Path == "/update_event" {

	}

}

// дата и содержание
func DeleteEventHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" && request.URL.Path == "/delete_event" {

	}

}

func EventsForDayHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" && request.URL.Path == "/events_for_day" {

	}

}
func EventsForWeekHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" && request.URL.Path == "/events_for_week" {

	}

}
func EventsForMonthHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" && request.URL.Path == "/events_for_month" {

	}

}
