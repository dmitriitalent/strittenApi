package entity

type RefreshToken struct {
	Id           int    `json:"-"`
	RefreshToken string `json:"refreshToken"`
	UserId       int    `json:"userId"`
}
