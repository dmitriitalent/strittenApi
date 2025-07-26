package eventRepository

import (
	"fmt"
	"strings"

	"github.com/dmitriitalent/strittenApi/internal/entity"
	additionalDataRespository "github.com/dmitriitalent/strittenApi/internal/repositories/additionalDatas"
)

func (r *EventRepository) CreateEvent(event entity.Event, additionalDatas entity.AdditionalDatas) (entity.Event, error) {
	var createdEvent entity.Event;

	tx, err := r.db.Begin()
	if err != nil {
		return event, fmt.Errorf("EventRepository:CreateEvent: Failed to begin transaction");
	}
	defer tx.Rollback()

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

	row := tx.QueryRow(query, 
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

	placeholders := []string{}
	values := []interface{}{}
	i := 1
	for _, additionalData := range additionalDatas {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d)", i, i+1, i+2))
		values = append(values, additionalData.Key, additionalData.Value, createdEvent.Id)
		i += 3
	}

	query = fmt.Sprintf("INSERT INTO %s (key, value, event_id) VALUES %s", additionalDataRespository.AdditionalDatasTable, strings.Join(placeholders, ","))
	if _, err := tx.Exec(query, values...); err != nil {
		return event, fmt.Errorf("EventRepository:CreateEvent: Failed to add additionalDatas, %s, %s", query, values);
	}

	err = tx.Commit()
	if err != nil {
		return event, fmt.Errorf("EventRepository:CreateEvent: Failed to commit transaction");
	}

	return createdEvent, nil;
}
