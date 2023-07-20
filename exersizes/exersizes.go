package exercises

import (
	"../util"
)

const (
	saveExerciseQuery     = "INSERT INTO exercises (name) VALUES ($1);"
	getAllExercisesQuery  = "SELECT name FROM exercises;"
)

type Exercise struct {
	Name string `json:"name"`
}

func (e *Exercise) Save(db util.DB) error {
	_, err := db.Exec(saveExerciseQuery, e.Name)
	return err
}

func GetAll(db util.DB) ([]*Exercise, error) {
	rows, err := db.Query(getAllExercisesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	exercises := make([]*Exercise, 0)
	for rows.Next() {
		exercise := &Exercise{}
		err := rows.Scan(&exercise.Name)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	// Check if rows.Next() stopped due to an error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exercises, nil
}
