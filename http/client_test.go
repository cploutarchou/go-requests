package http

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var baseURL = "https://go-requests.wiremockapi.cloud"

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
			name: "Test Get",
			args: args{
				tag: "cm8rvd96sgb7ev7dmli6pqz8vlpfx86egsiw6cejq1q1npe9yu45q27260b5td9ee90eiie7" +
					"q49rb2xtmo26qq4shqfh6farkm8fz5ddpn7jq64dtdd16e1j8z99cesaxz65bj252y930hbsbfchir4l030z2rhuaf",
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
			resp, err := client.Get(baseURL+"/pet/findByTags", nil)
			if err != nil {
				t.Errorf("findByTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var got PetsTags
			err = resp.Unmarshal(&got)
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
	type Order struct {
		Id          int64     `json:"id"`
		PetId       int64     `json:"petId"`
		Quantity    int64     `json:"quantity"`
		ShipDate    time.Time `json:"shipDate"`
		Status      string    `json:"status"`
		Complete    bool      `json:"complete"`
		OrderStatus string    `json:"orderStatus"`
	}
	type args struct {
		item Order
	}
	t1, _ := time.Parse("2006-01-02 15:04:05", "2022-12-12 00:00:00")
	tests := []struct {
		name    string
		args    args
		want    Order
		wantErr bool
	}{
		{
			name: "Test Post",
			args: args{
				item: Order{
					PetId:    670792158758028421,
					Quantity: 2,
					Id:       6075746898333402660,
					ShipDate: t1,
					Complete: true,
					Status:   "approved",
				},
			},
			want: Order{
				PetId:    670792158758028421,
				Quantity: 2,
				Id:       6075746898333402660,
				ShipDate: t1,
				Complete: true,
				Status:   "approved",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.args.item)
			if err != nil {
				t.Errorf("placeOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			builder := NewBuilder()
			builder.Headers().
				SetContentType("application/json").
				SetAccept("application/json")
			
			builder.SetRequestTimeout(10 * time.Second).
				SetResponseTimeout(10 * time.Second).
				SetMaxIdleConnections(10)
			client := builder.Build()
			res, err := client.Post(baseURL+"/store/order", nil, data)
			if err != nil {
				t.Errorf("placeOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var got Order
			err = res.Unmarshal(&got)
			if (err != nil) != tt.wantErr {
				t.Errorf("placePetOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Id != tt.want.Id {
				t.Errorf("placePetOrder() = %v, want %v", got.Id, tt.want.Id)
				return
			}
			
			// check if got is deep equal to want
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("placePetOrder() = \n%v \n, want \n%v \n", got, tt.want)
				return
			}
		})
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
