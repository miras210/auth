package service

func GetLogin(email, username string) string {
	if len(email) == 0 {
		return username
	}
	return email
}
