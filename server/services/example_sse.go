// EXAMPLE: SSE — remove with example
package services

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"sync"
	"time"
)

type ExampleSSESalesService struct {
	mu    sync.Mutex
	value int
	subs  map[chan string]struct{}
}

func NewExampleSSESalesService() *ExampleSSESalesService {
	s := &ExampleSSESalesService{
		value: 1000,
		subs:  make(map[chan string]struct{}),
	}
	go s.run()
	return s
}

func (s *ExampleSSESalesService) run() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		s.mu.Lock()
		s.value += rand.IntN(11) + 10
		formatted := formatUSD(s.value)
		subs := make([]chan string, 0, len(s.subs))
		for ch := range s.subs {
			subs = append(subs, ch)
		}
		s.mu.Unlock()
		for _, ch := range subs {
			select {
			case ch <- formatted:
			default:
			}
		}
	}
}

func (s *ExampleSSESalesService) Subscribe() (chan string, func()) {
	ch := make(chan string, 1)
	s.mu.Lock()
	s.subs[ch] = struct{}{}
	s.mu.Unlock()
	return ch, func() {
		s.mu.Lock()
		if _, ok := s.subs[ch]; ok {
			delete(s.subs, ch)
			close(ch)
		}
		s.mu.Unlock()
	}
}

func (s *ExampleSSESalesService) Current() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return formatUSD(s.value)
}

func formatUSD(n int) string {
	s := fmt.Sprintf("%d", n)
	var b strings.Builder
	b.WriteByte('$')
	pre := len(s) % 3
	if pre > 0 {
		b.WriteString(s[:pre])
		if len(s) > pre {
			b.WriteByte(',')
		}
	}
	for i := pre; i < len(s); i += 3 {
		b.WriteString(s[i : i+3])
		if i+3 < len(s) {
			b.WriteByte(',')
		}
	}
	return b.String()
}
