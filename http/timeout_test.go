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
