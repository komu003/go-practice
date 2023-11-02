package models

type Micropost struct {
	ID        int       `db:"id"`
	Content   string    `db:"content"`
	UserID    int       `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
