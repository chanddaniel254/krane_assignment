package service

import (
	"errors"
	"event_management/database"
	"event_management/graph/model"
)

func FetchEventIdFromSession(sessionId string) (string, error) {
	db := database.Db

	rows, err := db.Query("select event_id from session where id = $1", sessionId)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var eventId string
	for rows.Next() {
		err = rows.Scan(&eventId)
	}

	return eventId, nil
}

func GetSessionsByEventId(eventId string) ([]*model.Session, error) {
	db := database.Db
	rows, err := db.Query("Select id,name,start_time,end_time from session where event_id = $1 ", eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sessions []*model.Session
	for rows.Next() {
		var session model.Session
		err = rows.Scan(&session.ID, &session.Name, &session.StartTime, &session.EndTime)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, &session)
	}

	return sessions, nil
}

func CreateSessionForEvent(name, start_time, end_time, event_id string) (*model.Session, error) {
	db := database.Db
	var sessionId string

	err := db.QueryRow("Insert into session (name,start_time,end_time,event_id) values ($1,$2,$3,$4) returning id", name, start_time, end_time, event_id).Scan(&sessionId)

	if err != nil {
		return nil, err
	}
	return &model.Session{
		ID:        sessionId,
		Name:      name,
		StartTime: start_time,
		EndTime:   &end_time,
	}, nil

}
func ScheduleSession(name, start_time, end_time, sessionId string) (*model.Session, error) {
	db := database.Db

	result, err := db.Exec(`UPDATE session
    SET  start_time = $1, end_time = $2, name = $3
    WHERE id = $4 RETURNING name`, start_time, end_time, name, sessionId)

	if err != nil {
		return nil, err
	}

	rowsEffected, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	if rowsEffected == 0 {
		return nil, errors.New("the session doesnt exists")
	}
	return &model.Session{
		ID:        sessionId,
		StartTime: start_time,
		EndTime:   &end_time,
		Name:      name,
	}, nil
}
