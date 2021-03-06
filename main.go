package main

import (
	"fmt"

	"github.com/unknowntpo/task/cmd"
	"github.com/unknowntpo/task/db"
)

func main() {
	err := db.Init("./my.db")
	if err != nil {
		fmt.Errorf("Failed to init database: %v", err)
	}

	err = cmd.Init()
	if err != nil {
		fmt.Errorf("Failed to init cmd package: %v", err)
	}
}
