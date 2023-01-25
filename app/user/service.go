package user

type UserService interface {
	Get(uint) (User, error)
	GetByUsername(string) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
}
