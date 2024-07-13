package handler

import (
	"awesomeProject/accounts/grpc/db/models"
	"context"
	"errors"
	"sync"

	pb "awesomeProject/proto"
)

type AccountsHandler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
	pb.AccountsServer
}

func New() *AccountsHandler {
	return &AccountsHandler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

func (s *AccountsHandler) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	if len(in.Name) == 0 {
		return nil, errors.New("name is empty")
	}

	if _, ok := s.accounts[in.Name]; ok {
		return nil, errors.New("account already exists")
	}

	s.accounts[in.Name] = &models.Account{
		Name:   in.Name,
		Amount: int(in.Amount),
	}

	return &pb.Empty{}, nil
}

func (s *AccountsHandler) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	s.guard.RLock()
	defer s.guard.RUnlock()

	if len(in.Name) == 0 {
		return nil, errors.New("name is empty")
	}

	account, ok := s.accounts[in.Name]

	if !ok {
		return nil, errors.New("account does not exist")
	}

	return &pb.GetAccountResponse{
		Name:   account.Name,
		Amount: int32(account.Amount),
	}, nil
}

func (s *AccountsHandler) PatchAccount(ctx context.Context, in *pb.PatchAccountRequest) (*pb.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	if len(in.Name) == 0 {
		return nil, errors.New("name is empty")
	}

	account, ok := s.accounts[in.Name]

	if !ok {
		return nil, errors.New("account does not exist")
	}

	account.Amount = int(in.Amount)

	return &pb.Empty{}, nil
}

func (s *AccountsHandler) ChangeAccount(ctx context.Context, in *pb.ChangeAccountRequest) (*pb.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	if len(in.Name) == 0 {
		return nil, errors.New("name is empty")
	}

	if len(in.Newname) == 0 {
		return nil, errors.New("new name is empty")
	}

	account, ok := s.accounts[in.Name]

	if !ok {
		return nil, errors.New("account does not exist")
	}

	if _, ok := s.accounts[in.Newname]; ok {
		return nil, errors.New("new account already exists")
	}

	account.Name = in.Newname
	s.accounts[in.Newname] = account
	delete(s.accounts, in.Name)

	return &pb.Empty{}, nil
}

func (s *AccountsHandler) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*pb.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	if len(in.Name) == 0 {
		return nil, errors.New("name is empty")
	}

	if _, ok := s.accounts[in.Name]; !ok {
		return nil, errors.New("account does not exist")
	}

	delete(s.accounts, in.Name)

	return &pb.Empty{}, nil
}
