package main

import "github.com/SGDIEGO/CleanCode/internal/core/app"

type user struct {
	name string
}

var USERS = []user{
	{"diego"},
	{"alonso"},
	{"manolo"},
}

func getAll(id int, c chan user) {
	c <- USERS[id]
}

func main() {
	{
		// var ids = []int{0, 1, 2}
		// idChan := make(chan user)

		// var USERSGet = []user{}

		// for _, id := range ids {
		// 	go getAll(id, idChan)
		// 	user := <-idChan
		// 	USERSGet = append(USERSGet, user)
		// }

		// fmt.Println(USERSGet)
	}
	// New App
	App := app.NewApp()
	// Run App
	App.RunApp()
}
