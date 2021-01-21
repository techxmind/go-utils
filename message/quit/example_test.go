package quit_test

import (
	"fmt"

	"github.com/techxmind/go-utils/message/quit"
)

type Service struct {
	quit       chan bool
	tasks      chan int
	bufferSize int
	buffer     []int
	result     int
}

func (s *Service) worker() {
	process := func(tasks []int) int {
		ret := 0
		for _, task := range tasks {
			ret += task
		}
		return ret
	}

	for {
		select {
		case task, ok := <-s.tasks:
			if ok {
				s.buffer = append(s.buffer, task)
				if len(s.buffer) >= s.bufferSize {
					s.result += process(s.buffer)
					s.buffer = make([]int, 0, s.bufferSize)
				}
			}
		case <-s.quit:
			s.result += process(s.buffer)
			s.quit <- true
			return
		}
	}
}

func (s *Service) Close() {
	s.quit <- true
	<-s.quit
}

func Example() {
	s := &Service{
		quit:       make(chan bool),
		tasks:      make(chan int),
		bufferSize: 1,
		buffer:     make([]int, 0, 1),
	}

	s2 := &Service{
		quit:       make(chan bool),
		tasks:      make(chan int),
		bufferSize: 5,
		buffer:     make([]int, 0, 5),
	}

	go s.worker()
	go s2.worker()

	quit.Register(s)
	quit.Register(s2)

	s.tasks <- 1
	s.tasks <- 1

	s2.tasks <- 1
	s2.tasks <- 1

	fmt.Println(s.result)
	fmt.Println(s2.result)

	quit.Broadcast()

	fmt.Println(s.result)
	fmt.Println(s2.result)

	//output:
	// 1
	// 0
	// 2
	// 2
}
