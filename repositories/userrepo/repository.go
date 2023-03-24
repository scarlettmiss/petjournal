package userrepo

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	"sync"
	"time"
)

type Repository struct {
	mux   sync.Mutex
	users map[uuid.UUID]user.User
}

func New() *Repository {
	return &Repository{
		users: map[uuid.UUID]user.User{},
	}
}

func (r *Repository) CreateUser(u user.User) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	id, err := uuid.NewUUID()
	if err != nil {
		return user.Nil, err
	}
	now := time.Now()
	u.Id = id
	u.CreatedAt = now
	u.UpdatedAt = now
	u.Deleted = false
	r.users[u.Id] = u

	return u, nil
}

func (r *Repository) User(id uuid.UUID) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	u, ok := r.users[id]
	if !ok {
		return user.Nil, user.ErrNotFound
	}

	return u, nil
}

func (r *Repository) Users() map[uuid.UUID]user.User {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.users
}

func (r *Repository) UpdateUser(u user.User) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.users[u.Id]
	if !ok {
		return user.Nil, user.ErrNotFound
	}
	now := time.Now()
	u.UpdatedAt = now

	r.users[u.Id] = u

	return u, nil
}

func (r *Repository) DeleteUser(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.users[id]
	if !ok {
		return user.ErrNotFound
	}

	delete(r.users, id)

	return nil
}
