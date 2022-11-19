package http

import (
	"reflect"
	"testing"
	"time"
)

func Test_builderImpl_SetMaxIdleConnections(t *testing.T) {
	type fields struct {
		header  Headers
		Timeout Timeout
	}
	type args struct {
		maxConnections int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Timeout
	}{
		{
			name: "Test SetMaxIdleConnections",
			fields: fields{
				header:  NewHeaders(),
				Timeout: newTimeouts().SetResponseTimeout(0 * time.Second).SetRequestTimeout(0 * time.Second).SetMaxIdleConnections(10),
			},
			args: args{
				maxConnections: 10,
			},
			want: timeoutImpl{
				ResponseTimeout:    0,
				RequestTimeout:     0,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := builderImpl{
				header:  tt.fields.header,
				Timeout: tt.fields.Timeout,
			}
			if got := c.SetMaxIdleConnections(tt.args.maxConnections); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxIdleConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builderImpl_GetMaxIdleConnections(t *testing.T) {
	type fields struct {
		header  Headers
		Timeout Timeout
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test GetMaxIdleConnections",
			fields: fields{
				header:  NewHeaders(),
				Timeout: newTimeouts().SetResponseTimeout(0 * time.Second).SetRequestTimeout(0 * time.Second).SetMaxIdleConnections(10),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := builderImpl{
				header:  tt.fields.header,
				Timeout: tt.fields.Timeout,
			}
			if got := c.GetMaxIdleConnections(); got != tt.want {
				t.Errorf("GetMaxIdleConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBuilder(t *testing.T) {
	tests := []struct {
		name string
		want Builder
	}{
		{
			name: "Test NewBuilder",
			want: &builderImpl{
				header:  NewHeaders(),
				Timeout: newTimeouts(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builderImpl_Build(t *testing.T) {
	var builder = NewBuilder()

	type fields struct {
		header  Headers
		Timeout Timeout
	}
	tests := []struct {
		name   string
		fields fields
		want   Client
	}{
		{
			name: "Test Build",
			fields: fields{
				header:  NewHeaders(),
				Timeout: newTimeouts(),
			},

			want: builder.Build(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := builderImpl{
				header:  tt.fields.header,
				Timeout: tt.fields.Timeout,
			}
			if got := c.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builderImpl_SetRequestTimeout(t *testing.T) {
	type fields struct {
		header  Headers
		Timeout Timeout
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Timeout
	}{
		{
			name: "Test SetRequestTimeout",
			fields: fields{
				header:  NewHeaders(),
				Timeout: newTimeouts().SetResponseTimeout(0 * time.Second).SetRequestTimeout(20 * time.Second).SetMaxIdleConnections(10),
			},
			args: args{
				timeout: 10 * time.Second,
			},
			want: timeoutImpl{
				ResponseTimeout:    0,
				RequestTimeout:     20 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := builderImpl{
				header:  tt.fields.header,
				Timeout: tt.fields.Timeout,
			}
			if got := c.SetRequestTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRequestTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builderImpl_SetResponseTimeout(t *testing.T) {
	type fields struct {
		header  Headers
		Timeout Timeout
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Timeout
	}{
		{
			name: "Test SetResponseTimeout",
			fields: fields{
				header:  NewHeaders(),
				Timeout: newTimeouts().SetResponseTimeout(20 * time.Second).SetRequestTimeout(0 * time.Second).SetMaxIdleConnections(10),
			},
			args: args{
				timeout: 10 * time.Second,
			},
			want: timeoutImpl{
				ResponseTimeout:    20 * time.Second,
				RequestTimeout:     0,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := builderImpl{
				header:  tt.fields.header,
				Timeout: tt.fields.Timeout,
			}
			if got := c.SetResponseTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetResponseTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
