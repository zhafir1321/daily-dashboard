package repositories

import (
	"backend/internal/app/database"
	"backend/internal/app/entities"
	"database/sql"
)

type Event struct {
	database.BaseSQLRepository[entities.Event]
}

func NewEventRepository(db *sql.DB) *Event {
	return &Event{
		BaseSQLRepository: database.BaseSQLRepository[entities.Event]{DB: db},
	}
}

func mapEvent(rows *sql.Row, e *entities.Event) error {
	return rows.Scan(&e.ID, &e.Title, &e.Description, &e.EventDate, &e.EventTime, &e.UserID, &e.CreatedAt, &e.UpdatedAt)
}

func mapEvents(rows *sql.Rows, e *entities.Event) error {
	return rows.Scan(&e.ID, &e.Title, &e.Description, &e.EventDate, &e.EventTime, &e.UserID, &e.CreatedAt, &e.UpdatedAt)
}

func (r *Event) FindByID(id int, userID int) (*entities.Event, error) {
	return r.SelectSingle(
		mapEvent,
		"SELECT id, title, description, TO_CHAR(event_date, 'YYYY-MM-DD') as event_date, TO_CHAR(event_time, 'HH24:MI') as event_time, user_id, created_at, updated_at FROM events WHERE id = $1 AND user_id = $2",
		id,
		userID,
	)
}

func (r *Event) GetAllEvents(userID int) ([]*entities.Event, error) {
	return r.SelectMultiple(
		mapEvents,
		"SELECT id, title, description, TO_CHAR(event_date, 'YYYY-MM-DD') as event_date, TO_CHAR(event_time, 'HH24:MI') as event_time, user_id, created_at, updated_at FROM events WHERE user_id = $1",
		userID,
	)
}

func (r *Event) CreateEvent(event *entities.Event) error {
	id, err := r.Insert(
		"INSERT INTO events (title, description, event_date, event_time, user_id) VALUES ($1, $2, $3, $4, $5)",
		event.Title,
		event.Description,
		event.EventDate,
		event.EventTime,
		event.UserID,
	)
	event.ID = id
	return err
}

func (r *Event) UpdateEvent(event *entities.Event) error {
	_, err := r.ExecuteQuery(
		"UPDATE events SET title = $1, description = $2, event_date = $3, event_time = $4, updated_at = NOW() WHERE id = $5 AND user_id = $6",
		event.Title,
		event.Description,
		event.EventDate,
		event.EventTime,
		event.ID,
		event.UserID,
	)
	return err
}

func (r *Event) DeleteEvent(id int, userID int) error {
	_, err := r.ExecuteQuery(
		"DELETE FROM events WHERE id = $1 AND user_id = $2",
		id,
		userID,
	)
	return err
}
