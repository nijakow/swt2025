package main

import (
	"fmt"
	"time"
)

type Session struct {
	SelectedZettels []string
	Name            string
}

func genSessionName() string {
	// This function should generate a random name for the session.
	// For simplicity, we return a static name here.
	return fmt.Sprintf("session-%d", time.Now().UnixNano())
}

func NewSession() *Session {
	return &Session{
		SelectedZettels: make([]string, 0),
		Name:            genSessionName(),
	}
}

func (s *Session) AddZettel(zettel string) {
	s.SelectedZettels = append(s.SelectedZettels, zettel)
}

func (s *Session) RemoveZettel(zettel string) {
	for i, z := range s.SelectedZettels {
		if z == zettel {
			s.SelectedZettels = append(s.SelectedZettels[:i], s.SelectedZettels[i+1:]...)
			return
		}
	}
}

func (s *Session) ContainsZettel(zettel string) bool {
	for _, i := range s.SelectedZettels {
		if i == zettel {
			return true
		}
	}
	return false
}

func (s *Session) ClearZettels() {
	s.SelectedZettels = make([]string, 0)
}

func (s *Session) GetZettels() []string {
	return s.SelectedZettels
}
