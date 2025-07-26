package eventRepository

import (
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
)

func (r *EventRepository) GetEvent(id int) (entity.Event, error) {
	var event entity.Event;

	query := fmt.Sprintf("SELECT id, name, description, place, date, count, fundraising, user_id FROM %s WHERE id = $1", eventsTable);
	row := r.db.QueryRow(query, id)
	if err := row.Scan(
		&event.Id,
		&event.Name,
		&event.Description,
		&event.Place,
		&event.Date,
		&event.Count,
		&event.Fundraising,
		&event.UserId,
	); err != nil {
		return event, err;
	}

	return event, nil;
}
