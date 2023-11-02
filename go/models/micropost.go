package models

type Micropost struct {
	ID        int       `db:"id"`
	Content   string    `db:"content"`
	UserID    int       `db:"user_id"`
}
