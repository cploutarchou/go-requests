package examples

import (
	"reflect"
	"testing"
)

func TestGetContasts(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want []Contact
	}{
		{
			name: "Test Get Endpoint",
			args: args{
				endpoint: baseURL + "v1/contacts",
			},
			want: []Contact{
				{
					ID:        "11111",
					FirstName: "Tom",
					LastName:  "Smith",
					Email:     "tom.smith@example.com",
					DateAdded: "2021-01-03",
					CompanyID: "123",
				},
				{
					ID:        "22222",
					FirstName: "Suki",
					LastName:  "Patel",
					Email:     "spatel@example.com",
					DateAdded: "2020-11-12",
					CompanyID: "123",
				},
				{
					ID:        "33333",
					FirstName: "Lexine",
					LastName:  "Barnfield",
					Email:     "barnfield8@example.com",
					DateAdded: "2021-01-03",
					CompanyID: "234",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetContacts(tt.args.endpoint)
			if err != nil {
				t.Errorf("GetContacts() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContacts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContactByID(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want Contact
	}{
		{
			name: "Test Get Contact By ID",
			args: args{
				url: baseURL + "v1/contacts",
			},
			want: Contact{
				ID:        "433",
				FirstName: "lq8g98",
				LastName:  "ws4uj23qtc",
				Email:     "0t0jweqtyxdz@example.com",
				DateAdded: "2022-09-24",
				CompanyID: "123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetContactByID(tt.args.url, 433)
			if err != nil {
				t.Errorf("GetContactByID() error = %v", err)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("GetContactByID() got = %v, \n want %v", *got, tt.want)
			}
		})
	}
}
