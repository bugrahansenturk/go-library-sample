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

func (borrowService *BorrowService) ListBorrows() []domain.Borrow {
	borrows := make([]domain.Borrow, 0, len(borrowService.borrows))
	for _, borrow := range borrowService.borrows {
		borrows = append(borrows, borrow)
	}
	return borrows
}

func (borrowService *BorrowService) GetBorrowByID(id int) (domain.Borrow, error) {
	borrow, ok := borrowService.borrows[id]
	if !ok {
		return domain.Borrow{}, errors.New("borrow not found")
	}
	return borrow, nil
}

func (borrowService *BorrowService) AddBorrow(borrow domain.Borrow) (domain.Borrow, error) {
	user, err := borrowService.userService.GetUserByID(borrow.UserID)
	if err != nil {
		return domain.Borrow{}, errors.New("user not found")
	}

	if user.Role == domain.Expired {
		return domain.Borrow{}, errors.New("user membership is expired")
	}

	borrowService.idMutex.Lock()
	borrow.ID = borrowService.nextID
	borrowService.nextID++
	borrowService.idMutex.Unlock()

	borrow.DueDate = time.Now().Add(14 * 24 * time.Hour)
	borrowService.borrows[borrow.ID] = borrow
	return borrow, nil
}

func (borrowService *BorrowService) UpdateBorrow(updatedBorrow domain.Borrow) error {
	_, ok := borrowService.borrows[updatedBorrow.ID]
	if !ok {
		return errors.New("borrow not found")
	}
	borrowService.borrows[updatedBorrow.ID] = updatedBorrow
	return nil
}

func (borrowService *BorrowService) DeleteBorrow(id int) error {
	_, ok := borrowService.borrows[id]
	if !ok {
		return errors.New("borrow not found")
	}
	delete(borrowService.borrows, id)
	return nil
}
