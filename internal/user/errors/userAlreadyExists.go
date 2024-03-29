package errors

type UserAlreadyExists struct{}

func (e UserAlreadyExists) Error() string {
	return "User already exists in db"
}
