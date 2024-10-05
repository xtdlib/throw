package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/xtdlib/throw"
)

func foo() (error) {
	return throw.Wrap1(os.Open("non-existing-file"))
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	err := foo()
	slog.Error("something is wrong", throw.SlogAttr(foo()))
	slog.Info(fmt.Sprintf("this should print true: %v", errors.Is(err, os.ErrNotExist)))
}
