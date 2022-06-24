package models

import "time"

type User struct {
	Id int
	Uuid string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

// 为现有用户创建一个新的session
func (user *User) CreateSession() (session Session, err error) {
	sql := "insert into sessions(uuid, email, user_id, created_at) values(?, ?, ?, ?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()

	stmt.Exec(uuid, user.Email, user.Id, time.Now())

	sql = "select id, uuid, email, user_id, created_at from sessions where uuid = ?"
	stmt, err = Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.CreatedAt)
	return
}

func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("select id, uuid, email, user_id, created_at from sessions where user_id = ?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func (user *User) Create() (err error) {
	sql := "insert into users(uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	stmt.Exec(uuid, user.Name, user.Email, user.Password, time.Now())

	sql = "select id, uuid, created_at from users where uuid = ?"
	stmt, err = Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(uuid).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

func (user *User) Delete() (err error) {
	sql := "delete from users where id = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

func (user *User) Update() (err error) {
	sql := "update users set name = ?, email = ? where id = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Id)
	return
}

func UserDeleteAll() (err error) {
	sql := "delete from users"
	_, err = Db.Exec(sql)
	return
}

func Users() (users []User, err error) {
	rows, err := Db.Query("select id, uuid, name, email, password, created_at from users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid, name, email, password, created_at from users where uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// 创建一个群组
func (user *User) CreateThread(topic string) (thread Thread, err error) {
	sql := "insert into threads(uuid, topic, user_id, created_at) values(?, ?, ?, ?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	stmt.Exec(uuid, topic, user.Id, time.Now())

	sql = "select id, uuid, topic, user_id, created_at from threads where uuid = ?"
	stmt, err = Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(uuid).
		Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	return
}

// 创建一个帖子
func (user *User) CreatePost(thread Thread, body string) (post Post, err error) {
	sql := "insert into posts(uuid, body, user_id, thread_id, created_at) values(?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	stmt.Exec(uuid, body, user.Id, thread.Id, time.Now())

	sql = "select id, uuid, body, user_id, thread_id, created_at from posts where uuid = ?"
	stmt, err = Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(uuid).
		Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}