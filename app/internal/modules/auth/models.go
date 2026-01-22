package auth

type User struct {
	ID       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}

type UserWithPassword struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type AuthResult struct {
	User  *User
	Token string
}
