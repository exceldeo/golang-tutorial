package park

import (
	"testing"
)

func TestCalculateParkingBill(t *testing.T) {
	type args struct {
		vehicle  VehicleType
		duration int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "testing car with 48 hours",
			args: args{
				vehicle:  VehicleType(0),
				duration: 48,
			},
			want: 577000,
		},
		{
			name: "testing car with 0 hours",
			args: args{
				vehicle:  VehicleType(0),
				duration: 0,
			},
			want: 7000,
		},
		{
			name: "testing motocycle with 8 hours",
			args: args{
				vehicle:  VehicleType(1),
				duration: 8,
			},
			want: 38000,
		},
		{
			name: "testing motocycle with 0 hours",
			args: args{
				vehicle:  VehicleType(0),
				duration: 0,
			},
			want: 7000,
		},
		{
			name: "testing Bicycle with 8 hours",
			args: args{
				vehicle:  VehicleType(2),
				duration: 8,
			},
			want: 7500,
		},
		{
			name: "testing Bicycle with 0 hours",
			args: args{
				vehicle:  VehicleType(3),
				duration: 0,
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateParkingBill(tt.args.vehicle, tt.args.duration); got != tt.want {
				t.Errorf("CalculateParkingBill() = %v, want %v", got, tt.want)
			}
		})
	}
}
