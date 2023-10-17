package service

import (
	"errors"
	"event_management/database"
	"event_management/graph/model"
	"fmt"
)

func GetEvents(userId string) ([]*model.Event, error) {
	db := database.Db
	rows, err := db.Query(`	SELECT event_id FROM organizer where user_id = $1 `, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []*model.Event
	for rows.Next() {
		var eventId string
		err = rows.Scan(&eventId)
		if err != nil {
			return nil, err
		}
		rows, err = db.Query(`SELECT id ,name, start_date, end_date ,location FROM events where id = $1`, eventId)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {

			var event model.Event
			err = rows.Scan(&event.ID, &event.Name, &event.StartDate, &event.EndDate, &event.Location)
			events = append(events, &event)
		}

	}

	return events, nil
}

func CreateEvent(name, startDate, endDate, location, adminId string) (*model.Event, error) {
	db := database.Db
	var eventId int
	err := db.QueryRow("Insert into \"events\" (name,start_date,end_date,location) values ($1,$2,$3,$4) returning id ", name, startDate, endDate, location).Scan(&eventId)

	if err != nil {
		return nil, err
	}

	_, err = db.Query("Insert into \"organizer\" (event_id,user_id,role) values ($1,$2,$3) ", eventId, adminId, "admin")

	if err != nil {
		_, err = db.Query("Delete from \"events\" where id = $1", eventId)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &model.Event{
		ID:        fmt.Sprint(eventId),
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Location:  location,
	}, nil
}

func EditEventSchedule(startDate, endDate, newLocation, eventId string) (*model.Event, error) {
	db := database.Db
	var eventName string
	err := db.QueryRow(`UPDATE events
    SET  start_date = $1, end_date = $2, location = $3
    WHERE id = $4 RETURNING name
`, startDate, endDate, newLocation, eventId).Scan(&eventName)

	if err != nil {
		return nil, err
	}
	return &model.Event{
		Name:      eventName,
		StartDate: startDate,
		EndDate:   endDate,
		Location:  newLocation,
		ID:        eventId,
	}, nil

}

func IsUserRelatedToEvent(userId string, eventId string) (bool, error) {
	db := database.Db

	rows, err := db.Query("select id from organizer where user_id = $1 AND event_id = $2", userId, eventId)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if !rows.Next() {
		return false, errors.New("you cant access this event")
	}

	return true, nil
}
