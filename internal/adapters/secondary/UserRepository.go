package repository

import (
	"hexagonal_part2/internal/core/domain"

	"golang.org/x/exp/slog"
)

var (
	log *slog.Logger
)

func initLogger(logger *slog.Logger) {
	log = logger
}

type UserRepository struct{}

func NewUserRepository(logger *slog.Logger) *UserRepository {
	initLogger(logger)
	log.Info("The adapter UserRepository for the secondary (outBound) port is created")
	return &UserRepository{}
}

// the level of the  logs
func (ur *UserRepository) Info(info string) {
	log.Info(info)
}
func (ur *UserRepository) Warn(info string) {
	log.Warn(info)
}
func (ur *UserRepository) Error(info string) {
	log.Error(info)
}

// service simulates access to a database
func (ur *UserRepository) ListUsers() ([]domain.User, error) {
	ur_list := []domain.User{
		{ID: "u1", Name: "John"},
		{ID: "u2", Name: "Peter"},
		{ID: "u3", Name: "Elizabeth"},
	}
	return ur_list, nil
}
