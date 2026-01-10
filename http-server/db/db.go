package db

import (
	"fmt"
	"sync"
	"time"
)

// defining the schema of our User Model
/*
	{
		id: "1",
		username: "test",
		email: "test@gmail.com",
		updated_at: Datetime,
		created_at: Datetime
	}
*/
type User struct {
	Id         string    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Updated_at time.Time `json:"updated_at"`
	Created_at time.Time `json:"created_at"`
}

type Database struct {
	mu    sync.Mutex // for sync (only one goroutine reads/writes on the database at a time)
	users []User
}

// create a new instance of database and returns the pointer to it
func NewDatabase() *Database {
	return &Database{
		users: make([]User, 0),
	}
}

// gets all the users from the database
func (db *Database) GetAllUsers() []User {
	db.mu.Lock()
	defer db.mu.Unlock()

	users := make([]User, len(db.users))
	copy(users, db.users)
	return users
}

// create a new user in the database
func (db *Database) CreateNewUser(email, username string) User {
	db.mu.Lock()
	defer db.mu.Unlock()

	now := time.Now()
	user := User{
		Id:         fmt.Sprintf("%d", len(db.users)+1),
		Username:   username,
		Email:      email,
		Created_at: now,
		Updated_at: now,
	}

	db.users = append(db.users, user)
	return user
}
