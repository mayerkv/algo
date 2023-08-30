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
			name: "add to head",
			fields: fields{
				arr: []int{0, 1, 2, 3, 4},
			},
			args: args{
				item:  10,
				index: 0,
			},
			want: []int{10, 0, 1, 2, 3, 4},
		},
		{
			name: "add to tail",
			fields: fields{
				arr: []int{0, 1, 2, 3, 4},
			},
			args: args{
				item:  10,
				index: 5,
			},
			want: []int{0, 1, 2, 3, 4, 10},
		},
		{
			name: "add in center",
			fields: fields{
				arr: []int{0, 1, 2, 3, 4},
			},
			args: args{
				item:  10,
				index: 3,
			},
			want: []int{0, 1, 2, 10, 3, 4},
		},
		{
			name: "add with out of range",
			fields: fields{
				arr: []int{0, 1, 2, 3, 4},
			},
			args: args{
				item:  10,
				index: 10,
			},
			want: []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0, 10},
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
		name    string
		fields  fields
		args    args
		want    int
		wantArr []int
	}{
		{
			name: "remove first",
			fields: fields{
				arr: []int{0, 1, 2, 3, 4},
			},
			args: args{
				index: 0,
			},
			want:    0,
			wantArr: []int{1, 2, 3, 4, 0},
		},
		{
			name: "remove last",
			fields: fields{
				arr: []int{0, 1, 2, 3, 4},
			},
			args: args{
				index: 4,
			},
			want:    4,
			wantArr: []int{0, 1, 2, 3, 0},
		},
		{
			name: "remove in center",
			fields: fields{
				arr: []int{0, 1, 2, 3, 4},
			},
			args: args{
				index: 2,
			},
			want:    2,
			wantArr: []int{0, 1, 3, 4, 0},
		},
		{
			name: "remove with out of range",
			fields: fields{
				arr: []int{0, 1, 2, 3, 4},
			},
			args: args{
				index: 10,
			},
			want:    0,
			wantArr: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				a := &SingleArray[int]{
					arr: tt.fields.arr,
				}
				actual := a.Remove(tt.args.index)
				require.Equal(t, tt.want, actual)
				require.Equal(t, tt.wantArr, a.arr)
			},
		)
	}
}
