package examples

import (
	"fmt"
	"time"
)

type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Username  string    `json:"username"`
	KnownIps  []string  `json:"knownIps"`
	Profile   Profile   `json:"profile"`
	Id        string    `json:"id"`
}
type Profile struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	StaticData []int  `json:"staticData"`
}
type Users []User

func GetUsers() (Users, error) {
	res, err := client.Get(baseURL, nil)
	if err != nil {
		return nil, err
	}
	var users Users
	err = res.Unmarshal(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func GetUserByID(id int) (*User, error) {
	url := fmt.Sprintf("%s%d", baseURL, id)
	res, err := client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var user User
	err = res.Unmarshal(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
