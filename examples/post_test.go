package examples

import (
	"fmt"
	"testing"
	"time"
)

func TestPostCreateUser(t *testing.T) {
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Test Create User",
			args: args{
				user: User{
					CreatedAt: time.Now(),
					Name:      "John Doe",
					Avatar:    "https://s3.amazonaws.com/uifaces/faces/twitter/calebogden/128.jpg",
					Username:  "johndoe",
					KnownIps: []string{
						"201.175.188.189",
						"43de:bbde:97fa:a4c2:fa61:7f56:d7e7:548c",
					},
					Profile: Profile{
						FirstName:  "John",
						LastName:   "Doe",
						StaticData: []int{100, 200, 300},
					},
				},
			},
			want: &User{
				CreatedAt: time.Now(),
				Name:      "John Doe",
				Avatar:    "https://s3.amazonaws.com/uifaces/faces/twitter/calebogden/128.jpg",
				Username:  "johndoe",
				KnownIps: []string{
					"201.175.188.189",
					"43de:bbde:97fa:a4c2:fa61:7f56:d7e7:548c",
				},
				Profile: Profile{
					FirstName:  "John",
					LastName:   "Doe",
					StaticData: []int{100, 200, 300},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostCreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostCreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("PostCreateUser() = %v, want %v", got, tt.want)
			}
			fmt.Print("got: ", got.Status())
			fmt.Print("got: ", got.String())
			fmt.Print("got: ", got)
		})
	}
}
