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

func (userService *UserService) ListUsers() []domain.User {
	users := make([]domain.User, 0, len(userService.users))
	for _, user := range userService.users {
		users = append(users, user)
	}
	return users
}

func (userService *UserService) GetUserByID(id int) (domain.User, error) {
	user, ok := userService.users[id]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (userService *UserService) AddUser(user domain.User) (domain.User, error) {
	userService.idMutex.Lock()
	user.ID = userService.nextID
	userService.nextID++
	userService.idMutex.Unlock()

	userService.users[user.ID] = user
	return user, nil
}

func (userService *UserService) UpdateUser(updatedUser domain.User) error {
	_, ok := userService.users[updatedUser.ID]
	if !ok {
		return errors.New("user not found")
	}
	userService.users[updatedUser.ID] = updatedUser
	return nil
}

func (userService *UserService) DeleteUser(id int) error {
	_, ok := userService.users[id]
	if !ok {
		return errors.New("user not found")
	}
	delete(userService.users, id)
	return nil
}
