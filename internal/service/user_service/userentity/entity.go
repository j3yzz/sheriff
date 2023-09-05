package userentity

type UserRegisterEntity struct {
	ID       uint
	Phone    string
	Password string
}

type LoginEntity struct {
	Phone    string
	Password string
}
