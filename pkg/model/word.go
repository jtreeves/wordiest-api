package model

type Word struct {
	ID     int    `json:"id"`
	Value  string `json:"value"`
	Length int    `json:"length"`
	Letter string `json:"letter"`
}
