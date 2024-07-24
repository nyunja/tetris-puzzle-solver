package lib

import (
	"testing"
)

func Test_fits(t *testing.T) {
	type args struct {
		tetromino   []string
		tetSolution [][]string
		x           int
		y           int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "fits",
			args: args{
				tetromino: []string{"#", "#", "#", "#"},
				tetSolution: [][]string{
					{".", ".", ".", "."},
					{".", ".", ".", "."},
					{".", ".", ".", "."},
					{".", ".", ".", "."},
				},
				x: 0,
				y: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fits(tt.args.tetromino, tt.args.tetSolution, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("fits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtrackTetris(t *testing.T) {
	type args struct {
		tetrominos  [][]string
		tetSolution [][]string
		index       int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "backtrackTetris",
			args: args{
				tetrominos: [][]string{
					{"#", "#", "#", "#"},
					{"##", "##"},
				},
				tetSolution: [][]string{
					{"#", "#", "#", "."},
					{"#", "#", "#", "."},
					{"#", ".", ".", "."},
					{"#", ".", ".", "."},
				},
				index: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backtrackTetris(tt.args.tetrominos, tt.args.tetSolution, tt.args.index); got != tt.want {
				t.Errorf("backtrackTetris() = %v, want %v", got, tt.want)
			}
		})
	}
}
