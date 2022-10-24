package user

type User struct {
	ID       string `json:"id"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	PassHash string `json:"password"`
}

type SignIn struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

// NewUser
// is used for creating new user record in the database
type NewUser struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

// UpdateUser
// not sure about internals yet, but it is used for updating user data
type UpdateUser struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type Model struct {
	ID       *string
	Login    *string
	PassHash *string
}
