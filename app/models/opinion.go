package models

import (
	"log"
	"time"
)

type Opinion struct {
	ID        int
	UUID      string
	Name      string
	Opinion   string
	CreatedAt time.Time
}

func (o *Opinion) CreateOpinion() (err error) {
	cmd := `insert into opinion (
		uuid,
		name,
		opinion,
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		o.Name,
		o.Opinion,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
