package sets

import (
	"../util"
	"fmt"
	"time"
)

const (
	insertSetQuery = `INSERT INTO sets(completed_at, exercise, reps, user_id, weight, unit, notes) 
		VALUES ($1, $2, $3, $4, $5, $6, $7);`
	updateSetQuery = `UPDATE sets SET 
			completed_at=$1, exercise = $2, reps=$3, user_id=$4, weight=$5, unit=$6, notes=$7
			WHERE id=$8;`
	selectSetQuery = `SELECT id, completed_at, exercise, reps, user_id, weight, unit, notes FROM
		sets WHERE user_id = $1;`
)

type Set struct {
	Id          int64
	CompletedAt time.Time
	Exercise    string
	Reps        int
	UserId      int64
	Weight      int
	Unit        string
	Notes       string
}

func New(exercise string, reps int, userid int64, weight int, unit string) *Set {
	// Database time is in milliseconds, truncate time to be precise only to the second 
	now := time.Now().UTC()
	return &Set{
		CompletedAt: time.Unix(now.Unix(), 0),
		Exercise:    exercise,
		Reps:        reps,
		UserId:      userid,
		Weight:      weight,
		Unit:        unit,
		Notes:       ""}
}

func (s *Set) Save(db util.DB) error {
	var err error
	if s.Id > 0 {
		_, err = db.Exec(updateSetQuery,
			s.CompletedAt,
			s.Exercise,
			s.Reps,
			s.UserId,
			s.Weight,
			s.Unit,
			s.Notes,
			s.Id)
	} else {
		_, err = db.Exec(insertSetQuery,
			s.CompletedAt,
			s.Exercise,
			s.Reps,
			s.UserId,
			s.Weight,
			s.Unit,
			s.Notes)
		if err != nil {
			return fmt.Errorf("error while inserting set to database: %v", err)
		}
		id, err := util.GetLastId(db, "sets")
		s.Id = id
	}
	return err
}

func GetByUserId(user_id int64, db util.DB) ([]*Set, error) {
	rows, err := db.Query(selectSetQuery, user_id)
	if err != nil {
		return nil, fmt.Errorf("error while querying sets from database: %v", err)
	}
	sets := make([]*Set, 0)
	for rows.Next() {
		set := new(Set)
		err := rows.Scan(&set.Id, &set.CompletedAt, &set.Exercise, &set.Reps, &set.UserId, &set.Weight, &set.Unit, &set.Notes)
		if err != nil {
			return nil, fmt.Errorf("error while scanning set from database rows: %v", err)
		}
		sets = append(sets, set)
	}
	return sets, nil
}
