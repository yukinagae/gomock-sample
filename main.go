package main

import "fmt"

func main() {
	var repo Repository
	repo = UserRepository{}
	user := repo.Get(1)
	fmt.Printf("%v\n", user)
}
