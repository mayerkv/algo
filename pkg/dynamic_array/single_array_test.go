package dynamic_array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleArray_Add(t *testing.T) {
	type fields struct {
		arr []int
	}
	type args struct {
		item  int
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "add 7 to 0",
			fields: fields{
				arr: []int{1, 2, 3, 4, 5},
			},
			args: args{
				item:  7,
				index: 0,
			},
			want: []int{7, 2, 3, 4, 5},
		},
		{
			name: "add 7 to 10 (out of range)",
			fields: fields{
				arr: []int{1, 2, 3, 4, 5},
			},
			args: args{
				item:  7,
				index: 10,
			},
			want: []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 7},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				a := &SingleArray[int]{
					arr: tt.fields.arr,
				}
				a.Add(tt.args.item, tt.args.index)
				require.Equal(t, tt.want, a.arr)
			},
		)
	}
}

func TestSingleArray_Remove(t *testing.T) {
	type fields struct {
		arr []int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		arr    []int
	}{
		{
			name: "remove 0",
			fields: fields{
				arr: []int{1, 2, 3, 4, 5},
			},
			args: args{
				index: 0,
			},
			want: 1,
			arr:  []int{2, 3, 4, 5},
		},
		{
			name: "remove 3",
			fields: fields{
				arr: []int{1, 2, 3, 4, 5},
			},
			args: args{
				index: 2,
			},
			want: 3,
			arr:  []int{1, 2, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				a := &SingleArray[int]{
					arr: tt.fields.arr,
				}
				got := a.Remove(tt.args.index)
				require.Equal(t, tt.want, got)
				require.Equal(t, tt.arr, a.arr)
			},
		)
	}
}
