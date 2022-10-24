package repository

func GetLogin(email, username *string) *string {
	if email == nil {
		return username
	}
	return email
}
