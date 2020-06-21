package modeltests

import (
	"log"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllUsers(t *testing.T) {
	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}
	users, err := userInstance.FindAllUsers(s.DB)
	if err != nil {
		t.Errorf("this is the error getting th users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}
