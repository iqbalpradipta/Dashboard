package login

type Login struct {
	Id       uint
	Email    string
	Role     string
	Password string
}

type UsecaseLogin interface {
	Login(userData Login) (token string, err error)
}

type UserInterface interface {
	SelectUserByEmail(email string) (loginEntity Login, err error)
}
