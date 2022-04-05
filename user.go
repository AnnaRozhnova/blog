package blog

type User struct {
	Username string `json:"username" db:"username" binding:"required"`
	//Email    string `json:"email" db:"email" binding:"required"`
	Name string `json:"name" db:"name" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}


