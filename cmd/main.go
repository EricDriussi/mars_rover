package cmd

import (
	"database/sql"
	"fmt"
	"mars_rover/internal/infra"
	"mars_rover/internal/use_case/create"
	"mars_rover/internal/use_case/move"
)

// TODO: LIST OF THINGS!
// GUI
// API
func Sample1() {
	db, repository := infra.InitFS()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("err closing db connection")
		}
	}(db)
	curiosity, err := create.Random(repository)
	if err != nil {
		panic("Could not initialize game :c")
	}
	moveUseCase := move.For(curiosity, repository)
	movementErrors := moveUseCase.MoveSequence("ffrfblf")

	for _, err := range movementErrors {
		fmt.Println(err)
	}

	fmt.Println("Rover finished moving")
}

func Sample2() {
	db, repository := infra.InitFS()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("err closing db connection")
		}
	}(db)
	curiosity, err := create.Random(repository)
	if err != nil {
		panic("Could not initialize game :c")
	}
	moveUseCase := move.For(curiosity, repository)
	err = moveUseCase.MoveSequenceAborting("ffrfblf")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rover finished moving")
	}
}
