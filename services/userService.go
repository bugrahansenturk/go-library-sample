package services

import (
	"errors"
	"sync"

	domain "library-sample/domains"
)

type UserService struct {
	users   map[int]domain.User
	idMutex sync.Mutex
	nextID  int
}

func NewUserService() *UserService {
	service := &UserService{
		users:  make(map[int]domain.User),
		nextID: 1,
	}
	// Add initial users
	service.AddUser(domain.User{
		FirstName: "Alice",
		LastName:  "Smith",
		Email:     "alice@example.com",
		Role:      domain.Member,
	})
	service.AddUser(domain.User{
		FirstName: "Bob",
		LastName:  "Johnson",
		Email:     "bob@example.com",
		Role:      domain.Member,
	})
	service.AddUser(domain.User{
		FirstName: "Charlie",
		LastName:  "Brown",
		Email:     "charlie@example.com",
		Role:      domain.Expired,
	})
	return service
}

func (s *UserService) ListUsers() []domain.User {
	users := make([]domain.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

func (s *UserService) GetUserByID(id int) (domain.User, error) {
	user, ok := s.users[id]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) AddUser(user domain.User) (domain.User, error) {
	s.idMutex.Lock()
	user.ID = s.nextID
	s.nextID++
	s.idMutex.Unlock()

	s.users[user.ID] = user
	return user, nil
}

func (s *UserService) UpdateUser(updatedUser domain.User) error {
	_, ok := s.users[updatedUser.ID]
	if !ok {
		return errors.New("user not found")
	}
	s.users[updatedUser.ID] = updatedUser
	return nil
}

func (s *UserService) DeleteUser(id int) error {
	_, ok := s.users[id]
	if !ok {
		return errors.New("user not found")
	}
	delete(s.users, id)
	return nil
}
