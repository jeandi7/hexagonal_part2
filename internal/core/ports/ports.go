package ports

import "hexagonal_part2/internal/core/domain"

type UserPort interface {

	// return all the users in the repository
	ListUsers() ([]domain.User, error)

	// UserPort logger
	Info(s string)
	Warn(s string)
	Error(s string)
}

type ActionPort interface {

	// simulates the return of  all the users in a technical way  (http, grpc, ...)
	ListUsers() [][]string
}
