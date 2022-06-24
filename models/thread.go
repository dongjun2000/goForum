package models

import (
	"time"
)

type Thread struct {
	Id int
	Uuid string
	Topic string
	UserId int
	CreatedAt time.Time
}

// 格式化时间
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("2006-01-02 15:04:05")
}

// 获取群组中的主题数
func (thread *Thread) NumReplies() (count int) {
	sql := "select count(*) from posts where thread_id=?"
	rows, err := Db.Query(sql, thread.Id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// 获取群组里所有的主题
func (thread *Thread) Posts() (posts []Post, err error) {
	sql := "select id, uuid, body, user_id, thread_id, created_at from posts where thread_id = ?"
	rows, err := Db.Query(sql, thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// 获取所有的群组
func Threads() (threads []Thread, err error) {
	sql := "select id, uuid, topic, user_id, created_at from threads order by created_at desc"
	rows, err := Db.Query(sql)
	if err != nil {
		return
	}
	for rows.Next() {
		thread := Thread{}
		if err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt); err != nil {
			return
		}
		threads = append(threads, thread)
	}
	rows.Close()
	return
}

// 通过uuid获取群组
func ThreadByUUID(uuid string) (thread Thread, err error) {
	thread = Thread{}
	sql := "select id, uuid, topic, user_id, created_at from threads where uuid = ?"
	err = Db.QueryRow(sql, uuid).
		Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	return
}

// 获取创建这个群组的用户
func (thread *Thread) User() (user User) {
	user = User{}
	sql := "select id, uuid, name, email, created_at from users where id = ?"
	Db.QueryRow(sql, thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}