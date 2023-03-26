package services

import (
	"errors"
	"sync"

	"library-sample/domain"
)

type BorrowService struct {
	borrows map[int]domain.Borrow
	idMutex sync.Mutex
	nextID  int
}

func NewBorrowService() *BorrowService {
	return &BorrowService{
		borrows: make(map[int]domain.Borrow),
		nextID:  1,
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
	s.idMutex.Lock()
	borrow.ID = s.nextID
	s.nextID++
	s.idMutex.Unlock()

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
