package examples

import (
	"reflect"
	"testing"
	"time"
)

func TestGetUsers(t *testing.T) {

	tests := []struct {
		name string
		want int
	}{
		{
			name: "Test get all users",
			want: 39,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUsers()
			if err != nil {
				t.Errorf("GetUsers() error = %v", err)
				return
			}
			if len(got) != tt.want {
				t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
			for _, user := range got {
				if user.Id == "" {
					t.Errorf("GetUsers() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
func TestGetUserByID(t *testing.T) {
	createdAt, err := time.Parse(time.RFC3339, "2022-12-25T08:54:29.695Z")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name string
		want User
	}{
		{
			name: "Test get user by id",
			want: User{
				CreatedAt: createdAt,
				Name:      "Mario Little",
				Avatar:    "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1009.jpg",
				Username:  "Everett_Paucek",
				KnownIps: []string{"201.175.188.189",
					"43de:bbde:97fa:a4c2:fa61:7f56:d7e7:548c"},
				Profile: Profile{
					FirstName:  "Pearlie",
					LastName:   "Bartoletti",
					StaticData: []int{100, 200, 300},
				},
				Id: "1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserByID(1)
			if err != nil {
				t.Errorf("GetContactByID() error = %v", err)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("GetContactByID() \ngot = %v, \n want %v", *got, tt.want)
			}
		})
	}
}
