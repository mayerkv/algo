package bits

import "testing"

func TestBishop(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "",
			args: args{
				n: 13,
			},
			want: 4328785936,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Bishop(tt.args.n); got != tt.want {
					t.Errorf("Bishop() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
