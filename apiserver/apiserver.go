package apiserver

import (
	"fmt"
	"github.com/snarlysodboxer/k8s-spec/engine"
	"github.com/snarlysodboxer/k8s-spec/state"
)

type Server struct {
	store  state.Interface
	engine engine.Interface
}

func NewServer(store state.Interface, eng engine.Interface) *Server {
	return &Server{store, eng}
}

func (server *Server) ListenAndServe() error {
	fmt.Println("Listen and Serve placeholder")
	return nil
}
