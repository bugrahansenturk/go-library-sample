package services

import (
	"errors"
	domain "library-sample/domains"
	"sync"
	"time"
)

type BorrowService struct {
	borrows     map[int]domain.Borrow
	idMutex     sync.Mutex
	nextID      int
	userService *UserService
}

func NewBorrowService(userService *UserService) *BorrowService {
	if userService == nil {
		panic("userService is nil")
	}

	return &BorrowService{
		borrows:     make(map[int]domain.Borrow),
		nextID:      1,
		userService: userService,
	}
}

func (s *BorrowService) ListBorrows() []domain.Borrow {
	borrows := make([]domain.Borrow, 0, len(s.borrows))
	for _, borrow := range s.borrows {
		borrows = append(borrows, borrow)
	}
	return borrows
}

func (s *BorrowService) GetBorrowByID(id int) (domain.Borrow, error) {
	borrow, ok := s.borrows[id]
	if !ok {
		return domain.Borrow{}, errors.New("borrow not found")
	}
	return borrow, nil
}

func (s *BorrowService) AddBorrow(borrow domain.Borrow) (domain.Borrow, error) {
	user, err := s.userService.GetUserByID(borrow.UserID)
	if err != nil {
		return domain.Borrow{}, errors.New("user not found")
	}

	if user.Role == domain.Expired {
		return domain.Borrow{}, errors.New("user membership is expired")
	}

	s.idMutex.Lock()
	borrow.ID = s.nextID
	s.nextID++
	s.idMutex.Unlock()

	borrow.DueDate = time.Now().Add(14 * 24 * time.Hour)
	s.borrows[borrow.ID] = borrow
	return borrow, nil
}

func (s *BorrowService) UpdateBorrow(updatedBorrow domain.Borrow) error {
	_, ok := s.borrows[updatedBorrow.ID]
	if !ok {
		return errors.New("borrow not found")
	}
	s.borrows[updatedBorrow.ID] = updatedBorrow
	return nil
}

func (s *BorrowService) DeleteBorrow(id int) error {
	_, ok := s.borrows[id]
	if !ok {
		return errors.New("borrow not found")
	}
	delete(s.borrows, id)
	return nil
}
