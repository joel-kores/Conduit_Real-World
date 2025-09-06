package repository

import (
	"sync"
	"time"

	"github.com/joel-kores/Conduit_Real-World/internal/domain/user"
)

type InMemoryUserRepository struct {
	users map[string]*user.User
	mu sync.RWMutex
}

func NewInMemoryUserRespository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*user.User),
	}
}

func (r *InMemoryUserRepository) Create(u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[u.Email]; exists {
		return user.ErrUserAlreadyExists
	}

	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	r.users[u.Email] = u
	return nil
}

