package models

import (
	"time"
)

type Session struct {
	Id int
	Uuid string
	Email string
	UserId int
	CreatedAt time.Time
}

func (session *Session) Check() (valid bool, err error) {
	sql := "select id, uuid, email, user_id, created_at from sessions where uuid = ?"
	err = Db.QueryRow(sql, session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (session *Session) DeleteByUUID() (err error) {
	sql := "delete from sessions where uuid = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

func (session *Session) User() (user User, err error) {
	user = User{}
	sql := "SELECT id, uuid, name, email, created_at from users where id = ?"
	err = Db.QueryRow(sql, session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func SessionDeleteAll() (err error) {
	sql := "delete from sessions"
	_, err = Db.Exec(sql)
	return
}