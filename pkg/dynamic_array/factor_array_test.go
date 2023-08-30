package dynamic_array

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFactorArray_Add(t *testing.T) {
	type fields struct {
		arr    []int
		factor int
		size   int
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
				arr:    []int{0, 1, 2, 3, 4},
				factor: 50,
				size:   5,
			},
			args: args{
				item:  10,
				index: 0,
			},
			want: []int{10, 0, 1, 2, 3, 4, 0},
		},
		{
			name: "add to tail",
			fields: fields{
				arr:    []int{0, 1, 2, 3, 4},
				factor: 50,
				size:   5,
			},
			args: args{
				item:  10,
				index: 5,
			},
			want: []int{0, 1, 2, 3, 4, 10, 0},
		},
		{
			name: "add in center",
			fields: fields{
				arr:    []int{0, 1, 2, 3, 4},
				factor: 50,
				size:   5,
			},
			args: args{
				item:  10,
				index: 3,
			},
			want: []int{0, 1, 2, 10, 3, 4, 0},
		},
		{
			name: "add with out of range",
			fields: fields{
				arr:    []int{0, 1, 2, 3, 4},
				factor: 50,
				size:   5,
			},
			args: args{
				item:  10,
				index: 10,
			},
			want: []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0, 10, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				a := &FactorArray[int]{
					arr:    tt.fields.arr,
					factor: tt.fields.factor,
					size:   tt.fields.size,
				}
				a.Add(tt.args.item, tt.args.index)
				require.Equal(t, tt.want, a.arr)
			},
		)
	}
}

func TestFactorArray_Remove(t *testing.T) {
	type fields struct {
		arr    []int
		factor int
		size   int
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
				arr:    []int{0, 1, 2, 3, 4},
				size:   5,
				factor: 50,
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
				arr:    []int{0, 1, 2, 3, 4},
				size:   5,
				factor: 50,
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
				arr:    []int{0, 1, 2, 3, 4},
				size:   5,
				factor: 50,
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
				arr:    []int{0, 1, 2, 3, 4},
				size:   5,
				factor: 50,
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
				a := &FactorArray[int]{
					arr:    tt.fields.arr,
					factor: tt.fields.factor,
					size:   tt.fields.size,
				}
				if got := a.Remove(tt.args.index); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Remove() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func BenchmarkFactorArray_Add(b *testing.B) {
	a := &FactorArray[int]{
		arr:    make([]int, b.N),
		factor: 50,
		size:   b.N,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Add(b.N, i)
	}
}

func BenchmarkFactorArray_Remove(b *testing.B) {
	a := &FactorArray[int]{
		arr:    make([]int, b.N),
		factor: 50,
		size:   b.N,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Remove(i)
	}
}
