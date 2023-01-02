package examples

import "testing"

func Test_updatePet(t *testing.T) {
	type args struct {
		item *Pet
	}
	tests := []struct {
		name    string
		args    args
		want    *updateRes
		wantErr bool
	}{
		{
			name: "Test_updatePet",
			args: args{
				item: &Pet{
					Category: struct {
						Name string "json:\"name\""
						ID   int64  "json:\"id\""
					}{Name: "string", ID: 0},
					Name:      "doggie",
					PhotoUrls: []string{"string"},
					Tags: []struct {
						Name string "json:\"name\""
						ID   int64  "json:\"id\""
					}{{Name: "string", ID: 0}},
					Status: "available",
					ID:     0,
				},
			},
			want: &updateRes{
				Code:   200,
				Status: "ok",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updatePet(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("updatePet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Code != tt.want.Code {
				t.Errorf("updatePet() got = %v, want %v", got, tt.want)
			}
		})
	}
}
