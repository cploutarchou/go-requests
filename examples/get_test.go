package examples

import (
	"testing"
)

func Test_findPetsByTagJSON(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test_findPetsByTagJSON",
			args: args{
				tag: "cm8rvd96sgb7ev7dmli6pqz8vlpfx86egsiw6cejq1q1npe9yu45q27260b5td9ee90eiie7q49rb2xtmo26qq4shqfh6farkm8fz5ddpn7jq64dtdd16e1j8z99cesaxz65bj252y930hbsbfchir4l030z2rhuaf",
			},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findPetsByTagJSON(tt.args.tag)
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

func Test_findPetsByTagXML(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test_findPetsByTagXML",
			args: args{
				tag: "2vjdqvwzhxku1jp06yaa9v5raeco135nuabv8uh8xr3ve6c9dp5cc08v8y1kjq4cgr27cqdhrbt4eezd0l5q3ka77k1qqht84qc17rs7k7noxgzuqp6m2i",
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findPetsByTagXML(tt.args.tag)
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
