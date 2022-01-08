package main

type Pool struct {
	servers []*Server
	current int64
}

// AddBackend to the server pool
func (p *Pool) AddServer(server *Server) {
	p.servers = append(p.servers, server)
}
