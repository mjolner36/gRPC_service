package server

import (
	"context"
	"log"
	"mjolner36/gRPC_service/pb"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	users  map[string]string
	tokens map[string]string
	mu     sync.Mutex
}

func NewAuthService() *AuthService {
	return &AuthService{
		users:  make(map[string]string),
		tokens: make(map[string]string),
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[req.Username]; exists {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	s.users[req.Username] = req.Password
	log.Printf("User registered: %s\n", req.Username)
	return &pb.RegisterResponse{Message: "User registered successfully"}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pass, exists := s.users[req.Username]
	if !exists || pass != req.Password {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	token := "token123" // Можно заменить на JWT или UUID
	s.tokens[token] = req.Username
	log.Printf("User logged in: %s\n", req.Username)
	return &pb.LoginResponse{Token: token}, nil
}

func (s *AuthService) CheckToken(ctx context.Context, req *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, valid := s.tokens[req.Token]
	return &pb.CheckTokenResponse{Valid: valid}, nil
}
