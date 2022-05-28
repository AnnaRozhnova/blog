package blog

type Post struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title" binding:"required"`
	Body     string `json:"body" db:"body" binding:"required"`
	Username string `json:"username" db:"username"`
}

type Comment struct {
	Id       int    `json:"-" db:"id"`
	Body     string `json:"body" db:"body" binding:"required"`
	Username string `json:"username" db:"username"`
	PostId   int    `json:"post_id" db:"post_id" binding:"required"`
}