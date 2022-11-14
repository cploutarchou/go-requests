package http

import (
	"reflect"
	"testing"
)

func Test_authorizationImpl_Bearer(t *testing.T) {
	type fields struct {
		authorizationType AuthorizationType
		value             string
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Authorization
	}{
		{
			name:   "test1",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			args:   args{token: "test"},
			want:   &authorizationImpl{authorizationType: AuthorizationTypeBearer, value: "test"},
		},
		{
			name:   "test2",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			args:   args{token: ""},
			want:   &authorizationImpl{authorizationType: AuthorizationTypeBearer, value: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizationImpl{
				authorizationType: tt.fields.authorizationType,
				value:             tt.fields.value,
			}
			if got := a.Bearer(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authorizationImpl.Bearer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorizationImpl_Basic(t *testing.T) {
	type fields struct {
		authorizationType AuthorizationType
		value             string
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Authorization
	}{

		{
			name:   "test1",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			args:   args{username: "test", password: "test"},
			want:   &authorizationImpl{authorizationType: AuthorizationTypeBasic, value: "test:test"},
		},
		{
			name:   "test2",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			args:   args{username: "", password: ""},
			want:   &authorizationImpl{authorizationType: AuthorizationTypeBasic, value: ":"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizationImpl{
				authorizationType: tt.fields.authorizationType,
				value:             tt.fields.value,
			}
			if got := a.Basic(tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authorizationImpl.Basic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorizationImpl_String(t *testing.T) {
	type fields struct {
		authorizationType AuthorizationType
		value             string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "test1",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			want:   "Basic test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizationImpl{
				authorizationType: tt.fields.authorizationType,
				value:             tt.fields.value,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("authorizationImpl.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorizationImpl_Type(t *testing.T) {
	type fields struct {
		authorizationType AuthorizationType
		value             string
	}
	tests := []struct {
		name   string
		fields fields
		want   AuthorizationType
	}{
		{
			name:   "test1",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			want:   AuthorizationTypeBasic,
		},

		{
			name:   "test2",
			fields: fields{authorizationType: AuthorizationTypeBearer, value: "test"},
			want:   AuthorizationTypeBearer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizationImpl{
				authorizationType: tt.fields.authorizationType,
				value:             tt.fields.value,
			}
			if got := a.Type(); got != tt.want {
				t.Errorf("authorizationImpl.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorizationImpl_Value(t *testing.T) {
	type fields struct {
		authorizationType AuthorizationType
		value             string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "test1",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			want:   "test",
		},
		{
			name:   "test2",
			fields: fields{authorizationType: AuthorizationTypeBearer, value: "test"},
			want:   "test",
		},
		{
			name:   "test3",
			fields: fields{authorizationType: AuthorizationTypeBearer, value: ""},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizationImpl{
				authorizationType: tt.fields.authorizationType,
				value:             tt.fields.value,
			}
			if got := a.Value(); got != tt.want {
				t.Errorf("authorizationImpl.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorizationImpl_IsBasic(t *testing.T) {
	type fields struct {
		authorizationType AuthorizationType
		value             string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{

		{
			name:   "test1",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			want:   true,
		},
		{
			name:   "test2",
			fields: fields{authorizationType: AuthorizationTypeBearer, value: "test"},
			want:   false,
		},
		{
			name:   "test3",
			fields: fields{authorizationType: AuthorizationTypeBearer, value: ""},
			want:   false,
		},
		{
			name:   "test4",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: ""},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizationImpl{
				authorizationType: tt.fields.authorizationType,
				value:             tt.fields.value,
			}
			if got := a.IsBasic(); got != tt.want {
				t.Errorf("authorizationImpl.IsBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorizationImpl_IsBearer(t *testing.T) {
	type fields struct {
		authorizationType AuthorizationType
		value             string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{

		{
			name:   "test1",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			want:   false,
		},
		{
			name:   "test2",
			fields: fields{authorizationType: AuthorizationTypeBearer, value: "test"},
			want:   true,
		},
		{
			name:   "test3",
			fields: fields{authorizationType: AuthorizationTypeBearer, value: ""},
			want:   true,
		},
		{
			name:   "test4",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: ""},
			want:   false,
		},
		{
			name:   "test5",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizationImpl{
				authorizationType: tt.fields.authorizationType,
				value:             tt.fields.value,
			}
			if got := a.IsBearer(); got != tt.want {
				t.Errorf("authorizationImpl.IsBearer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorizationImpl_IsEmpty(t *testing.T) {
	type fields struct {
		authorizationType AuthorizationType
		value             string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "test1",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			want:   false,
		},
		{
			name:   "test2",
			fields: fields{authorizationType: AuthorizationTypeBearer, value: "test"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizationImpl{
				authorizationType: tt.fields.authorizationType,
				value:             tt.fields.value,
			}
			if got := a.IsEmpty(); got != tt.want {
				t.Errorf("authorizationImpl.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorizationImpl_IsSet(t *testing.T) {
	type fields struct {
		authorizationType AuthorizationType
		value             string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "test1",
			fields: fields{authorizationType: AuthorizationTypeBasic, value: "test"},
			want:   true,
		},
		{
			name:   "test2",
			fields: fields{authorizationType: AuthorizationTypeBearer, value: "test"},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizationImpl{
				authorizationType: tt.fields.authorizationType,
				value:             tt.fields.value,
			}
			if got := a.IsSet(); got != tt.want {
				t.Errorf("authorizationImpl.IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
