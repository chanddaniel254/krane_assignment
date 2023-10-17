package auth

import (
	"errors"
	"event_management/database"
)

func IsAdmin(userId string, eventId string, extendAuth bool, haveOrganizerId bool) (bool, error) {
	db := database.Db

	finalEventId := eventId

	if haveOrganizerId {
		rows, err := db.Query("SELECT event_id FROM organizer WHERE id=$1", finalEventId)
		if err != nil {
			return false, err
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&finalEventId)
			if err != nil {
				return false, err

			}
		}
	}

	rows, err := db.Query("select role from organizer where user_id = $1 AND event_id  = $2 ", userId, finalEventId)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	for rows.Next() {

		var role string
		err = rows.Scan(&role)

		if err != nil {
			return false, nil
		}
		if extendAuth {
			if role == "admin" || role == "contributor" {
				return true, nil
			}
		}
		if role == "admin" {
			return true, nil
		}
	}
	return false, errors.New("Unauthorized access --")

}
