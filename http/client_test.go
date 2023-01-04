package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func Test_goHTTPClient_Get(t *testing.T) {
	type args struct {
		tag string
	}
	type PetTag struct {
		PhotoUrls []string `json:"photoUrls"`
		Name      string   `json:"name"`
		ID        int64    `json:"id"`
		Category  struct {
			Name string `json:"name"`
			ID   int64  `json:"id"`
		} `json:"category"`
		Tags []struct {
			Name string `json:"name"`
			ID   int64  `json:"id"`
		} `json:"tags"`
		Status string `json:"status"`
	}
	type PetsTags []PetTag
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test_findPetsByTagXML",
			args: args{
				tag: "cm8rvd96sgb7ev7dmli6pqz8vlpfx86egsiw6cejq1q1npe9yu45q27260b5td9ee90eiie7q49rb2xtmo26qq4shqfh6farkm8fz5ddpn7jq64dtdd16e1j8z99cesaxz65bj252y930hbsbfchir4l030z2rhuaf",
			},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewBuilder()
			builder.Headers().
				SetContentType("application/json").
				SetAccept("application/json")
			
			builder.SetRequestTimeout(10 * time.Second).
				SetResponseTimeout(10 * time.Second).
				SetMaxIdleConnections(10)
			client := builder.Build()
			client.QueryParams().Set("tags", tt.args.tag)
			resp, err := client.Get("https://go-requests.wiremockapi.cloud/pet/findByTags", nil)
			if err != nil {
				t.Errorf("findByTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var pets PetsTags
			err = resp.Unmarshal(&pets)
			got := pets
			if (err != nil) != tt.wantErr {
				t.Errorf("findByTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("findByTag() = %v, want %v", len(got), tt.want)
				return
			}
		})
	}
}

func Test_goHTTPClient_Post(t *testing.T) {
	type User struct {
		Name     string `json:"name"`
		UserName string `json:"username"`
	}
	user := User{
		Name:     "Christos Ploutarchou",
		UserName: "christos",
	}
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Post",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url: "https://api.github.com",
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				client.DisableTimeouts()
				data, err := json.Marshal(user)
				if err != nil {
					t.Errorf("goHTTPClient.Post() Unable to marshal data.  error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				res, err := client.Post(tt.args.url, tt.fields.Headers.GetAll(), data)
				got := &Response{
					statusCode: res.StatusCode(),
					status:     res.Status(),
				}
				fmt.Println(res.String())
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Post() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Post() =  %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_goHTTPClient_Put(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
		body    []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Put",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
				body:    nil,
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				res, err := client.Put(tt.args.url, tt.fields.Headers.GetAll(), tt.args.body)
				got := &Response{
					statusCode: res.StatusCode(),
					status:     res.Status(),
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Put() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Put() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_goHTTPClient_Patch(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
		body    []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Patch",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
				body:    nil,
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				res, err := client.Patch(tt.args.url, tt.fields.Headers.GetAll(), tt.args.body)
				got := &Response{
					statusCode: res.StatusCode(),
					status:     res.Status(),
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Patch() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Patch() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_goHTTPClient_Delete(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
		body    []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Delete",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
				body:    nil,
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				res, err := client.Delete(tt.args.url, tt.fields.Headers.GetAll(), tt.args.body)
				got := &Response{
					statusCode: res.StatusCode(),
					status:     res.Status(),
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Delete() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Delete() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_goHTTPClient_Head(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
		body    []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Head",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
				body:    nil,
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				res, err := client.Head(tt.args.url, tt.fields.Headers.GetAll(), tt.args.body)
				got := &Response{
					statusCode: res.StatusCode(),
					status:     res.Status(),
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Head() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Head() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
