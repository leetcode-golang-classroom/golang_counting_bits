package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	n := 12
	for idx := 0; idx < b.N; idx++ {
		countBits(n)
	}
}
func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "n = 2",
			args: args{n: 2},
			want: []int{0, 1, 1},
		},
		{
			name: "n = 5",
			args: args{n: 5},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
