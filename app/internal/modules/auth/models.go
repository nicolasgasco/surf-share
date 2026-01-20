package auth

type User struct {
	ID       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}

type UserCredentials struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
