package MutexSet

import (
	"sync"
)

/*
	Протестируйте производительность операций чтения и записи
	на множестве действительных чисел, безопасность которого
	обеспечивается sync.Mutex и sync.RWMutex для разных вариантов
	использования: 10% запись, 90% чтение; 50% запись, 50% чтение;
	90% запись, 10% чтение
*/

type Set struct {
	sync.Mutex
	mm map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}

func(s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func(s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}


