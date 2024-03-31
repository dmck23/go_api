package errors

type EmailAlreadyInUse struct{}

func (e EmailAlreadyInUse) Error() string {
	return "Email already registered to user"
}
