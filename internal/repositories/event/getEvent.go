package eventRepository

import (
	"encoding/json"
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
)

func (r *EventRepository) GetEvent(id int) (entity.Event, entity.AdditionalDatas, error) {
	var event entity.Event;
	var aditionalDataJson string;

	query := `
		SELECT
		    e.id,
		    e.name,
		    e.description,
		    e.place,
		    e.date,
		    e.count,
		    e.fundraising,
		    e.user_id,
		    COALESCE(
		        json_agg(
		            json_build_object(
		                'id', ad.id,
		                'key', ad.key,
		                'value', ad.value,
		                'event_id', ad.event_id
		            )
		        ) FILTER (WHERE ad.id IS NOT NULL), '[]'
		    ) AS additional_data
		FROM events e
		LEFT JOIN additional_datas ad ON ad.event_id = e.id
		WHERE e.id = $1
		GROUP BY e.id;
		`;
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
		&aditionalDataJson,
	); err != nil {
		return event, entity.AdditionalDatas{}, fmt.Errorf("EventRepository:GetEvent:scan: %s", err.Error());
	}

	var additionalDatas entity.AdditionalDatas;
	if err := json.Unmarshal([]byte(aditionalDataJson), &additionalDatas); err != nil {
		return event, entity.AdditionalDatas{}, fmt.Errorf("EventRepository:GetEvent:unmarshal: %s", err.Error());
	}

	return event, additionalDatas, nil;
}
