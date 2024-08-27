package model

import "time"

type Conference struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  int       `json:"category_id"`
	Theme       string    `json:"theme"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CityID      int       `json:"city_id"`
	Website     string    `json:"website"`
	OrganizerID int       `json:"organizer_id"`
}
