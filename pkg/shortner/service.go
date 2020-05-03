package shortner

import (
	"fmt"
	"io"
)

const (
	ErrInvalidRedirect = "%s service.redirect.store"
)

type Service interface {
	NewRedirect(url string) (*Redirect, error)
	NewRedirectIoReader(data io.Reader) (*Redirect, error)
	Encode(redirect *Redirect) ([]byte, error)
	Find(code interface{}) (*Redirect, error)
	Store(redirect *Redirect) error
}

func NewService(repository Repository, usecases UseCases) Service {
	return &service{
		Repository: repository,
		UseCases:   usecases,
	}
}

type service struct {
	Repository Repository
	UseCases   UseCases
}

func (s *service) NewRedirect(url string) (*Redirect, error) {
	return s.UseCases.NewRedirect(url)
}

func (s *service) NewRedirectIoReader(data io.Reader) (*Redirect, error) {
	return s.UseCases.NewRedirectIoReader(data)
}

func (s *service) Encode(redirect *Redirect) ([]byte, error) {
	return s.UseCases.Encode(redirect)
}

func (s *service) Find(code interface{}) (*Redirect, error) {
	return s.Repository.Find(code)
}

func (s *service) Store(redirect *Redirect) error {
	err := s.UseCases.Validate(redirect)
	if err != nil {
		return fmt.Errorf(ErrInvalidRedirect, err.Error())
	}

	return s.Repository.Store(redirect)
}
