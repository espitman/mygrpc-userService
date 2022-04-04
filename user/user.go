package user

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) GetUser(ctx context.Context, in *UserRequest) (*UserResponse, error) {
	log.Println("Receive message body from client:")
	return &UserResponse{Id: in.Id, Fname: "Saeed", Lname: "Heidari"}, nil
}
