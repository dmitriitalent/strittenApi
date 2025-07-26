package eventRepository

import (
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
)

func (r *EventRepository) CreateEvent(event entity.Event) (entity.Event, error) {
	var createdEvent entity.Event;

	query := fmt.Sprintf(`INSERT INTO %s (
		name,
		description,
		place,
		date,
		count,
		fundraising,
		user_id
	) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING
		id,
		name,
		description,
		place,
		date,
		count,
		fundraising,
		user_id
	`, eventsTable);

	row := r.db.QueryRow(query, 
		event.Name,
		event.Description,
		event.Place,
		event.Date,
		event.Count,
		event.Fundraising,
		event.UserId,
	)
	
	if err := row.Scan(
		&createdEvent.Id,
		&createdEvent.Name,
		&createdEvent.Description,
		&createdEvent.Place,
		&createdEvent.Date,
		&createdEvent.Count,
		&createdEvent.Fundraising,
		&createdEvent.UserId,
	); err != nil {
		return event, err;
	}

	return createdEvent, nil;
}
