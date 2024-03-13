package main

import (
	"fmt"

	"company.com/chainOfResponsibility/authentication"
	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")
	var authenticator authentication.Authenticator = (authentication.NewCustomerAuthenticator())
	authenticator.
		SetNext(authentication.NewEmployeeAuthenticator()).
		SetNext(authentication.NewNoUserAuthenticator())
	var loginService = &authentication.LoginService{
		Authenticator: authenticator,
	}

	log.Info().Msg(fmt.Sprintf("Company login: %v", loginService.Login("customer@gmail.com", "password")))
	log.Info().Msg(fmt.Sprintf("Employee login:%v ", loginService.Login("employee@company.com", "password")))
	log.Info().Msg(fmt.Sprintf("Invalid login: %v", loginService.Login("abcd@efgh.com", "password")))
}
