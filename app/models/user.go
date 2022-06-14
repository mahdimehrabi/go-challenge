package models

import (
	"encoding/json"
	"io"
	"time"
)

type User struct {
	ID            string    `json:"ID" binding:"required"`
	Segment       string    `json:"segment" binding:"required"`
	ExpireSegment time.Time `json:"expireSegment" binding:"required"`
}

func (u *User) SliceToModel(data []interface{}) error {
	u.ID = data[0].(string)
	u.Segment = data[1].(string)
	u.ExpireSegment = data[2].(time.Time)
	return nil
}

type Segment struct {
	Segment    string `json:"segmentTitle"`
	UsersCount int64  `json:"usersCount"`
}

func (p *Segment) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
