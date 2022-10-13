package main

import (
	"fmt"
	"github.com/SBEKzy/testTask2/config"
	"github.com/SBEKzy/testTask2/model"
	"github.com/SBEKzy/testTask2/repository"
	"github.com/SBEKzy/testTask2/service"
)

/*
1 - user
var exampleUser = model.User{
	Name:     "BBB",
	Email:    "BBB@BB.BBB",
	Password: "BbBbBbBb",
}
*/
/*
//2 - user
var exampleUser = model.User{
	Name:     "DDD",
	Email:    "DDD@DD.DDD",
	Password: "DdDdDdDd",
}
*/
//3-user
var exampleUser = model.User{
	Name:     "LLL",
	Email:    "LLL@LL.LLL",
	Password: "LdLdLdLd",
}
var replacementUser = model.User{
	Name: "DDDGGGDDD",
}

func main() {
	config.New()
	db := repository.ConnectDB()
	repo := repository.New(db)
	s := service.New(repo)

	err := s.CreateUser(&exampleUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("added", exampleUser.Email)

	user, err := s.GetUser(exampleUser.Email)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)

	exampleUser.Name = "replacedName"
	exampleUser.Password = ""
	err = s.UpdateUser(exampleUser.Email, &exampleUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("UPDATED")

	user, err = s.GetUser(exampleUser.Email)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
	/*
		err := s.DeleteUser(exampleUser.Email)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("deleted", exampleUser.Email)*/
	fmt.Println("done")
}
