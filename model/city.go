package model

type City struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	CountryCode string `json:"country_code"`
}
