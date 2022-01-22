package models

import (
	"errors"
	"time"
)

type Grade struct {
	ID uint `json:"id"`
	Score   uint   `json:"score"`
	Index   string `json:"index"`
	StudentId uint `json:"student_id"`
	SubjectId uint `json:"subject_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (r *Grade) ConvertScore() {
	if r.Score >= 80 {
		r.Index = "A"
	} else if r.Score >= 70 && r.Score < 80 {
		r.Index = "B"
	} else if r.Score >= 60 && r.Score < 70 {
		r.Index = "C"
	} else if r.Score >= 50 && r.Score < 60 {
		r.Index = "D"
	} else if r.Score < 50 {
		r.Index = "E"
	}
}

func (r *Grade) Validate() error {
	if r.Score > 100 {
		return errors.New("grade should no more than 100")
	}
	return nil
}