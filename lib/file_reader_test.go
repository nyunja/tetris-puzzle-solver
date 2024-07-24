package lib

import (
	"reflect"
	"testing"
)

func TestInputFileReader(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		want    *Tetrominos
		wantErr bool
	}{
		{
			name: "InputFileReader",
			file: "sample1.txt",
			want: &Tetrominos{Tet: [][]string{
				{"....","AAA.",".A..","...."},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InputFileReader(tt.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("InputFileReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InputFileReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
