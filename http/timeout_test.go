package http

import (
	"reflect"
	"testing"
	"time"
)

func Test_newTimeouts(t *testing.T) {
	tests := []struct {
		name string
		want *timeoutImpl
	}{
		{
			name: "Test newTimeouts",
			want: &timeoutImpl{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTimeouts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTimeouts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeoutImpl_GetRequestTimeout(t *testing.T) {
	type fields struct {
		ResponseTimeout    time.Duration
		RequestTimeout     time.Duration
		MaxIdleConnections int
		DisableTimeouts    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "Test GetRequestTimeout",
			fields: fields{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
			want: 5 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := timeoutImpl{
				ResponseTimeout:    tt.fields.ResponseTimeout,
				RequestTimeout:     tt.fields.RequestTimeout,
				MaxIdleConnections: tt.fields.MaxIdleConnections,
				DisableTimeouts:    tt.fields.DisableTimeouts,
			}
			if got := c.GetRequestTimeout(); got != tt.want {
				t.Errorf("GetRequestTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeoutImpl_GetResponseTimeout(t *testing.T) {
	type fields struct {
		ResponseTimeout    time.Duration
		RequestTimeout     time.Duration
		MaxIdleConnections int
		DisableTimeouts    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "Test GetResponseTimeout",
			fields: fields{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
			want: 5 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := timeoutImpl{
				ResponseTimeout:    tt.fields.ResponseTimeout,
				RequestTimeout:     tt.fields.RequestTimeout,
				MaxIdleConnections: tt.fields.MaxIdleConnections,
				DisableTimeouts:    tt.fields.DisableTimeouts,
			}
			if got := c.GetResponseTimeout(); got != tt.want {
				t.Errorf("GetResponseTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeoutImpl_GetMaxIdleConnections(t *testing.T) {
	type fields struct {
		ResponseTimeout    time.Duration
		RequestTimeout     time.Duration
		MaxIdleConnections int
		DisableTimeouts    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test GetMaxIdleConnections",
			fields: fields{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := timeoutImpl{
				ResponseTimeout:    tt.fields.ResponseTimeout,
				RequestTimeout:     tt.fields.RequestTimeout,
				MaxIdleConnections: tt.fields.MaxIdleConnections,
				DisableTimeouts:    tt.fields.DisableTimeouts,
			}
			if got := c.GetMaxIdleConnections(); got != tt.want {
				t.Errorf("GetMaxIdleConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeoutImpl_Disable(t *testing.T) {
	type fields struct {
		ResponseTimeout    time.Duration
		RequestTimeout     time.Duration
		MaxIdleConnections int
		DisableTimeouts    bool
	}
	type args struct {
		disable bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Timeout
	}{
		{
			name: "Test DisableTimeouts Enabled",
			fields: fields{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
			args: args{
				disable: false,
			},
			want: timeoutImpl{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
		},
		{
			name: "Test DisableTimeouts Disabled",
			fields: fields{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    true,
			},
			args: args{
				disable: true,
			},
			want: timeoutImpl{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := timeoutImpl{
				ResponseTimeout:    tt.fields.ResponseTimeout,
				RequestTimeout:     tt.fields.RequestTimeout,
				MaxIdleConnections: tt.fields.MaxIdleConnections,
				DisableTimeouts:    tt.fields.DisableTimeouts,
			}
			if got := c.Disable(tt.args.disable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Disable() = %v, want %v", got, tt.want)
			}
		})
	}
}
