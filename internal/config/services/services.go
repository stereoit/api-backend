package services

import (
	"github.com/sarulabs/di"
	user "github.com/stereoit/eventival/pkg/user/di"
)

// Services represents all general services like logger...
var services = []di.Def{
	{
		Name: "logger",
		Build: func(ctn di.Container) (interface{}, error) {
			return nil, nil
		},
	},
}

// New returns complete Services
// additional services are configured here
func New() []di.Def {
	// add user services
	svc := append(services, user.Services...)
	return svc
}
