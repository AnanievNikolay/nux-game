package game

import (
	"math/rand"
	"time"
)

func (s *Service) getNumber() int {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	return rnd.Intn(s.maxGameNumber)
}
