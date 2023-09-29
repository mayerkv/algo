package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "random",
			args: args{
				arr: []int{5, 8, 3, 4, 2, 0, 7, 1, 9, 6},
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				CountSort(tt.args.arr)
				require.Equal(t, tt.want, tt.args.arr)
			},
		)
	}
}

func BenchmarkCountSort(b *testing.B) {
	inputSize := []int{100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000}
	for _, size := range inputSize {
		b.Run(
			fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
				testList := Random(size)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					clone := make([]int, size)
					copy(clone, testList)
					CountSort(clone)
				}
			},
		)
	}
}
