package main

import (
	"errors"
	"sync"
)

type Employee struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Sex string `json:"sex"`
	Age int `json:"age"`
	Salary int `json:"salary"`
	Office int `json:"office"`
}

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int)
	GetOffice(id int) (Employee, error)
	Check(office int, e Employee) (choose bool)
}

type MemoryStorage struct {
	counter int 
	data map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()

	e.ID = s.counter
	s.data[e.ID] = *e

	s.counter++

	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	employee, ok := s.data[id]
	if !ok {
		return employee, errors.New("employee not found")
	}

	return employee, nil
}

func (s *MemoryStorage) Update(id int, e Employee) {
	s.Lock()
	s.data[id] = e
	s.Unlock()
}

func (s *MemoryStorage) Delete(id int) {
	s.Lock()
	delete(s.data, id)
	s.Unlock()
}

func (s *MemoryStorage) GetOffice(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	employee, ok := s.data[id]
	if !ok {
		return employee, errors.New("employee not found")
	}
	
	return employee, nil
}



func (s *MemoryStorage) Check(office int, e Employee) (choose bool) {
	s.Lock()
	defer s.Unlock()
	
	switch e.Office {
	case office:
		choose := true
		return choose
	default:
		choose := false
		return choose
	}
}
