package dto

type RegistrationUserDTO struct {
	Email    string
	Password string
}

type LoginUserDTO struct {
	Email    string
	Password string
}

type GetUserDTO struct {
	ID uint64
}
