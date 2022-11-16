package http

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_goHTTPClient_Get(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{

		{
			name: "Test goHTTPClient Get",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
			},
			want: &http.Response{
				StatusCode: 200,
				Status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewBuilder()
			client := builder.Build()
			res, err := client.Get(tt.args.url, tt.fields.Headers.GetAll())
			got := &http.Response{
				StatusCode: res.StatusCode,
				Status:     res.Status,
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("goHTTPClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goHTTPClient.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_goHTTPClient_Post(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
		body    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Post",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
				body:    nil,
			},
			want: &http.Response{
				StatusCode: 200,
				Status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				res, err := client.Post(tt.args.url, tt.fields.Headers.GetAll(), tt.args.body)
				got := &http.Response{
					StatusCode: res.StatusCode,
					Status:     res.Status,
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Post() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Post() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
