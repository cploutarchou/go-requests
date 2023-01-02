package examples

import (
	"reflect"
	"testing"
	"time"
)

func Test_placePetOrder(t *testing.T) {
	type args struct {
		item Order
	}
	t1, _ := time.Parse("2006-01-02 15:04:05", "2022-12-12 00:00:00")
	tests := []struct {
		name    string
		args    args
		want    Order
		wantErr bool
	}{
		{
			name: "Test_placePetOrder",
			args: args{
				item: Order{
					PetId:    670792158758028421,
					Quantity: 2,
					Id:       6075746898333402660,
					ShipDate: t1,
					Complete: true,
					Status:   "approved",
				},
			},
			want: Order{
				PetId:    670792158758028421,
				Quantity: 2,
				Id:       6075746898333402660,
				ShipDate: t1,
				Complete: true,
				Status:   "approved",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := placePetOrder(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("placePetOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Id != tt.want.Id {
				t.Errorf("placePetOrder() = %v, want %v", got.Id, tt.want.Id)
				return
			}

			// check if got is deep equal to want
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("placePetOrder() = \n%v \n, want \n%v \n", got, tt.want)
				return
			}
		})
	}
}
