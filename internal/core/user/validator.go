package user

func (u NewUser) Validate() map[string]string {
	errors := make(map[string]string)

	if len(u.Password) == 0 {
		errors["password"] = "invalid password, password can not be empty"
	}

	if len(u.Username) == 0 && len(u.Email) == 0 {
		errors["login"] = "invalid username or email, they can not be empty"
	}

	if len(errors) == 0 {
		return nil
	}

	return errors
}

func (u UpdateUser) Validate() map[string]string {
	errors := make(map[string]string)

	if len(u.Username) == 0 && len(u.Email) == 0 {
		errors["login"] = "invalid username or email, they can not be empty"
	}

	if len(errors) == 0 {
		return nil
	}

	return errors
}
