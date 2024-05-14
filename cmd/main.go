package main

import (
	action "hexagonal_part2/internal/adapters/primary"

	"os"

	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Hexagonal_part2 start ...")

	UserAction := action.NewUserAction(logger)
	logger.Info("Response to the command -->", "resp", UserAction.ListUsers())

	logger.Info("Hexagonal_part2 end")

}
