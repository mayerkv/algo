package search

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBST_Insert(t1 *testing.T) {
	type fields struct {
		Tree Tree
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Tree
	}{
		{
			name: "insert to root",
			fields: fields{
				Tree: Tree{},
			},
			args: args{
				x: 1,
			},
			want: Tree{
				root: &Node{value: 1},
			},
		},
		{
			name: "insert great than root",
			fields: fields{
				Tree: Tree{
					root: &Node{value: 10},
				},
			},
			args: args{
				x: 11,
			},
			want: Tree{
				root: &Node{
					value: 10,
					right: &Node{value: 11},
				},
			},
		},
		{
			name: "insert less than root",
			fields: fields{
				Tree: Tree{
					root: &Node{value: 10},
				},
			},
			args: args{
				x: 9,
			},
			want: Tree{
				root: &Node{
					value: 10,
					left:  &Node{value: 9},
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(
			tt.name, func(t1 *testing.T) {
				t := &BST{
					Tree: tt.fields.Tree,
				}
				t.Insert(tt.args.x)
				require.Equal(t1, t.Tree, tt.want)
			},
		)
	}
}

func BenchmarkBST_Insert(b *testing.B) {
	rnd := rand.New(rand.NewSource(1))
	bst := NewBST()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Insert(rnd.Int())
	}
}
