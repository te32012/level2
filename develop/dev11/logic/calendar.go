package dev11

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

type Database struct {
	Data map[int]Event
}

type Event struct {
	Id           int
	Date         string
	Notification string
}

func (note *Event) Event(id int, date string, notification string) {
	note.Id = id
	note.Date = date
	note.Notification = notification
}

func (data *Database) AddEventFromJson(read io.ReadCloser) error {
	var event Event = Event{}
	buffer, err := io.ReadAll(read)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buffer, &event)
	if err != nil {
		return errors.New("ошибка формата данных")

	}
	data.Data[event.Id] = event
	return nil
}

func (data *Database) RemoveEventFromJson(read io.ReadCloser) error {
	var event Event = Event{}
	buffer, err := io.ReadAll(read)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buffer, &event)
	if err != nil {
		return errors.New("ошибка формата данных")
	}
	_, ok := data.Data[event.Id]
	if ok {
		delete(data.Data, event.Id)
	} else {
		return errors.New("ключа не существует")
	}
	return nil
}

func (data *Database) UpdateEventFromJson(read io.ReadCloser) error {
	var event Event = Event{}
	buffer, err := io.ReadAll(read)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buffer, &event)
	if err != nil {
		return err
	}
	value, ok := data.Data[event.Id]
	if ok {
		data.Data[event.Id] = value
	} else {
		return errors.New("такого события не существует")
	}
	return nil
}

func (data *Database) GetAllEventInDataDay(day string) ([]byte, error) {
	t, err := time.Parse("02-01-2006", day)
	if err != nil {
		return nil, errors.New("ошибка формата даты")
	}
	events := make([]Event, 0)
	for _, value := range data.Data {
		t1, err := time.Parse("02-01-2006", value.Date)
		if err != nil {
			return nil, errors.New("ошибка формата даты")
		}
		if t.Day() == t1.Day() && t.Year() == t1.Year() && t.Month() == t1.Month() {
			events = append(events, value)
		}
	}
	ts := struct {
		Events []Event `json:"result"`
	}{Events: events}
	tmp, err := json.Marshal(ts)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

func (data *Database) GetAllEventInDataWeek(day string) ([]byte, error) {
	t, err := time.Parse("02-01-2006", day)
	if err != nil {
		return nil, err
	}
	dw := int(t.Weekday())
	d := time.Hour * 24
	week := make([]time.Time, 0)
	switch {
	case dw == 0:
		week = append(week, t)          // пн
		week = append(week, t.Add(d))   // вт
		week = append(week, t.Add(d*2)) // ср
		week = append(week, t.Add(d*3)) // чт
		week = append(week, t.Add(d*4)) // пт
		week = append(week, t.Add(d*5)) // сб
		week = append(week, t.Add(d*6)) // вс
	case dw == 1:
		week = append(week, t)          // вт
		week = append(week, t.Add(-d))  // пн
		week = append(week, t.Add(d*2)) // ср
		week = append(week, t.Add(d*3)) // чт
		week = append(week, t.Add(d*4)) // пт
		week = append(week, t.Add(d*5)) // сб
		week = append(week, t.Add(d*6)) // вс
	case dw == 2:
		week = append(week, t)           // ср
		week = append(week, t.Add(-d))   // пн
		week = append(week, t.Add(-d*2)) // вт
		week = append(week, t.Add(d*3))  // чт
		week = append(week, t.Add(d*4))  //пт
		week = append(week, t.Add(d*5))  // сб
		week = append(week, t.Add(d*6))  // вс
	case dw == 3:
		week = append(week, t)           // чт
		week = append(week, t.Add(-d))   // пн
		week = append(week, t.Add(-d*2)) //вт
		week = append(week, t.Add(-d*3)) // ср
		week = append(week, t.Add(d*4))  // пт
		week = append(week, t.Add(d*5))  // сб
		week = append(week, t.Add(d*6))  // вск
	case dw == 4:
		week = append(week, t)           // пт
		week = append(week, t.Add(-d))   // пн
		week = append(week, t.Add(-d*2)) // вт
		week = append(week, t.Add(-d*3)) // ср
		week = append(week, t.Add(-d*4)) // чт
		week = append(week, t.Add(d*5))  // сб
		week = append(week, t.Add(d*6))  // вс
	case dw == 5:
		week = append(week, t)           // сб
		week = append(week, t.Add(-d))   //пн
		week = append(week, t.Add(-d*2)) //вт
		week = append(week, t.Add(-d*3)) // ср
		week = append(week, t.Add(-d*4)) // чт
		week = append(week, t.Add(-d*5)) // пт
		week = append(week, t.Add(d*6))  // вс
	case dw == 6:
		week = append(week, t)           // вск
		week = append(week, t.Add(-d))   // пн
		week = append(week, t.Add(-d*2)) // вт
		week = append(week, t.Add(-d*3)) // ср
		week = append(week, t.Add(-d*4)) // чт
		week = append(week, t.Add(-d*5)) // пт
		week = append(week, t.Add(-d*6)) // сб
	}
	var ans []Event = make([]Event, 0)
	for _, v := range data.Data {
		for _, d := range week {
			t1, err := time.Parse("02-01-2006", v.Date)
			if err != nil {
				return nil, errors.New("ошибка формата даты")
			}
			if d.Day() == t1.Day() && d.Year() == t1.Year() && d.Month() == t1.Month() {
				ans = append(ans, v)
			}
		}
	}
	ts := struct {
		Events []Event `json:"result"`
	}{Events: ans}
	tmp, err := json.Marshal(ts)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

func (data *Database) GetAllEventInDataMonth(day string) ([]byte, error) {
	t, err := time.Parse("02-01-2006", day)
	if err != nil {
		return nil, err
	}
	events := make([]Event, 0)
	for _, value := range data.Data {
		t1, err := time.Parse("02-01-2006", value.Date)
		if err != nil {
			return nil, errors.New("ошибка формата даты")
		}
		if t1.Month() == t.Month() {
			events = append(events, value)
		}
	}
	ts := struct {
		Events []Event `json:"result"`
	}{Events: events}
	tmp, err := json.Marshal(ts)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}
