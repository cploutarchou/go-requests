package http

import (
	"reflect"
	"testing"
)

func Test_headerImpl_Set(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test Set",
			fields: fields{
				values: map[string]string{
					"key": "value",
				},
			},
			args: args{
				key:   "key",
				value: "value",
			},
			want: &headerImpl{
				values: map[string]string{
					"key": "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.Set(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
