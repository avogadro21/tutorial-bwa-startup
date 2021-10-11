package user

type DTORegisterUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
}

func GetDTORegisterUser(user User) DTORegisterUser {
	payload := DTORegisterUser{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Occupation: user.Occupation,
	}

	return payload
}

type DTOLogin struct {
	DTORegisterUser
	Token string `json:"token"`
}

func GetDTOLogin(user User, token string) DTOLogin {
	payload := DTOLogin{
		DTORegisterUser: DTORegisterUser{
			ID:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			Occupation: user.Occupation,
		},
		Token: token,
	}

	return payload
}
