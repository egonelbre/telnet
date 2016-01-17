package telnet

import (
	"log"
	"net"
)

type Handler func(*Conn)

type Server struct {
	Addr     string
	Handler  Handler
	ErrorLog *log.Logger
}

func ListenAndServe(addr string, handler Handler) error {
	server := &Server{addr, handler, nil}
	return server.ListenAndServe()
}

func (srv *Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}

func (srv *Server) logf(format string, args ...interface{}) {
	if srv.ErrorLog != nil {
		srv.ErrorLog.Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func (srv *Server) Serve(ln net.Listener) error {
	defer ln.Close()

	for {
		rw, err := ln.Accept()
		if err != nil {
			srv.logf("failed to accept connection: %v", err)
			continue
		}

		go srv.Handler(srv.newConn(rw))
	}
	return nil
}

func (srv *Server) newConn(rwc net.Conn) *Conn {
	return NewConn(rwc)
}
