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
				DisableTimeouts:    true,
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
			if got := c.Disable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Disable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeoutImpl_SetRequestTimeout(t *testing.T) {
	type fields struct {
		ResponseTimeout    time.Duration
		RequestTimeout     time.Duration
		MaxIdleConnections int
		DisableTimeouts    bool
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
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
			args: args{
				timeout: 10 * time.Second,
			},
			want: timeoutImpl{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     10 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
		},
		{
			name: "Test SetRequestTimeout",
			fields: fields{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
			args: args{
				timeout: 0,
			},
			want: timeoutImpl{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     0,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
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
			if got := c.SetRequestTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRequestTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeoutImpl_SetResponseTimeout(t *testing.T) {
	type fields struct {
		ResponseTimeout    time.Duration
		RequestTimeout     time.Duration
		MaxIdleConnections int
		DisableTimeouts    bool
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
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},
			args: args{
				timeout: 10 * time.Second,
			},
			want: timeoutImpl{
				ResponseTimeout:    10 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
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
			if got := c.SetResponseTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetResponseTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeoutImpl_SetMaxIdleConnections(t *testing.T) {
	type fields struct {
		ResponseTimeout    time.Duration
		RequestTimeout     time.Duration
		MaxIdleConnections int
		DisableTimeouts    bool
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
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 10,
				DisableTimeouts:    false,
			},

			args: args{
				maxConnections: 20,
			},
			want: timeoutImpl{
				ResponseTimeout:    5 * time.Second,
				RequestTimeout:     5 * time.Second,
				MaxIdleConnections: 20,
				DisableTimeouts:    false,
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
			if got := c.SetMaxIdleConnections(tt.args.maxConnections); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxIdleConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}
