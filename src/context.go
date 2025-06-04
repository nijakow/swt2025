package main

import (
	"fmt"
	"math/rand"
	"net/http"
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

// Create a global context variable
var GlobalContext = NewContext()

// Create a system that handles the cookie stuff based on a HTTP request
type CookieBlock struct {
	sessionid string
	Session   *Session
}

func handleRequestCookie(request *http.Request) CookieBlock {
	cookie, err := request.Cookie("session_id")

	var session *Session
	var sessionid string

	if err == nil {
		sessionid = cookie.Value
		session = GlobalContext.GetSession(sessionid)
	} else {
		sessionid, session = GlobalContext.CreateSession()
	}

	return CookieBlock{
		sessionid: sessionid,
		Session:   session,
	}
}

// Make sure the correct headers are sent
func handleResponseCookie(w http.ResponseWriter, block *CookieBlock) {
	w.Header().Set("Set-Cookie", fmt.Sprintf("session_id=%s; Path=/; HttpOnly; Secure", block.sessionid))
}
