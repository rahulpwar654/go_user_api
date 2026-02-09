package repository

import (
	"sort"

	"gin-user-api/internal/model"
)

type UserRepository interface {
	GetByID(id int64) (*model.User, error)
	ListPaged(limit, offset int) ([]model.User, int64, error)
	Create(u *model.User) error
	Update(u *model.User) error
	Delete(id int64) error
}

type InMemoryUserRepository struct {
	data map[int64]*model.User
	next int64
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data: make(map[int64]*model.User),
		next: 1,
	}
}

func (r *InMemoryUserRepository) GetByID(id int64) (*model.User, error) {
	u, ok := r.data[id]
	if !ok {
		return nil, nil
	}
	return u, nil
}

func (r *InMemoryUserRepository) Create(u *model.User) error {
	u.ID = r.next
	r.next++
	r.data[u.ID] = u
	return nil
}

func (r *InMemoryUserRepository) ListPaged(limit, offset int) ([]model.User, int64, error) {
	users := make([]model.User, 0, len(r.data))
	for _, u := range r.data {
		users = append(users, *u)
	}
	sort.Slice(users, func(i, j int) bool { return users[i].ID < users[j].ID })

	total := int64(len(users))
	if limit <= 0 || offset < 0 || offset >= len(users) {
		return []model.User{}, total, nil
	}

	start := offset
	end := offset + limit
	if end > len(users) {
		end = len(users)
	}
	return users[start:end], total, nil
}

func (r *InMemoryUserRepository) Update(u *model.User) error {
	if _, ok := r.data[u.ID]; !ok {
		return nil
	}
	r.data[u.ID] = u
	return nil
}

func (r *InMemoryUserRepository) Delete(id int64) error {
	delete(r.data, id)
	return nil
}
