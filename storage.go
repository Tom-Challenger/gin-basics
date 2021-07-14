package main

import (
	"errors"
	"sync"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(*Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int)
}

type MemoryStorage struct {
	data   map[int]Employee
	conter int
	sync.Mutex
}

// Constructor method
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:   make(map[int]Employee),
		conter: 1,
	}
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()

	e.ID = s.conter
	s.data[e.ID] = *e

	s.conter++

	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	s.Lock()

	// Отмечаем, что комада будет вызвана не посредственно перед return
	defer s.Unlock()

	employee, isExist := s.data[id]

	if !isExist {
		return employee, errors.New("employee not fount")
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