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
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int)
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}



func NewMemoryStorage() *MemoryStorage {
//Adding 2 people initially
	m1 := make(map[int]Employee)
	var e1= Employee{ID:1, Name:"yura", Sex:"M", Age:25, Salary:1500  }
	var e2= Employee{ID:2, Name:"alex", Sex:"M", Age:25, Salary:1500  }
	m1[1]=e1
	m1[2]=e2
	//var mem=MemoryStorage{counter:3, data:m1}

	return &MemoryStorage{
		//data:    make(map[int]Employee),
		data: m1,
		counter: 3,
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
