package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/krlosaren/hostal-app/internal/users/models"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&models.User{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, stmts)
}
