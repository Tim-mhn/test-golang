package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.Id != 0 {
		return User{}, errors.New("User must not contain id or id must equal 0")
	}
	u.Id = nextId
	nextId++
	fmt.Println("next id is ", nextId)
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("could not find user with id %d", id)
}

func DeleteById(id int) error {
	for index, u := range users {
		if u.Id == id {
			updatedUsers := users[:index]
			updatedUsers = append(updatedUsers, users[index+1:]...)
			users = updatedUsers
			return nil
		}
	}

	return fmt.Errorf("could not find user with id %d", id)
}
