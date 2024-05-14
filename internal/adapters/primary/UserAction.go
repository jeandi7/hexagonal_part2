package action

import (
	repository "hexagonal_part2/internal/adapters/secondary"
	"hexagonal_part2/internal/core/usecases"

	"golang.org/x/exp/slog"
)

var (
	log *slog.Logger
	uc  *usecases.UserUseCase
)

func initLogger(logger *slog.Logger) {
	log = logger
}

func initUserUseCase(ucase *usecases.UserUseCase) {
	uc = ucase

}

type UserAction struct{}

func NewUserAction(logger *slog.Logger) *UserAction {
	initLogger(logger)
	log.Info("The adapter UserAction for the primary (inBound) port is created")
	UserRepository := repository.NewUserRepository(log)
	// Next step is the moment when the adapters are plugged (injected) into the ports
	// for the example, i put this code here
	log.Info("The adapters will be plugged (injected) into the ports... ")
	ua := &UserAction{}
	uc := usecases.NewUserUseCase(ua, UserRepository)
	initUserUseCase(uc)
	return ua
}

// the level of the  logs
func (ua *UserAction) Info(info string) {
	log.Info(info)
}

func (ua *UserAction) ListUsers() [][]string {
	list_users, _ := uc.ListUsers()
	return list_users
}
