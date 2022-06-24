package models

import "time"

type Post struct {
	Id int
	Uuid string
	Body string
	UserId int
	ThreadId int
	CreatedAt time.Time
}

// 格式化时间
func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("2006-01-02 15:04:05")
}

// 获取写这个主题的用户
func (post *Post) User() (user User) {
	user = User{}
	sql := "select id, uuid, name, email, created_at from users where id = ?"
	Db.QueryRow(sql, post.Id).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}