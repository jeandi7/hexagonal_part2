package usecases

import (
	"hexagonal_part2/internal/core/ports"
)

// the core domain is driven by the primary port
// the core domain is using the secondary ports to perform operations

type UserUseCase struct {
	actionPort ports.ActionPort // primary port
	userPort   ports.UserPort   // secondary Port
}

// dependancy injection here
func NewUserUseCase(actionPort ports.ActionPort, userPort ports.UserPort) *UserUseCase {
	userPort.Info("Dependancy injection of adapters into the ports is done : UserUseCase is created")
	return &UserUseCase{actionPort: actionPort, userPort: userPort}
}

func (t *UserUseCase) ListUsers() ([][]string, error) {

	// Need to Request the UserRepository port
	t.userPort.Info("UserUseCase : Command ListUsers is received by the Core Domain through the primary(inBound) port ...")

	list_users, err := t.userPort.ListUsers()
	if err != nil {
		t.userPort.Warn("User Storie :ListUsers : error")
		return nil, err
	}

	t.userPort.Info("UserUseCase : Request to the secondary(outBound) port is done")

	// converts the response into a technical format
	// domain.User --> string

	var userList [][]string
	for _, user := range list_users {
		t.userPort.Info("Convert Object " + user.String() + " from Core Domain to Technical Representation (string, http,grpc,others)")
		userSlice := []string{user.ID, user.Name}
		userList = append(userList, userSlice)
	}

	return userList, nil
}
