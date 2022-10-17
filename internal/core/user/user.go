package user

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	PassHash string `json:"password"`
}

// NewUser
// is used for creating new user record in the database
type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUser
// not sure about internals yet, but it is used for updating user data
type UpdateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
