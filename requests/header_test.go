package requests

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
				"key":  {"value"},
				"key2": {"value2"},
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
					"key": "value",
				},
			},
			want: []string{"key"},
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
					"key": "value",
				},
			},
			want: 1,
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

func Test_headerImpl_SetContentLocation(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		contentLocation string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetContentLocation",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				contentLocation: "https://www.example.com",
			},
			want: &headerImpl{
				values: map[string]string{
					"Content-Location": "https://www.example.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetContentLocation(tt.args.contentLocation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetContentLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetContentMD5(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		contentMD5 string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetContentMD5",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				contentMD5: "Q2hlY2sgSW50ZWdyaXR5IQ==",
			},
			want: &headerImpl{
				values: map[string]string{
					"Content-MD5": "Q2hlY2sgSW50ZWdyaXR5IQ==",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetContentMD5(tt.args.contentMD5); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetContentMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetContentRange(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		contentRange string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetContentRange",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				contentRange: "bytes 200-1000/67589",
			},
			want: &headerImpl{
				values: map[string]string{
					"Content-Range": "bytes 200-1000/67589",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetContentRange(tt.args.contentRange); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetContentRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetCookie(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		cookie string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetCookie",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				cookie: "theme=light;Token=abc123",
			},
			want: &headerImpl{
				values: map[string]string{
					"Cookie": "theme=light;Token=abc123",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetCookie(tt.args.cookie); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetDate(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		date string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetDate",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				date: "Tue, 15 Nov 1994 08:12:31 GMT",
			},
			want: &headerImpl{
				values: map[string]string{
					"Date": "Tue, 15 Nov 1994 08:12:31 GMT",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetDate(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetETag(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		etag string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetETag",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				etag: "737060cd8c284d8af7ad3082f209582d",
			},
			want: &headerImpl{
				values: map[string]string{
					"ETag": "737060cd8c284d8af7ad3082f209582d",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetETag(tt.args.etag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetETag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetExpires(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		expires string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetExpires",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				expires: "Thu, 22 Dec 2022 16:00:00 GMT",
			},
			want: &headerImpl{
				values: map[string]string{
					"Expires": "Thu, 22 Dec 2022 16:00:00 GMT",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetExpires(tt.args.expires); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExpires() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetAcceptLanguage(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		acceptLanguage string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetAcceptLanguage",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				acceptLanguage: "en-US",
			},
			want: &headerImpl{
				values: map[string]string{
					"Accept-Language": "en-US",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetAcceptLanguage(tt.args.acceptLanguage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAcceptLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetAcceptRanges(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		acceptRanges string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetAcceptRanges",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				acceptRanges: "bytes",
			},
			want: &headerImpl{
				values: map[string]string{
					"Accept-Ranges": "bytes",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetAcceptRanges(tt.args.acceptRanges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAcceptRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetAge(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		age string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetAge",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				age: "12",
			},
			want: &headerImpl{
				values: map[string]string{
					"Age": "12",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetAge(tt.args.age); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_SetAllow(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	type args struct {
		allow string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Headers
	}{
		{
			name: "Test SetAllow",
			fields: fields{
				values: map[string]string{},
			},
			args: args{
				allow: "GET, HEAD",
			},
			want: &headerImpl{
				values: map[string]string{
					"Allow": "GET, HEAD",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.SetAllow(tt.args.allow); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAllow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerImpl_Values(t *testing.T) {
	type fields struct {
		values map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name: "Test Values",
			fields: fields{
				values: map[string]string{
					"Accept": "application/json",
				},
			},
			want: map[string]string{
				"Accept": "application/json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headerImpl{
				values: tt.fields.values,
			}
			if got := h.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
