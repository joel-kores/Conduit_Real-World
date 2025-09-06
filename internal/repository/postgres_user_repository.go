package repository

import (
	"time"

	"github.com/joel-kores/Conduit_Real-World/internal/domain/user"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) CreateUser(person *user.User) error {
	result := r.db.Create(person)
	if result.Error != nil {
		if isDuplicatedKeyError(result.Error) {
			return user.ErrUserAlreadyExists
		}
	}
	return nil
}

func (r *PostgresUserRepository) GetUser(username string) (*user.User, error) {
	var u user.User
	result := r.db.Where("username=?", username).First(&u)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, user.ErrUserNotFound
		}
		return nil, result.Error
	}
	return &u, nil
}

func (r *PostgresUserRepository) Update(u *user.User) error {
	u.UpdatedAt = time.Now()
	result := r.db.Save(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PostgresUserRepository) Delete (id string) error {
	result := r.db.Where("id = ?", id).Delete(&user.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return user.ErrUserNotFound
	}
	return nil
}

func isDuplicatedKeyError(err error) bool {
	return err != nil && (err.Error() == "pq: duplicate key value violates unique constraint" ||
		err.Error() == "ERROR: duplicate key value violates unique constraint")
}
