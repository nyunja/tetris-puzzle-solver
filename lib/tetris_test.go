package lib

import (
	"reflect"
	"testing"
)

func Test_calculateInitialGridSize(t *testing.T) {
	type args struct {
		tet *Tetrominos
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{name: "calculateInitialGridSize", args: args{tet: &Tetrominos{[][]string{
			{"...#", "...#", "...#", "...#"},
		}}}, want: 4, want1: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calculateInitialGridSize(tt.args.tet)
			if got != tt.want {
				t.Errorf("calculateInitialGridSize() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calculateInitialGridSize() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_createGrid(t *testing.T) {
	type args struct {
		maxWidth  int
		maxHeight int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "createGrid", args: args{maxWidth: 4, maxHeight: 4}, want: [][]string{
			{".", ".", ".", "."},
			{".", ".", ".", "."},
			{".", ".", ".", "."},
			{".", ".", ".", "."},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createGrid(tt.args.maxWidth, tt.args.maxHeight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCleanTetromino(t *testing.T) {
	type args struct {
		tet *Tetrominos
	}
	tests := []struct {
		name    string
		args    args
		want    *Tetrominos
		wantErr bool
	}{
		{name: "CleanTetromino", 
		args: args{tet: &Tetrominos{[][]string{
			{"...#", "...#", "...#", "...#"},
		}}}, 
		want: &Tetrominos{[][]string{
			{"#", "#", "#", "#"},
		}}, 
		wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CleanTetromino(tt.args.tet)
			if (err != nil) != tt.wantErr {
				t.Errorf("CleanTetromino() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CleanTetromino() = %v, want %v", got, tt.want)
			}
		})
	}
}
