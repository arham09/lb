package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Server struct {
	URL          *url.URL
	alive        bool
	mux          sync.RWMutex
	reverseProxy *httputil.ReverseProxy
}

// SetAlive for this backend
func (s *Server) SetAlive(alive bool) {
	s.mux.Lock()
	s.alive = alive
	s.mux.Unlock()
}

func (s *Server) IsAlive() bool {
	s.mux.RLock()
	alive := s.alive
	s.mux.RUnlock()
	return alive
}
