package rack

import (
	"bike-api/model"
	"fmt"
)

type Rack struct {
	Id          int     `json:"id"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
}

func GetAll() ([]Rack, error) {
	conn, err := model.ConnectToDB()
	defer model.SafeClose(conn, &err)

	rows, err := conn.Query("call spRackGetAll()")
	if err != nil {
		return nil, err
	}
	defer model.SafeClose(rows, &err)

	var r Rack
	var racks []Rack
	for rows.Next() {
		err = rows.Scan(
			&r.Id,
			&r.Latitude,
			&r.Longitude,
			&r.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows from database: %v", err)
		}
		racks = append(racks, r)
	}
	return racks, nil
}
