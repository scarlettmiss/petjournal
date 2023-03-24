package userrepo

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	"sync"
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

func (r *Repository) CreateUser(u user.User) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.users[u.Id] = u

	return nil
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

func (r *Repository) UpdateUser(u user.User) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.users[u.Id]
	if !ok {
		return user.ErrNotFound
	}

	r.users[u.Id] = u

	return nil
}

func (r *Repository) DeleteUser(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	u, ok := r.users[id]
	if !ok {
		return user.ErrNotFound
	}

	u.Deleted = true
	r.users[u.Id] = u

	return nil
}
