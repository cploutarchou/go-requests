package http

import (
	"net/http"
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

func Test_headerImpl_SetContentLength(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		contentLength int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetContentLength",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				contentLength: 10,
			},
			want: &headerImpl{
				values: map[string]string{
					"Content-Length": "10",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetContentLength(tt.args.contentLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetContentLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetContentDisposition(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		contentDisposition string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{

		{
			name: "Test SetContentDisposition",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				contentDisposition: "attachment",
			},
			want: &headerImpl{
				values: map[string]string{
					"Content-Disposition": "attachment",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetContentDisposition(tt.args.contentDisposition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetContentDisposition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHeaders(t *testing.T) {
	tests := []struct {
		name string
		want Headers
	}{
		{
			name: "Test NewHeaders",
			want: &headerImpl{
				values: map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeaders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_Clone(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   Headers
	}{
		{
			name: "Test Clone",
			fields: fields{
				values: map[string]string{
					"key": "value",
				},
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
			if got := h.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_Del(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test Del",
			fields: fields{
				values: map[string]string{
					"key": "value",
				},
			},
			args: args{
				key: "key",
			},
			want: &headerImpl{
				values: map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.Del(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_Get(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test Get",
			fields: fields{
				values: map[string]string{
					"key": "value",
				},
			},
			args: args{
				key: "key",
			},
			want: "value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.Get(tt.args.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_GetAll(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string][]string
	}{
		{
			name: "Test GetAll",
			fields: fields{
				values: map[string]string{
					"key":  "value",
					"key2": "value2",
				},
			},
			want: map[string][]string{
				"key":  []string{"value"},
				"key2": []string{"value2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_GetAllHttpHeaders(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   http.Header
	}{
		{
			name: "Test GetAllHttpHeaders",
			fields: fields{
				values: map[string]string{
					"key": "value",
				},
			},
			want: http.Header{
				"Key": []string{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.GetAllHttpHeaders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllHttpHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_IsEmpty(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test IsEmpty",
			fields: fields{
				values: map[string]string{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_IsSet(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test IsSet",
			fields: fields{
				values: map[string]string{
					"key": "value",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.IsSet(); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_Keys(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Test Keys",
			fields: fields{
				values: map[string]string{
					"key":  "value",
					"key2": "value2",
				},
			},
			want: []string{"key", "key2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_Len(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test Len",
			fields: fields{
				values: map[string]string{
					"key":  "value",
					"key2": "value2",
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetAccept(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		accept string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetAccept",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				accept: "application/json",
			},
			want: &headerImpl{
				values: map[string]string{
					"Accept": "application/json",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetAccept(tt.args.accept); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAccept() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetAcceptCharset(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		acceptCharset string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetAcceptCharset",
			fields: fields{
				values: map[string]string{},
			},

			args: args{
				acceptCharset: "utf-8",
			},
			want: &headerImpl{
				values: map[string]string{
					"Accept-Charset": "utf-8",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetAcceptCharset(tt.args.acceptCharset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAcceptCharset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetContentEncoding(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		contentEncoding string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetContentEncoding",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				contentEncoding: "gzip",
			},
			want: &headerImpl{
				values: map[string]string{
					"Content-Encoding": "gzip",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetContentEncoding(tt.args.contentEncoding); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetContentEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetContentLanguage(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		contentLanguage string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetContentLanguage",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				contentLanguage: "en",
			},
			want: &headerImpl{
				values: map[string]string{
					"Content-Language": "en",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetContentLanguage(tt.args.contentLanguage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetContentLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}