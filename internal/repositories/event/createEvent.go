package eventRepository

import (
	"fmt"
	"strings"

	"github.com/dmitriitalent/strittenApi/internal/entity"
	additionalDataRespository "github.com/dmitriitalent/strittenApi/internal/repositories/additionalDatas"
)

func (r *EventRepository) CreateEvent(event entity.Event, additionalDatas entity.AdditionalDatas) (entity.Event, entity.AdditionalDatas, error) {
	var createdEvent entity.Event;

	tx, err := r.db.Begin()
	if err != nil {
		return event, additionalDatas, fmt.Errorf("EventRepository:CreateEvent: Failed to begin transaction: %s", err.Error());
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
		return event, additionalDatas, err;
	}

	placeholders := []string{}
	values := []interface{}{}
	i := 1
	for _, additionalData := range additionalDatas {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d)", i, i+1, i+2))
		values = append(values, additionalData.Key, additionalData.Value, createdEvent.Id)
		i += 3
	}

	query = fmt.Sprintf("INSERT INTO %s (key, value, event_id) VALUES %s RETURNING id, key, value, event_id", additionalDataRespository.AdditionalDatasTable, strings.Join(placeholders, ","))
	rows, err := tx.Query(query, values...); 
	if err != nil {
		return event, additionalDatas, fmt.Errorf("EventRepository:CreateEvent: Failed to add additionalDatas, %s, %s: %s", query, values, err.Error());
	}
	if rows == nil {
		tx.Commit()
		return createdEvent, additionalDatas, nil
	}

	var createdAdditionalDatas entity.AdditionalDatas;
	for rows.Next() {
		var additionalData entity.AdditionalData;
		if err := rows.Scan(
			&additionalData.Id,
			&additionalData.Key,
			&additionalData.Value,
			&additionalData.EventId,
		); err != nil {
			return event, additionalDatas, fmt.Errorf("EventRepository:CreateEvent: Failed to scan additionalDatas, %s, %s: %s", query, values, err.Error());
		}
		createdAdditionalDatas = append(createdAdditionalDatas, additionalData)
	}

	err = tx.Commit()
	if err != nil {
		return event, additionalDatas, fmt.Errorf("EventRepository:CreateEvent: Failed to commit transaction: %s", err.Error());
	}

	return createdEvent, createdAdditionalDatas, nil;
}
