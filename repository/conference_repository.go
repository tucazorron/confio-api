package repository

import (
	"database/sql"
	"fmt"
	"tucazorron/confio-api/model"
)

type ConferenceRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ConferenceRepository {
	return ConferenceRepository{
		connection: connection,
	}
}

func (rep *ConferenceRepository) GetConferences() ([]model.Conference, error) {
	query := "select * from conference"
	rows, err := rep.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Conference{}, err
	}
	var conferenceList []model.Conference
	var conferenceObj model.Conference
	for rows.Next() {
		err = rows.Scan(
			&conferenceObj.ID,
			&conferenceObj.Name,
			&conferenceObj.Description,
			&conferenceObj.CategoryID,
			&conferenceObj.Theme,
			&conferenceObj.StartDate,
			&conferenceObj.EndDate,
			&conferenceObj.CityID,
			&conferenceObj.Website,
			&conferenceObj.OrganizerID,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Conference{}, err
		}
		conferenceList = append(conferenceList, conferenceObj)
	}
	rows.Close()
	return conferenceList, nil
}

func (rep *ConferenceRepository) CreateConference(conference model.Conference) (int, error) {
	var id int
	query, err := rep.connection.Prepare(`
		insert into conference
			(name, description, category_id, theme, start_date, end_date, city_id, website, organizer_id)
		values
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`)
	if err != nil {
		return 0, err
	}
	err = query.QueryRow(
		conference.Name,
		conference.Description,
		conference.CategoryID,
		conference.Theme,
		conference.StartDate,
		conference.EndDate,
		conference.CityID,
		conference.Website,
		conference.OrganizerID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
