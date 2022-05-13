package blog

type User struct {
	Username string `json:"username" db:"username" binding:"required"`
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}
