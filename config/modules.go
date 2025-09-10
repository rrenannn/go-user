package config

import (
	"github.com/rrenannn/go-user/infra/crypt"
	"github.com/rrenannn/go-user/internal/user"
	"go.uber.org/fx"
)

var AllModules = fx.Options(
	user.Module,
	crypt.Module,
)
