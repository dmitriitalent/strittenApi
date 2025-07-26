package entity

type AdditionalDatas []AdditionalData
type AdditionalData struct {
	Id      int
	Key     string
	Value   string
	EventId int
}