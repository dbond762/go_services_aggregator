package domain

import (
	"fmt"
)

type ServiceProvider func() (Service, error)

var serviceProviders = make(map[string]ServiceProvider)

func RegisterService(ident string, provider ServiceProvider) {
	if provider == nil {
		panic("services provider: register provider is nil")
	}
	if _, dup := serviceProviders[ident]; dup {
		panic(fmt.Sprintf("services provider: register called twice for provider %s", ident))
	}

	serviceProviders[ident] = provider
}

func CreateService(ident string) (Service, error) {
	provider, ok := serviceProviders[ident]
	if !ok {
		return nil, fmt.Errorf("services provider: unknown service %q", ident)
	}

	return provider()
}
