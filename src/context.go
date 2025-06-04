package main

import (
	"fmt"
	"math/rand"
)

func generateSessionID() string {
	// Generate a unique session ID, in hex
	return fmt.Sprintf("%016x", rand.Int63())
}

type Context struct {
	// Mapping cookies to sessions
	Sessions map[string]*Session
}

func NewContext() *Context {
	return &Context{
		Sessions: make(map[string]*Session),
	}
}

func (c *Context) GetSession(cookie string) *Session {
	if session, exists := c.Sessions[cookie]; exists {
		return session
	}

	return nil
}

func (c *Context) CreateSession() (string, *Session) {
	sessionID := generateSessionID()
	session := NewSession()
	c.Sessions[sessionID] = session
	return sessionID, session
}
