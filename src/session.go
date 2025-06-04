package main

type Session struct {
	SelectedZettels []string
}

func NewSession() *Session {
	return &Session{
		SelectedZettels: make([]string, 0),
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

func (s *Session) ClearZettels() {
	s.SelectedZettels = make([]string, 0)
}

func (s *Session) GetZettels() []string {
	return s.SelectedZettels
}
