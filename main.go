package main

import (
	"fmt"

	"github.com/yukinagae/gomock-sample/user"
)

func main() {
	var repo user.Repository
	repo = user.UserRepository{}
	user := repo.Get(1)
	fmt.Printf("%v\n", user)
}
