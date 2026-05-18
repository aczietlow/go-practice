package main

import (
	"fmt"
	"math/rand"
)

type User struct {
	id   int
	name string
}

func getUser(id int) User {
	u := User{
		id: id,
	}

	userNames := []string{
		"Blake",
		"Ricky",
		"Shelley",
		"Dave",
		"George",
		"John",
		"James",
		"Mitch",
		"Williamson",
		"Burry",
		"Vennett",
		"Shipley",
		"Geller",
		"Rickert",
		"Carrell",
		"Baum",
		"Brownfield",
		"Lippmann",
		"Moses",
	}
	u.name = fmt.Sprintf("%s#%d", userNames[id%len(userNames)], id)

	return u
}

func getUsers(num int) []User {
	// Deterministic random numbers.
	r := rand.New(rand.NewSource(1))
	users := []User{}
	ids := []int{}
	// Build a pool of ids three times larger than needed, so we can randomly select ids.
	for i := range num * 3 {
		ids = append(ids, i)
	}
	r.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})
	ids = ids[:num]
	for _, id := range ids {
		users = append(users, getUser(id))
	}

	return users
}

func (u *User) lt(otherUser User) bool {
	return u.id < otherUser.id
}

func (u *User) gt(otherUser User) bool {
	return u.id > otherUser.id
}

func (u *User) equal(otherUser User) bool {
	return u.id == otherUser.id
}
