package main

import (
	"fmt"
	"go_postgresql/lib"

	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("This is main")

	persons := lib.GetPersons()

	fmt.Println(*persons)
}
