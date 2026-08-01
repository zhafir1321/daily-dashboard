package entities

type Event struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UserID      int     `json:"user_id"`
	EventDate   string  `json:"event_date"`
	EventTime   *string `json:"event_time"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
