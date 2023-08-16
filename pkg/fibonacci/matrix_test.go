package fibonacci

import (
	"reflect"
	"testing"
)

func Test_matrix_multiply(t *testing.T) {
	type fields struct {
		a, b int
		c, d int
	}
	type args struct {
		m2 matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   matrix
	}{
		{
			name: "first",
			fields: fields{
				1, 1,
				1, 0,
			},
			args: args{m2: matrix{
				1, 1,
				1, 0,
			}},
			want: matrix{
				2, 1,
				1, 1,
			},
		},
		{
			name: "second",
			fields: fields{
				2, 1,
				1, 1,
			},
			args: args{m2: matrix{
				2, 1,
				1, 1,
			}},
			want: matrix{
				5, 3,
				3, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				m1 := matrix{
					a: tt.fields.a,
					b: tt.fields.b,
					c: tt.fields.c,
					d: tt.fields.d,
				}
				if got := m1.multiply(tt.args.m2); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("multiply() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_matrix_pow(t *testing.T) {
	type fields struct {
		a, b int
		c, d int
	}
	type args struct {
		n int
	}
	var (
		tests = []struct {
			name   string
			fields fields
			args   args
			want   matrix
		}{
			{
				name: "first",
				fields: fields{
					1, 1,
					1, 0,
				},
				args: args{
					n: 2,
				},
				want: matrix{
					2, 1,
					1, 1,
				},
			},
			{
				name: "second",
				fields: fields{
					1, 1,
					1, 0,
				},
				args: args{
					n: 4,
				},
				want: matrix{
					5, 3,
					3, 2,
				},
			},
			{
				name: "third",
				fields: fields{
					1, 1,
					1, 0,
				},
				args: args{
					n: 8,
				},
				want: matrix{
					34, 21,
					21, 13,
				},
			},
			{
				name: "base",
				fields: fields{
					10, 20,
					30, 40,
				},
				args: args{
					n: 0,
				},
				want: matrix{
					1, 0,
					0, 1,
				},
			},
		}
	)
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				m1 := matrix{
					a: tt.fields.a,
					b: tt.fields.b,
					c: tt.fields.c,
					d: tt.fields.d,
				}
				if got := m1.pow(tt.args.n); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("pow() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
