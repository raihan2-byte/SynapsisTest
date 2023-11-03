package user

// import "time"

type RegisterFormatter struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name" `
	Email    string `json:"email" `
	Password string `json:"password" `
	Role     int    `json:"role" `
	Balance  int    `json:"balance" `
}

func FormatterRegister(user User) RegisterFormatter {
	formatterRegister := RegisterFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Balance:  user.Balance,
	}
	return formatterRegister

}

type UserFormatter struct {
	Token string `json:"token"`
}

func FormatterUser(Token string) UserFormatter {
	formatterLogin := UserFormatter{
		Token: Token,
	}
	return formatterLogin
}

// type DeletedUserFormatter struct {
// 	Message string `json:"message"`
// }

// func FormatterDeletedUser(user string) DeletedUserFormatter {
// 	formatterDeletedUser := DeletedUserFormatter{
// 		Message: user,
// 	}
// 	return formatterDeletedUser
// }
