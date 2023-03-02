package main

import (
	"errors"
	"sync"
)

type Department struct {
	Office int `json:"office"`
	Title string `json:"title"`
}

type Managment interface {
	InsertM(d *Department)
	GetM(office int) (Department, error)
	UpdateM(office int, d Department)
	DeleteM(office int)
}

type MemoryManagment struct {
	counter int 
	data map[int]Department
	sync.Mutex
}

func NewMemoryManagment() *MemoryManagment {
	return &MemoryManagment{
		data: make(map[int]Department),
		counter: 1,
	}
}

func (m *MemoryManagment) InsertM(d *Department) {
	m.Lock()

	d.Office = m.counter
	m.data[d.Office] = *d
	

	m.counter++

	m.Unlock()
}

func (m *MemoryManagment) GetM(office int) (Department, error) {
	m.Lock()
	defer m.Unlock()

	department, ok := m.data[office]
	if !ok {
		return department, errors.New("department not found")
	}

	return department, nil
}


func (m *MemoryManagment) UpdateM(office int, d Department) {
	m.Lock()
	m.data[office] = d
	m.Unlock()
}

func (m *MemoryManagment) DeleteM(office int) {
	m.Lock()
	delete(m.data, office)
	m.Unlock()
}