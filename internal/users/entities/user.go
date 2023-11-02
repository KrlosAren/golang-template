package entities

import (
	"fmt"
	"time"

	"github.com/krlosaren/hostal-app/internal/users/dto"
)

type (
	User struct {
		UserID    int64
		FullName  string
		Email     string
		Password  string
		Age       int8
		CreatedAt time.Time
		UpadateAt time.Time
	}
)

func NewUser(data dto.CreateUserRequest) *User {

	fullname := fmt.Sprintf("%s %s", data.Name, data.LastName)

	return &User{
		FullName:  fullname,
		Email:     data.Email,
		Password:  data.Password,
		Age:       data.Age,
		CreatedAt: time.Now(),
		UpadateAt: time.Now(),
	}
}
