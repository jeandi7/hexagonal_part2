package domain

import "fmt"

type User struct {
	ID   string
	Name string
}

func (u *User) String() string {
	return fmt.Sprintf("User : {%s,%s}", u.ID, u.Name)
}
