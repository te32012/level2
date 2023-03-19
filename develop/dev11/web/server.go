package dev11

import (
	logic "dev11/logic"
	"log"
	"net/http"
)

type Application struct {
	Server *http.Server
	Data   *logic.Database
	Logger *log.Logger
}

func (myserver *Application) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", myserver.CreateEventHandler)
	mux.HandleFunc("/update_event", myserver.UpdateEventHandler)
	mux.HandleFunc("/delete_event", myserver.DeleteEventHandler)
	mux.HandleFunc("/events_for_day", myserver.EventsForDayHandler)
	mux.HandleFunc("/events_for_week", myserver.EventsForWeekHandler)
	mux.HandleFunc("/events_for_month", myserver.EventsForMonthHandler)
	return mux
}

// дата и содержание заметки
func (myserver *Application) CreateEventHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" && request.URL.Path == "/create_event" {
		myserver.Logger.Println("ошибка обрабочика требуемый метод : POST, требуемый путь /create_event. Имеющийся метод", request.Method, "имеющийся путь", request.URL.Path)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	myserver.Logger.Println("добавляем событие в хранилище")
	err := myserver.Data.AddEventFromJson(request.Body)
	if err != nil {
		myserver.Logger.Println("Некорретое поведение бизнесс логики при POST запросе к /create_event : ошибка формата ", err.Error())
		if err.Error() == "ошибка формата данных" {
			response.WriteHeader(http.StatusBadRequest)
			return
		}
		response.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("событие создано"))
}

// дата и содержание
func (myserver *Application) UpdateEventHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" && request.URL.Path == "/update_event" {
		myserver.Logger.Println("ошибка обрабочика требуемый метод : POST, требуемый путь /update_event. Имеющийся метод", request.Method, "имеющийся путь", request.URL.Path)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	myserver.Logger.Println("обновляем событие в хранилище")
	err := myserver.Data.UpdateEventFromJson(request.Body)
	if err != nil {
		myserver.Logger.Println("Некорретое поведение бизнесс логики при POST запросе к /create_event : ошибка формата ", err.Error())
		if err.Error() == "ошибка формата данных" {
			response.WriteHeader(http.StatusBadRequest)
			return
		}
		response.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("событие обновлено"))
}

// дата и содержание
func (myserver *Application) DeleteEventHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" && request.URL.Path == "/delete_event" {
		myserver.Logger.Println("ошибка обрабочика требуемый метод : POST, требуемый путь /delete_event. Имеющийся метод", request.Method, "имеющийся путь", request.URL.Path)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	myserver.Logger.Println("удаляем событие из хранилища")
	err := myserver.Data.RemoveEventFromJson(request.Body)
	if err != nil {
		myserver.Logger.Println("Некорретое поведение бизнесс логики при POST запросе к /create_event : ошибка формата ", err.Error())
		if err.Error() == "ошибка формата данных" {
			response.WriteHeader(http.StatusBadRequest)
			return
		}
		response.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("событие удалено"))
}

func (myserver *Application) EventsForDayHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" && request.URL.Path == "/events_for_day" {
		myserver.Logger.Println("ошибка обрабочика требуемый метод : GET, требуемый путь /events_for_day. Имеющийся метод", request.Method, "имеющийся путь", request.URL.Path)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	var key string = "abc"
	if request.URL.Query().Has("data") {
		key = request.URL.Query().Get("data")
	}
	myserver.Logger.Println("получаем данные о событии из хранилища за день")
	ans, err := myserver.Data.GetAllEventInDataDay(key)
	if err != nil {
		myserver.Logger.Println("Некорретое поведение бизнесс логики при GET запросе к /events_for_day : ошибка формата ", err.Error())
		if err.Error() == "ошибка формата даты" {
			response.WriteHeader(http.StatusBadRequest)
			return
		}
		response.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(ans)
}
func (myserver *Application) EventsForWeekHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" && request.URL.Path == "/events_for_week" {
		myserver.Logger.Println("ошибка обрабочика требуемый метод : GET, требуемый путь /events_for_week. Имеющийся метод", request.Method, "имеющийся путь", request.URL.Path)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	var key string = "abc"
	if request.URL.Query().Has("data") {
		key = request.URL.Query().Get("data")
	}
	myserver.Logger.Println("получаем данные о событии из хранилища за неделю")

	ans, err := myserver.Data.GetAllEventInDataWeek(key)
	if err != nil {
		myserver.Logger.Println("Некорретое поведение бизнесс логики при GET запросе к /events_for_week : ошибка формата ", err.Error())
		if err.Error() == "ошибка формата даты" {
			response.WriteHeader(http.StatusBadRequest)
			return
		}
		response.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(ans)
}
func (myserver *Application) EventsForMonthHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" && request.URL.Path == "/events_for_month" {
		myserver.Logger.Println("ошибка обрабочика требуемый метод : GET, требуемый путь /events_for_month. Имеющийся метод", request.Method, "имеющийся путь", request.URL.Path)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	var key string = "abc"
	if request.URL.Query().Has("data") {
		key = request.URL.Query().Get("data")
	}
	myserver.Logger.Println("получаем данные о событии из хранилища за месяц")
	ans, err := myserver.Data.GetAllEventInDataMonth(key)
	if err != nil {
		myserver.Logger.Println("Некорретое поведение бизнесс логики при GET запросе к /events_for_month : ошибка формата ", err.Error())
		if err.Error() == "ошибка формата даты" {
			response.WriteHeader(http.StatusBadRequest)
			return
		}
		response.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(ans)
}
