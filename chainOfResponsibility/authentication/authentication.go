package authentication

import "github.com/rs/zerolog/log"

type Authenticator interface {
	SetNext(authenticator Authenticator) Authenticator
	Authenticate(username string, password string) bool
}

type EmployeeAuthenticator struct {
	next         Authenticator
	employeeList []string
}

func (a *EmployeeAuthenticator) SetNext(next Authenticator) Authenticator {
	a.next = next
	return next
}

func (a *EmployeeAuthenticator) Authenticate(username string, password string) bool {
	log.Info().Msg("EmployeeAuthenticator called")
	for _, emp := range a.employeeList {
		if emp == username {
			return true
		}
	}

	if a.next != nil {
		return a.next.Authenticate(username, password)
	}

	return false
}

func NewEmployeeAuthenticator() Authenticator {
	var list []string = make([]string, 2)
	list = append(list, "employee@company.com")
	list = append(list, "eployee2@company.com")
	return &EmployeeAuthenticator{
		employeeList: list,
	}
}

type CustomerAuthenticator struct {
	next         Authenticator
	customerList []string
}

func (a *CustomerAuthenticator) SetNext(next Authenticator) Authenticator {
	a.next = next
	return next
}

func (a *CustomerAuthenticator) Authenticate(username string, password string) bool {
	log.Info().Msg("CustomerAuthenticator called")
	for _, emp := range a.customerList {
		if emp == username {
			return true
		}
	}

	if a.next != nil {
		return a.next.Authenticate(username, password)
	}

	return false
}

func NewCustomerAuthenticator() Authenticator {
	var list []string = make([]string, 2)
	list = append(list, "customer@gmail.com")
	list = append(list, "customer2@gmail.com")
	return &CustomerAuthenticator{
		customerList: list,
	}
}

type NoUserAuthenticator struct {
	next Authenticator
}

func (a *NoUserAuthenticator) SetNext(next Authenticator) Authenticator {
	a.next = next
	return next
}

func (a *NoUserAuthenticator) Authenticate(username string, password string) bool {
	log.Info().Msg("NoUserAuthenticator called")
	return false
}

func NewNoUserAuthenticator() Authenticator {
	return &NoUserAuthenticator{}
}

type LoginService struct {
	Authenticator Authenticator
}

func (s *LoginService) Login(username string, password string) bool {
	return s.Authenticator.Authenticate(username, password)
}
