package comand

type CreateUserCommand struct {
	Name     string
	Password string
	IsAdmin  bool
}
