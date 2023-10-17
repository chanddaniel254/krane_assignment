package service

import (
	"errors"
	"event_management/database"
	"event_management/graph/model"
	"fmt"
)

func CreateOrganizer(eventId, userId, role string) (*model.Organizer, error) {
	db := database.Db
	var organizerId int
	rows, err := db.Query("select id from organizer where user_id = $1 AND event_id = $2", userId, eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userExist bool
	for rows.Next() {
		userExist = true
	}
	if userExist {
		return nil, errors.New("user has already been assigned role")
	}
	err = db.QueryRow("insert into \"organizer\" (event_id, user_id,role) values ($1,$2,$3) returning 1", eventId, userId, role).Scan(&organizerId)
	if err != nil {
		return nil, err
	}

	return &model.Organizer{
		ID:      fmt.Sprint(organizerId),
		EventID: &model.Event{ID: eventId},
		Role:    role,
		UserID:  &model.User{ID: userId},
	}, nil
}

func GetOrganizersByEventId(eventId string) ([]*model.Organizer, error) {
	db := database.Db
	rows, err := db.Query(`
        SELECT
            o.id AS organizer_id,
            o.role AS role,
            e.name AS event_name,
            e.start_date AS event_start_date,
            e.end_date AS event_end_date,
            e.location AS event_location,
			e.id AS event_id,
			u.name AS user_name,
			u.email AS user_email,
			u.phone AS user_phone,
			u.id AS user_id
        FROM
            organizer o 
        JOIN
            events e
        ON
            o.event_id = e.id
		JOIN 
		    "User" u
		ON 
		   o.user_id = u.id
		WHERE
		   o.event_id = $1
			`, eventId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var organizers []*model.Organizer
	for rows.Next() {
		var organizer model.Organizer
		var event model.Event // Create a new event instance
		var user model.User
		err = rows.Scan(&organizer.ID, &organizer.Role, &event.Name, &event.StartDate, &event.EndDate, &event.Location, &event.ID, &user.Name, &user.Email, &user.Phoneno, &user.ID)

		if err != nil {
			return nil, err
		}
		organizer.UserID = &user
		organizer.EventID = &event
		organizers = append(organizers, &organizer)
	}
	return organizers, nil

}

func GetOrganizerById(organizerId string) (*model.Organizer, error) {
	db := database.Db
	rows, err := db.Query(`
        SELECT
            o.id AS organizer_id,
            o.role AS role,
            e.name AS event_name,
            e.start_date AS event_start_date,
            e.end_date AS event_end_date,
            e.location AS event_location,
			e.id AS event_id,
			u.name AS user_name,
			u.email AS user_email,
			u.phone AS user_phone,
			u.id AS user_id
        FROM
            organizer o 
        JOIN
            events e
        ON
            o.event_id = e.id
		JOIN 
		    "User" u
		ON 
		   o.user_id = u.id
		WHERE
		   o.id = $1
			`, organizerId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var organizer model.Organizer
	for rows.Next() {
		var event model.Event // Create a new event instance
		var user model.User
		err = rows.Scan(&organizer.ID, &organizer.Role, &event.Name, &event.StartDate, &event.EndDate, &event.Location, &event.ID, &user.Name, &user.Email, &user.Phoneno, &user.ID)

		if err != nil {
			return nil, err
		}
		organizer.UserID = &user
		organizer.EventID = &event

	}
	return &organizer, nil
}

func RemoveOrganizer(organizerId string) (string, error) {
	db := database.Db
	result, err := db.Exec("DELETE FROM organizer WHERE id = $1", organizerId)

	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowsAffected == 0 {
		return "", errors.New("organizer doesnt exists")
	}

	return "success", nil

}
