package user

type UserRepository interface {
	CreateUser(user *User) error
	FindUserByID(id string) (*User, error)
	FindUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
	UserExistsByEmail(email string) (bool, error)
}
