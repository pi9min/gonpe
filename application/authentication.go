package application

import "github.com/ponpe/server/application/repository"

type AuthenticationApp struct {
	repos *repository.AllRepository
}

func NewAuthenticationApp(repos *repository.AllRepository) *AuthenticationApp {
	return &AuthenticationApp{
		repos: repos,
	}
}
