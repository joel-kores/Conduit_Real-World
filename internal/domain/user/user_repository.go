package user


type UserRepository interface {
	Create(user *User) error
	FindByID(id string) (*User, error)
	FindByEmail(email string) (User, error)
	Update(user User) error
	Delete(id string) error
	ExistsByEmail(email string) (bool, error)
	FindByRefreshToken(token string) (User, error)
}
