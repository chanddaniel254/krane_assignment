// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Event struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Location  string `json:"location"`
}

type Expense struct {
	ID          string `json:"id"`
	ItemName    string `json:"item_name"`
	Cost        string `json:"cost"`
	Description string `json:"description"`
	Category    string `json:"category"`
	EventID     *Event `json:"event_id"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewEvent struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Location  string `json:"location"`
}

type NewExpense struct {
	ItemName    string `json:"item_name"`
	Cost        string `json:"cost"`
	Description string `json:"description"`
	Category    string `json:"category"`
	EventID     string `json:"event_id"`
}

type NewOrganizer struct {
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
	Role    string `json:"role"`
}

type NewSession struct {
	EventID   string  `json:"event_id"`
	Name      string  `json:"name"`
	StartTime string  `json:"start_time"`
	EndTime   *string `json:"end_time,omitempty"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phoneno  string `json:"phoneno"`
}

type Organizer struct {
	ID      string `json:"id"`
	EventID *Event `json:"event_id"`
	UserID  *User  `json:"user_id"`
	Role    string `json:"role"`
}

type ScheduleEvent struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Location  string `json:"location"`
	EventID   string `json:"event_id"`
}

type ScheduleSession struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Name      string `json:"name"`
	ID        string `json:"id"`
}

type Session struct {
	ID        string  `json:"id"`
	EventID   *Event  `json:"event_id"`
	Name      string  `json:"name"`
	StartTime string  `json:"start_time"`
	EndTime   *string `json:"end_time,omitempty"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phoneno  string `json:"phoneno"`
}
