package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_list_addSort(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *list[int]
	}{
		{
			name: "add to center",
			fields: fields{
				head: &node[int]{
					value: 1,
					next: &node[int]{
						value: 2,
						next: &node[int]{
							value: 4,
						},
					},
				},
			},
			args: args{
				value: 3,
			},
			want: &list[int]{
				head: &node[int]{
					value: 1,
					next: &node[int]{
						value: 2,
						next: &node[int]{
							value: 3,
							next: &node[int]{
								value: 4,
							},
						},
					},
				},
			},
		},
		{
			name: "empty list",
			fields: fields{
				head: nil,
			},
			args: args{
				value: 1,
			},
			want: &list[int]{
				head: &node[int]{value: 1},
			},
		},
		{
			name: "add to head",
			fields: fields{
				head: &node[int]{value: 2},
			},
			args: args{
				value: 1,
			},
			want: &list[int]{
				head: &node[int]{
					value: 1,
					next: &node[int]{
						value: 2,
					},
				},
			},
		},
		{
			name: "add to tail",
			fields: fields{
				head: &node[int]{value: 1},
			},
			args: args{
				value: 4,
			},
			want: &list[int]{
				head: &node[int]{
					value: 1,
					next: &node[int]{
						value: 4,
					},
				},
			},
		},
		{
			name: "add double",
			fields: fields{
				head: &node[int]{value: 1},
			},
			args: args{
				value: 1,
			},
			want: &list[int]{
				head: &node[int]{
					value: 1,
					next: &node[int]{
						value: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				l := &list[int]{
					head: tt.fields.head,
				}
				l.addSort(tt.args.value)
				require.Equal(t, l, tt.want)
			},
		)
	}
}

func TestBucketSort(t *testing.T) {
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
				BucketSort(tt.args.arr)
				require.Equal(t, tt.want, tt.args.arr)
			},
		)
	}
}

func BenchmarkBucketSort(b *testing.B) {
	inputSize := []int{100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000}
	for _, size := range inputSize {
		b.Run(
			fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					testList := Random(size)
					BucketSort(testList)
				}
			},
		)
	}
}
