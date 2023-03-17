package main

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
