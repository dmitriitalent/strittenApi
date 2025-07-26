package entity

type AdditionalDatas []AdditionalData
type AdditionalData struct {
	Id      int    `json:"id"`
	Key     string `json:"key"`
	Value   string `json:"value"`
	EventId int    `json:"event_id"`
}