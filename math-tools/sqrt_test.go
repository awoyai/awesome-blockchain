package math_tools

import (
	"math/big"
	"reflect"
	"testing"
)

func TestSqrt(t *testing.T) {
	type args struct {
		y2 *big.Int
		p  *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "sqrt(36), p=19",
			args: args{y2: big.NewInt(36), p: big.NewInt(19)},
			want: big.NewInt(6),
		},{
			name: "sqrt(49), p=19",
			args: args{y2: big.NewInt(49), p: big.NewInt(19)},
			want: big.NewInt(7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.args.y2, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
