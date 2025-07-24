package dto

type CreateEvent struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Place       string `json:"place"`
	Count       int    `json:"count"`
	Fundraising string `json:"fundraising"`
}
